package repository

import (
	"time"

	"github.com/muratovdias/my-forum.2.0/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(models.User) error
	GetUserByEmail(string) (models.User, error)
	GetUserByUsername(string) (models.User, error)
	GetUserByToken(string) (models.User, error)
	SaveToken(string, string, time.Time) error
	DeleteToken(string) error
}

type Post interface {
	CreatePost(*models.Post) error
	GetAllPost() (*[]models.Post, error)
	GetPostByCategory(string) (*[]models.Post, error)
	GetPostByID(string) (*models.Post, error)
	MyPosts(string) (*[]models.Post, error)
	MyFavourites(int) (*[]models.Post, error)
}

type Comment interface {
	CreateComment(models.Comment) error
	CheckCommentExists(string) error
	GetCommentByPostID(int) (*[]models.Comment, error)
}

type UserPostVote interface{}

type UserCommentVote interface{}

type Repository struct {
	Authorization
	Post
	Comment
	UserPostVote
	UserCommentVote
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
		Post:          NewPostRepo(db),
		Comment:       NewCommentRepo(db),
	}
}
