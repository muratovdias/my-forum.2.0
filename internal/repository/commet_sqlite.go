package repository

import (
	"fmt"

	"github.com/muratovdias/my-forum.2.0/models"
	"gorm.io/gorm"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (r *CommentRepo) CreateComment(comment models.Comment) error {
	row := r.db.Create(&comment)
	if row.Error != nil {
		fmt.Printf("repo: %s\n", err)
		return fmt.Errorf(path+"create comment: %w", err)
	}
	return nil
}

func (r *CommentRepo) CheckCommentExists(id string) error {
	row := r.db.First(&models.Comment{}, "id=?", id)
	if row.Error != nil {
		fmt.Println("Check comment: ", row.Error)
		return row.Error
	}
	return nil
}

func (r *CommentRepo) GetCommentByPostID(id int) (*[]models.Comment, error) {
	var comments []models.Comment
	result := r.db.Where("user_id=?", id).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comments, nil
}
