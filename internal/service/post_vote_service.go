package service

import (
	"github.com/muratovdias/my-forum.2.0/internal/repository"
	"github.com/muratovdias/my-forum.2.0/models"
)

type PostVoteService struct {
	postRepo repository.PostVote
}

func NewPostVoteService(postRepo repository.PostVote) *PostVoteService {
	return &PostVoteService{
		postRepo: postRepo,
	}
}

func (p PostVoteService) ManipulationPostVote(models.UserPostVote) error {
	return nil
}
