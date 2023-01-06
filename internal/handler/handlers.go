package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"forum/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.userIdentity(h.home))
	mux.HandleFunc("/auth/sign-up", h.signUp)
	mux.HandleFunc("/auth/sign-in", h.signIn)
	mux.HandleFunc("/log-out", h.logOut)
	mux.HandleFunc("/post/create", h.userIdentity(h.createPost))
	mux.HandleFunc("/my-posts", h.userIdentity(h.myPosts))
	mux.HandleFunc("/my-favourites", h.userIdentity(h.myFavourites))
	mux.HandleFunc("/post/", h.userIdentity(h.post))
	mux.HandleFunc("/like-post", h.userIdentity(h.likePost))
	mux.HandleFunc("/dislike-post", h.userIdentity(h.dislikePost))
	mux.HandleFunc("/like-comment", h.userIdentity(h.likeComment))
	mux.HandleFunc("/dislike-comment", h.userIdentity(h.dislikeComment))
	mux.Handle("/ui/css/", http.StripPrefix("/ui/css/", http.FileServer(http.Dir("./ui/css/"))))
	// handler := h.Logging(mux)
	return mux
}

func (h *Handler) templExecute(w http.ResponseWriter, path string, data interface{}) error {
	templ, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println("ParseFiles()")
		h.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return err
	}
	if err = templ.Execute(w, data); err != nil {
		fmt.Println("templExecute()")
		h.ErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return err
	}
	return nil
}
