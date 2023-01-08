package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/muratovdias/my-forum.2.0/models"
)

func (h *Handler) home(c *gin.Context) {
	if c.Request.URL.Path != "/" {
		h.ErrorPage(c.Writer, http.StatusNotFound, "Sorry, but no pages were found for your request =(")
		return
	}
	var user models.User
	var posts **[]models.Post
	var err error
	userInterface, _ = c.Get("user")
	user, ok = userInterface.(models.User)
	if !ok {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	if err := c.Request.ParseForm(); err != nil {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, "something went wrong")
		return
	}
	category := strings.Join(c.Request.Form["category"], " ")
	if category != "" {
		posts, err = h.services.GetPostByCategory(category)
		if err != nil {
			h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	} else {
		posts, err = h.services.GetAllPost()
		if err != nil {
			h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	}
	pageData := models.PageData{
		Username: user.Username,
		Posts:    **posts,
	}
	if err := h.templExecute(c.Writer, "./ui/index.html", pageData); err != nil {
		return
	}
}
