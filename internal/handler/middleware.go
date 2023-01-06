package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"forum/models"
)

const ctxUserKey ctxKey = iota

type ctxKey int8

func (h *Handler) userIdentity(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Execute middleware 1")
		var user models.User
		c, err := r.Cookie("session_token")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxUserKey, models.User{})))
				return
			case errors.Is(err, c.Valid()):
				h.ErrorPage(w, http.StatusBadRequest, "invalid cookie value")
			}
			h.ErrorPage(w, http.StatusBadRequest, "failed to get cookie")
			return
		}
		user, err = h.services.Authorization.GetUserByToken(c.Value)
		if err != nil {
			handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxUserKey, models.User{})))
			return
		}
		if user.TokenDuration.Before(time.Now()) {
			if err := h.services.DeleteToken(user.Token); err != nil {
				h.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				return
			}
			http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
			return
		}
		handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxUserKey, user)))
	}
}

func (h *Handler) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Method: %s URL: %s Time: %s", r.Method, r.RequestURI, time.Since(start))
	})
}
