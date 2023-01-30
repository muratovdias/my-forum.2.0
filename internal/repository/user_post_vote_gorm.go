package repository

import "gorm.io/gorm"

type PostVoteRepo struct {
	db *gorm.DB
}

func NewPostVoteRepo(db *gorm.DB) *PostVoteRepo {
	return &PostVoteRepo{
		db: db,
	}
}


