package service

import (
	"github.com/muratovdias/my-forum.2.0/internal/repository"
	"github.com/muratovdias/my-forum.2.0/models"
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

func (s *CommentService) CheckCommentExists(id string) error {
	return nil
}
