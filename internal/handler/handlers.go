package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muratovdias/my-forum.2.0/internal/service"
	"honnef.co/go/tools/go/callgraph/static"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.GET("/sign-up")
		auth.GET("/sign-in")

		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}

	main := router.Group("/")
	{
		main.GET("log-out")
		main.GET("my-posts")
		main.GET("category")
		main.GET("my-favourites")
		main.POST("like")
		main.POST("dislike")
	}

	post := router.Group("/post")
	{
		post.GET("/create")
		post.GET("/:id")

		post.POST("/create")
		post.POST("/like-comment")
		post.POST("/dislike-comment")
	}

	router.Use(static.Serve("/css", static.LocalFile("ui/css", true)))

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
