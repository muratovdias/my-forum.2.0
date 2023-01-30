package repository

import "gorm.io/gorm"

type CommentVoteRepo struct {
	db *gorm.DB
}

func NewCommentVoteRepo(db *gorm.DB) *CommentVoteRepo {
	return &CommentVoteRepo{
		db: db,
	}
}
