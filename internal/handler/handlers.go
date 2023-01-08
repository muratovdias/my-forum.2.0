package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muratovdias/my-forum.2.0/internal/service"
	"github.com/muratovdias/my-forum.2.0/models"
)

var (
	userInterface any
	user          models.User
	ok            bool
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
		auth.GET("/sign-up", h.signUpGET)
		auth.POST("/sign-up", h.signUpPOST)
		auth.GET("/sign-in", h.signInGET)
		auth.POST("/sign-in", h.signInPOST)
	}

	main := router.Group("/")

	main.Use(h.userIdentity())
	{
		main.GET("/", h.home)
		main.GET("/log-out", h.logOut)
		main.GET("/my-posts", h.myPosts)
		main.GET("/my-favourites", h.myFavourites)
		main.POST("/like-post", h.likePost)
		main.POST("/dislike-post", h.dislikePost)
		post := main.Group("/post")
		{
			post.GET("/create", h.createPostGET)
			post.POST("/create", h.createPostPOST)

			post.GET("/:id", h.postGET)
			post.POST("/:id/comment", h.postPOST)

			comment := post.Group("/comment")
			{
				comment.POST("/:id/like", h.likeComment)
				comment.POST("/:id/dislike", h.dislikeComment)
			}
		}
	}
	router.Static("/ui/css", "./ui/css")

	return router
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
