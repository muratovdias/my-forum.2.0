package handler

import (
	"errors"
	"fmt"
	"net/http"

	"forum/internal/service"
	"forum/models"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := h.templExecute(w, "./ui/sign-up.html", nil); err != nil {
			return
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.ErrorPage(w, http.StatusInternalServerError, "something went wrong")
			return
		}
		username, ok1 := r.Form["username"]
		email, ok2 := r.Form["email"]
		password, ok3 := r.Form["password"]
		if !ok1 || !ok2 || !ok3 {
			w.WriteHeader(http.StatusBadRequest)
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
			h.ErrorPage(w, http.StatusBadRequest, fmt.Sprintf("%s\n", err))
			return
		} else if err != nil {
			h.ErrorPage(w, http.StatusInternalServerError, fmt.Sprintf("%s\n", err))
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := h.templExecute(w, "./ui/sign-in.html", nil); err != nil {
			return
		}
	case http.MethodPost:
		email := r.FormValue("email")
		password := r.FormValue("password")
		user, err := h.services.Authorization.GenerateToken(email, password)
		var status int
		if err == service.ErrMail || err == service.ErrPassword {
			if err == service.ErrMail {
				status = http.StatusUnauthorized
			} else {
				status = http.StatusBadRequest
			}
			h.ErrorPage(w, status, fmt.Sprintf("%s", err))
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:  "session_token",
			Value: user.Token,
			Path:  "/",
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) logOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.ErrorPage(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		fmt.Println("method: log-out")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: "",
		Path:  "/",
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
