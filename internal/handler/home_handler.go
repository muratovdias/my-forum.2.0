package handler

import (
	"net/http"
	"strings"

	"forum/models"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.ErrorPage(w, http.StatusNotFound, "Sorry, but no pages were found for your request =(")
		return
	}
	var user models.User
	var posts **[]models.Post
	var err error
	user = r.Context().Value(ctxUserKey).(models.User)
	switch r.Method {
	case http.MethodGet:
		r.ParseForm()
		category := strings.Join(r.Form["category"], " ")
		if category != "" {
			posts, err = h.services.GetPostByCategory(category)
			if err != nil {
				h.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				return
			}
		} else {
			posts, err = h.services.GetAllPost()
			if err != nil {
				h.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				return
			}
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	pageData := models.PageData{
		Username: user.Username,
		Posts:    **posts,
	}
	if err := h.templExecute(w, "./ui/index.html", pageData); err != nil {
		return
	}
}
