package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muratovdias/my-forum.2.0/models"
)

func (h *Handler) userIdentity() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		cookie, err := c.Cookie("session_token")
		c.Set("user", user)
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				c.Next()
				return
			case cookie == "":
				c.Next()
				return
			}
			h.ErrorPage(c.Writer, http.StatusBadRequest, "failed to get cookie")
			return
		}
		user, err = h.services.Authorization.GetUserByToken(cookie)
		if err != nil {
			return
		}
		if user.TokenDuration.Before(time.Now()) {
			if err := h.services.DeleteToken(user.Token); err != nil {
				h.ErrorPage(c.Writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				return
			}
			http.Redirect(c.Writer, c.Request, "/auth/sign-in", http.StatusSeeOther)
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

// func (h *Handler) Logging(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		next.ServeHTTP(w, r)
// 		log.Printf("Method: %s URL: %s Time: %s", r.Method, r.RequestURI, time.Since(start))
// 	})
// }
