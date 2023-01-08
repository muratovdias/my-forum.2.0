package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muratovdias/my-forum.2.0/models"
)

var post *models.Post

func (h *Handler) createPostGET(c *gin.Context) {
	userInterface, _ = c.Get("user")
	user, ok = userInterface.(models.User)
	if err := h.templExecute(c.Writer, "./ui/create_post.html", user); err != nil {
		return
	}
}

func (h *Handler) createPostPOST(c *gin.Context) {
	userInterface, _ = c.Get("user")
	user, ok = userInterface.(models.User)
	if err := c.Request.ParseForm(); err != nil {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, "something went wrong")
		return
	}
	title, ok1 := c.Request.Form["tittle"]
	categories, ok2 := c.Request.Form["categories"]
	category := strings.Join(categories, " ")
	content, ok3 := c.Request.Form["content"]
	if !ok1 || !ok2 || !ok3 {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(title[0]) == "" || strings.TrimSpace(content[0]) == "" || strings.TrimSpace(category) == "" {
		http.Redirect(c.Writer, c.Request, "/post/create", http.StatusSeeOther)
		return
	}
	post := models.Post{
		Title:    title[0],
		Category: category,
		Content:  content[0],
		Author:   user.Username,
		AuthorID: user.ID,
		Date:     time.Now().Format("January 2, 2006"),
	}
	err := h.services.CreatePost(&post)
	if err != nil {
		log.Println(err)
		h.ErrorPage(c.Writer, http.StatusInternalServerError, "can not creat post")
	}
	c.Redirect(http.StatusFound, "/")
}

func (h *Handler) myPosts(c *gin.Context) {
	userInterface, _ = c.Get("user")
	user, ok = userInterface.(models.User)
	if !ok {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	posts, err := h.services.Post.MyPosts(strconv.Itoa(user.ID))
	if err != nil {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	pageData := models.PageData{
		Username: user.Username,
		Posts:    *posts,
	}
	if err := h.templExecute(c.Writer, "./ui/index.html", pageData); err != nil {
		fmt.Println("my posts: templExecute()")
		return
	}
}

func (h *Handler) myFavourites(c *gin.Context) {
	userInterface, _ = c.Get("user")
	user, ok = userInterface.(models.User)
	if !ok {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	posts, err := h.services.Post.MyFavourites(user.ID)
	if err != nil {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	pageData := models.PageData{
		Username: user.Username,
		Posts:    *posts,
	}
	if err := h.templExecute(c.Writer, "./ui/index.html", pageData); err != nil {
		fmt.Println("my favourites: templExecute()")
		return
	}
}

func (h *Handler) postGET(c *gin.Context) {
	userInterface, _ = c.Get("user")
	user, ok = userInterface.(models.User)
	if !ok {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	postID = c.Param("id")
	post, err = h.services.Post.GetPostByID(postID)
	var emptyPost models.Post
	if err != nil || *post == emptyPost {
		h.ErrorPage(c.Writer, http.StatusNotFound, "Sorry, but no pages were found for your request =(")
		return
	}
	postData := models.PostData{
		Username: user.Username,
		Post:     *post,
		Comments: []models.Comment{},
	}
	comments, err := h.services.Comment.GetCommentByPostID(post.ID)
	if err != nil {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	postData.Comments = *comments

	if err := h.templExecute(c.Writer, "./ui/post.html", postData); err != nil {
		fmt.Printf("post_handler: %s\n", err)
		return
	}
}

func (h *Handler) postPOST(c *gin.Context) {
	user1 := models.User{}
	if user == user1 {
		http.Redirect(c.Writer, c.Request, "/auth/sign-in", http.StatusSeeOther)
		return
	}
	commentText := c.Request.FormValue("comment")
	if strings.TrimSpace(commentText) == "" {
		http.Redirect(c.Writer, c.Request, "/post/"+postID, http.StatusSeeOther)
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
		h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	c.Redirect(http.StatusFound, "/post/"+postID)
}
