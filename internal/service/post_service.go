package service

import (
	"fmt"
	"log"
	"strconv"

	"forum/internal/repository"
	"forum/models"
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
	return s.repo.MyFavourites(id)
}
