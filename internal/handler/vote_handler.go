package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muratovdias/my-forum.2.0/models"
)

var (
	postID    string
	commentID string
	url       string
	id        int
	err       error
)

func (h *Handler) likePost(c *gin.Context) {
	userInterface, _ = c.Get("user")
	user = userInterface.(models.User)
	if user == (models.User{}) {
		c.Redirect(http.StatusFound, "/auth/sign-in")
		return
	}
	postID1 := c.Request.FormValue("like1")
	fmt.Println(postID1)
	postID2 := c.Request.FormValue("like2")

	if postID1 == "" {
		postID = postID2
		url = "/post/" + postID
	} else {
		postID = postID1
		url = "/"
	}
	_, err = h.services.Post.GetPostByID(postID)
	if err != nil {
		h.ErrorPage(c.Writer, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	id, _ = strconv.Atoi(postID)
	if err != nil {
		h.ErrorPage(c.Writer, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	like := models.UserPostVote{
		UserID: user.ID,
		PostID: id,
		Vote:   true,
	}
	if err = h.services.PostVote.ManipulationPostVote(like); err != nil {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	c.Redirect(http.StatusSeeOther, url)
}

func (h *Handler) dislikePost(c *gin.Context) {
	userInterface, _ = c.Get("user")
	user = userInterface.(models.User)
	if user == (models.User{}) {
		c.Redirect(http.StatusFound, "/auth/sign-in")
		return
	}
	postID1 := c.Request.FormValue("dislike1")
	fmt.Println(postID1)
	postID2 := c.Request.FormValue("dislike2")

	if postID1 == "" {
		postID = postID2
		url = "/post/" + postID
	} else {
		postID = postID1
		url = "/"
	}
	_, err = h.services.Post.GetPostByID(postID)
	if err != nil {
		h.ErrorPage(c.Writer, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	id, _ = strconv.Atoi(postID)
	if err != nil {
		h.ErrorPage(c.Writer, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	dislike := models.UserPostVote{
		UserID: user.ID,
		PostID: id,
		Vote:   false,
	}
	if err = h.services.PostVote.ManipulationPostVote(dislike); err != nil {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	c.Redirect(http.StatusSeeOther, url)
}

func (h *Handler) likeComment(c *gin.Context) {
	userInterface, _ = c.Get("user")
	user = userInterface.(models.User)
	if user == (models.User{}) {
		c.Redirect(http.StatusFound, "/auth/sign-in")
		return
	}
	commentID = c.Param("id")
	if err := h.services.Comment.CheckCommentExists(commentID); err != nil {
		h.ErrorPage(c.Writer, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	id, _ = strconv.Atoi(commentID)
	if err != nil {
		h.ErrorPage(c.Writer, http.StatusNotFound, "Sorry, but no pages were found for your request =(")
		return
	}
	like := models.UserCommentVote{
		UserID:    user.ID,
		CommentID: id,
		Vote:      true,
	}
	if err = h.services.CommentVote.MaipulationCommentVote(like); err != nil {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	c.Redirect(http.StatusSeeOther, "/post/"+postID)
}

func (h *Handler) dislikeComment(c *gin.Context) {
	userInterface, _ = c.Get("user")
	user = userInterface.(models.User)
	if user == (models.User{}) {
		c.Redirect(http.StatusFound, "/auth/sign-in")
		return
	}
	commentID = c.Param("id")
	if err := h.services.Comment.CheckCommentExists(commentID); err != nil {
		h.ErrorPage(c.Writer, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	id, err = strconv.Atoi(commentID)
	if err != nil {
		log.Printf("id = %d", id)
		h.ErrorPage(c.Writer, http.StatusNotFound, "Sorry, but no pages were found for your request =(")
		return
	}
	dislike := models.UserCommentVote{
		UserID:    user.ID,
		CommentID: id,
		Vote:      false,
	}
	if err = h.services.CommentVote.MaipulationCommentVote(dislike); err != nil {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	c.Redirect(http.StatusSeeOther, "/post/"+postID)
}
