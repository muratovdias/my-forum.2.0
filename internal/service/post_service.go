package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/muratovdias/my-forum.2.0/internal/repository"
	"github.com/muratovdias/my-forum.2.0/models"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) CreatePost(post *models.Post) error {
	return s.repo.CreatePost(post)
}

func (s *PostService) GetAllPost() (**[]models.Post, error) {
	posts, err := s.repo.GetAllPost()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &posts, nil
}

func (s *PostService) GetPostByCategory(category string) (**[]models.Post, error) {
	posts, err := s.repo.GetPostByCategory(category)
	if err != nil {
		return nil, err
	}
	return &posts, nil
}

func (s *PostService) MyPosts(id string) (*[]models.Post, error) {
	return s.repo.MyPosts(id)
}

func (s *PostService) GetPostByID(id string) (*models.Post, error) {
	_, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	post, err := s.repo.GetPostByID(id)
	if err != nil {
		fmt.Printf("service: %s", err)
		return nil, err
	}
	return post, nil
}

func (s *PostService) MyFavourites(id int) (*[]models.Post, error) {
	_, err := s.repo.MyFavourites(id)
	if errors.Is(err, sql.ErrNoRows) {
		return &[]models.Post{}, nil
	} else {
		return s.repo.MyFavourites(id)
	}
}
