package service

import (
	"forum/internal/repository"
	"forum/models"
)

type Authorization interface {
	CreateUser(models.User) error
	GenerateToken(string, string) (models.User, error)
	GetUserByToken(string) (models.User, error)
	DeleteToken(string) error
}

type Post interface {
	CreatePost(*models.Post) error
	GetAllPost() (**[]models.Post, error)
	GetPostByCategory(string) (**[]models.Post, error)
	MyPosts(string) (*[]models.Post, error)
	MyFavourites(int) (*[]models.Post, error)
	GetPostByID(string) (*models.Post, error)
}

type Comment interface {
	CreateComment(models.Comment) error
	GetCommentByPostID(int) (*[]models.Comment, error)
}

type Like interface {
	SetPostLike(models.Like) error
	SetCommentLike(models.Like) error
}

type Dislike interface {
	SetPostDislike(models.DisLike) error
	SetCommentDislike(models.DisLike) error
}

type Service struct {
	Authorization
	Post
	Comment
	Like
	Dislike
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Post:          NewPostService(repo.Post),
		Comment:       NewCommetService(repo.Comment),
		Like:          NewLikeService(repo.Like),
		Dislike:       NewDislikeService(repo.Dislike),
	}
}
