package repository

import (
	"fmt"
	"log"

	"github.com/muratovdias/my-forum.2.0/models"
	"gorm.io/gorm"
)

var (
	rows *gorm.Rows
	err  error
)

type PostRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (r *PostRepo) CreatePost(post *models.Post) error {
	result := r.db.Create(&post)
	if result.Error != nil {
		log.Printf("error create post: %s", result.Error)
		return result.Error
	}
	return nil
}

func (r *PostRepo) GetAllPost() (*[]models.Post, error) {
	var posts *[]models.Post
	result := r.db.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (r *PostRepo) GetPostByCategory(category string) (*[]models.Post, error) {
	var posts *[]models.Post
	row := r.db.Where("category LIKE ?", "%"+category+"%").Order("id DESC").Find(&posts)
	if row.Error != nil {
		return nil, row.Error
	}
	return posts, nil
}

func (r *PostRepo) MyPosts(id string) (*[]models.Post, error) {
	var posts []models.Post
	row := r.db.Where("author_id=?", id).Find(&posts)
	if row.Error != nil {
		return nil, fmt.Errorf("me posts: %w", row.Error)
	}
	return &posts, nil
}

func (r *PostRepo) MyFavourites(id int) (*[]models.Post, error) {
	var ids []int
	row := r.db.Table("user_post_votes").Select("id").Where("user_id=? AND vote=true", id).Find(&ids)
	if row.Error != nil {
		return nil, row.Error
	}
	var posts []models.Post
	row = r.db.Find(&posts, ids)
	if row.Error != nil {
		return nil, fmt.Errorf("my favourites: %w", row.Error)
	}
	return &posts, nil
}

func (r *PostRepo) GetPostByID(id string) (*models.Post, error) {
	var post *models.Post
	row := r.db.Where("id=?", id).Find(&post)
	if row.Error != nil {
		return nil, fmt.Errorf("get post bi ID: %w", row.Error)
	}
	return post, nil
}
