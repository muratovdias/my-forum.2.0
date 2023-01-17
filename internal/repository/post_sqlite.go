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
	rows, err = r.db.Query(`SELECT * FROM post WHERE category LIKE '%` + category + `%'` + `ORDER BY id DESC`)
	if err != nil {
		log.Printf("by category: %s", err)
		return nil, fmt.Errorf(path+"get all post: %w", err)
	}
	var posts []models.Post
	for rows.Next() {
		post := new(models.Post)
		if err = rows.Scan(&post.ID, &post.AuthorID, &post.Likes, &post.Dislikes, &post.Title, &post.Category, &post.Content, &post.Author, &post.Date); err != nil {
			log.Printf("by category scan: %s", err)
			return nil, fmt.Errorf(path+"get all post, scan: %w", err)
		}
		posts = append(posts, *post)
	}
	return &posts, nil
}

func (r *PostRepo) MyPosts(id string) (*[]models.Post, error) {
	rows, err = r.db.Query(`SELECT * FROM post WHERE author_id=` + id + ` ORDER BY id DESC`)
	if err != nil {
		return nil, fmt.Errorf(path+"get my post: %w", err)
	}
	var posts []models.Post
	for rows.Next() {
		post := new(models.Post)
		if err = rows.Scan(&post.ID, &post.AuthorID, &post.Likes, &post.Dislikes, &post.Title, &post.Category, &post.Content, &post.Author, &post.Date); err != nil {
			return nil, fmt.Errorf(path+"get my posts: scan: %w", err)
		}
		posts = append(posts, *post)
	}
	return &posts, nil
}

func (r *PostRepo) MyFavourites(id int) (*[]models.Post, error) {
	var posts []models.Post
	query = `SELECT post_id FROM like WHERE user_id=$1 AND post_id != 0 AND active=1 ORDER BY id DESC`
	rows, err = r.db.Query(query, id)
	if err != nil {
		log.Printf("my favourites query: %s\n", err)
		return nil, fmt.Errorf(path+"select my favourites: %w", err)
	}
	var postsID []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf(path+"my favourites: scan post_id: %w", err)
		}
		postsID = append(postsID, id)
	}
	for _, id := range postsID {
		post := new(models.Post)
		query = `SELECT * FROM post WHERE id=$1 ORDER BY id DESC`
		row := r.db.QueryRow(query, id)
		if err = row.Scan(&post.ID, &post.AuthorID, &post.Likes, &post.Dislikes, &post.Title, &post.Category, &post.Content, &post.Author, &post.Date); err != nil {
			log.Printf("my favourites: scan post: %s", err)
			return nil, fmt.Errorf(path+"my favourites: scan post: %w", err)
		}
		posts = append(posts, *post)
	}
	return &posts, nil
}

func (r *PostRepo) GetPostByID(id string) (*models.Post, error) {
	rows, err := r.db.Query(`SELECT * FROM post WHERE id=` + id)
	if err != nil {
		return nil, fmt.Errorf(path+"get post by id: %w", err)
	}
	var post models.Post
	for rows.Next() {
		if err := rows.Scan(&post.ID, &post.AuthorID, &post.Likes, &post.Dislikes, &post.Title, &post.Category, &post.Content, &post.Author, &post.Date); err != nil {
			return nil, fmt.Errorf(path+"get post by id: scan: %w", err)
		}
	}
	return &post, nil
}
