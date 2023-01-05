package repository

import (
	"database/sql"
	"time"

	"forum/models"
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
	GetCommentByPostID(int) (*[]models.Comment, error)
}

type Like interface {
	SetPostLike(models.Like) error
	SetCommentLike(models.Like) error
	LikeDislike
}

type Dislike interface {
	SetPostDislike(models.DisLike) error
	SetCommentDislike(models.DisLike) error
	LikeDislike
}

type LikeDislike interface {
	// Post//
	CheckPostDislike(int, int) error
	CheckPostLike(int, int) error
	DeletePostDislike(int, int) error
	DeletePostLike(int, int) error
	UpdatePostVote(int) error
	// Comment//
	CheckCommentDislike(int, int) error
	CheckCommentLike(int, int) error
	DeleteCommentDislike(int, int) error
	DeleteCommentLike(int, int) error
	UpdateCommentVote(int) error
}

type Repository struct {
	Authorization
	Post
	Comment
	Like
	Dislike
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
		Post:          NewPostRepo(db),
		Comment:       NewCommentRepo(db),
		Like:          NewLikeRepo(db),
		Dislike:       NewDislikeRepo(db),
	}
}
