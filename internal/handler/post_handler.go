package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"forum/models"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ctxUserKey).(models.User)
	if user.Username == "" {
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
	}
	switch r.Method {
	case http.MethodGet:
		if err := h.templExecute(w, "./ui/create_post.html", user); err != nil {
			return
		}
	case http.MethodPost:
		title := r.FormValue("tittle")
		r.ParseForm()
		categories := r.Form["categories"]
		category := strings.Join(categories, " ")
		content := r.FormValue("content")
		if strings.TrimSpace(title) == "" || strings.TrimSpace(content) == "" || strings.TrimSpace(category) == "" {
			http.Redirect(w, r, "/post/create", http.StatusSeeOther)
			return
		}
		post := models.Post{
			Title:    title,
			Category: category,
			Content:  content,
			Author:   user.Username,
			AuthorID: user.ID,
			Date:     time.Now().Format("January 2, 2006"),
		}
		err := h.services.CreatePost(&post)
		if err != nil {
			log.Println(err)
			h.ErrorPage(w, http.StatusInternalServerError, "can not creat post")
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) myPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	user := r.Context().Value(ctxUserKey).(models.User)
	posts, err := h.services.Post.MyPosts(strconv.Itoa(user.ID))
	if err != nil {
		h.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	pageData := models.PageData{
		Username: user.Username,
		Posts:    *posts,
	}
	if err := h.templExecute(w, "./ui/index.html", pageData); err != nil {
		fmt.Println("my posts: templExecute()")
		return
	}
}

func (h *Handler) myFavourites(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	user := r.Context().Value(ctxUserKey).(models.User)
	posts, err := h.services.Post.MyFavourites(user.ID)
	if err != nil {
		h.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	pageData := models.PageData{
		Username: user.Username,
		Posts:    *posts,
	}
	if err := h.templExecute(w, "./ui/index.html", pageData); err != nil {
		fmt.Println("my favourites: templExecute()")
		return
	}
}

func (h *Handler) post(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ctxUserKey).(models.User)
	PostID = strings.TrimPrefix(r.URL.Path, "/post/")
	post, err := h.services.Post.GetPostByID(PostID)
	var emptyPost models.Post
	if err != nil || *post == emptyPost {
		h.ErrorPage(w, http.StatusNotFound, "Sorry, but no pages were found for your request =(")
		return
	}
	postData := models.PostData{
		Username: user.Username,
		Post:     *post,
		Comments: []models.Comment{},
	}
	switch r.Method {
	case http.MethodGet:
		comments, err := h.services.Comment.GetCommentByPostID(post.ID)
		if err != nil {
			h.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		postData.Comments = *comments
	case http.MethodPost:
		user1 := models.User{}
		if user == user1 {
			http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
			return
		}
		commentText := r.FormValue("comment")
		if strings.TrimSpace(commentText) == "" {
			http.Redirect(w, r, "/post/"+PostID, http.StatusSeeOther)
			return
		}
		comment := models.Comment{
			UserID: user.ID,
			PostID: post.ID,
			Text:   commentText,
			Author: user.Username,
			Date:   time.Now().Format("01-02-2006 15:04:05"),
		}
		if err := h.services.Comment.CreateComment(comment); err != nil {
			h.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if err := h.templExecute(w, "./ui/post.html", postData); err != nil {
		fmt.Printf("post_handler: %s\n", err)
		return
	}
}
