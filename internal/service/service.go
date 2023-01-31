package service

import (
	"github.com/muratovdias/my-forum.2.0/internal/repository"
	"github.com/muratovdias/my-forum.2.0/models"
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
	CheckCommentExists(string) error
	GetCommentByPostID(int) (*[]models.Comment, error)
}

type PostVote interface {
	ManipulationPostVote(models.UserPostVote) error
}

type CommentVote interface {
	MaipulationCommentVote(models.UserCommentVote) error
}

type Service struct {
	Authorization
	Post
	Comment
	PostVote
	CommentVote
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Post:          NewPostService(repo.Post),
		Comment:       NewCommetService(repo.Comment),
		PostVote:      NewPostVoteService(repo.PostVote),
		CommentVote:   NewCommentVoteService(repo.CommentVote),
	}
}
