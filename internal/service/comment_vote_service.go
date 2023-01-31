package service

import (
	"github.com/muratovdias/my-forum.2.0/internal/repository"
	"github.com/muratovdias/my-forum.2.0/models"
)

type CommentVoteService struct {
	commentRepo repository.CommentVote
}

func NewCommentVoteService(commentRepo repository.CommentVote) *CommentVoteService {
	return &CommentVoteService{
		commentRepo: commentRepo,
	}
}

func (c CommentVoteService) MaipulationCommentVote(models.UserCommentVote) error {
	return nil
}
