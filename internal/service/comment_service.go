package service

import (
	"forum/internal/repository"
	"forum/models"
)

type CommentService struct {
	repo repository.Comment
}

func NewCommetService(repo repository.Comment) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (s *CommentService) CreateComment(commnet models.Comment) error {
	return s.repo.CreateComment(commnet)
}

func (s *CommentService) GetCommentByPostID(id int) (*[]models.Comment, error) {
	return s.repo.GetCommentByPostID(id)
}
