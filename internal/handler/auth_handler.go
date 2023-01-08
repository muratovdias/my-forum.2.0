package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muratovdias/my-forum.2.0/models"

	"github.com/muratovdias/my-forum.2.0/internal/service"
)

func (h *Handler) signUpGET(c *gin.Context) {
	if err := h.templExecute(c.Writer, "./ui/sign-up.html", nil); err != nil {
		return
	}
}

func (h *Handler) signUpPOST(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, "something went wrong")
		return
	}
	username, ok1 := c.Request.Form["username"]
	email, ok2 := c.Request.Form["email"]
	password, ok3 := c.Request.Form["password"]
	if !ok1 || !ok2 || !ok3 {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	user := models.User{
		Email:    email[0],
		Username: username[0],
		Password: password[0],
	}
	err := h.services.Authorization.CreateUser(user)
	// Handle errors
	if errors.Is(err, service.ErrInvalidEmail) || errors.Is(err, service.ErrInvalidUsername) || errors.Is(err, service.ErrInvalidPassword) {
		h.ErrorPage(c.Writer, http.StatusBadRequest, fmt.Sprintf("%s\n", err))
		return
	} else if err != nil {
		h.ErrorPage(c.Writer, http.StatusInternalServerError, fmt.Sprintf("%s\n", err))
		return
	}
	http.Redirect(c.Writer, c.Request, "/", http.StatusSeeOther)
}

func (h *Handler) signInGET(c *gin.Context) {
	if err := h.templExecute(c.Writer, "./ui/sign-in.html", nil); err != nil {
		return
	}
}

func (h *Handler) signInPOST(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	email, ok1 := c.Request.Form["email"]
	password, ok2 := c.Request.Form["password"]
	if !ok1 || !ok2 {
		h.ErrorPage(c.Writer, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	user, err := h.services.Authorization.GenerateToken(email[0], password[0])
	var status int
	if err == service.ErrMail || err == service.ErrPassword {
		if err == service.ErrMail {
			status = http.StatusUnauthorized
		} else {
			status = http.StatusBadRequest
		}
		h.ErrorPage(c.Writer, status, fmt.Sprintf("%s", err))
		return
	}
	c.SetCookie("session_token", user.Token, 3600, "/", "localhost", false, true)

	http.Redirect(c.Writer, c.Request, "/", http.StatusSeeOther)
}

func (h *Handler) logOut(c *gin.Context) {
	c.SetCookie("session_token", "", 3600, "/", "localhost", true, false)
	http.Redirect(c.Writer, c.Request, "/", http.StatusSeeOther)
}
