package service

import (
	"database/sql"
	"errors"
	"log"

	"forum/internal/repository"
	"forum/models"
)

type DislikeService struct {
	repo repository.Dislike
}

func NewDislikeService(repo repository.Dislike) *DislikeService {
	return &DislikeService{
		repo: repo,
	}
}

func (s *DislikeService) SetPostDislike(dislike models.DisLike) error {
	if err := s.repo.CheckPostDislike(dislike.UserID, dislike.PostID); err == nil {
		if err = s.repo.DeletePostDislike(dislike.UserID, dislike.PostID); err != nil {
			log.Printf("service: %s", err)
			return err
		}
	} else if errors.Is(err, sql.ErrNoRows) {
		if err = s.repo.CheckPostLike(dislike.UserID, dislike.PostID); err == nil {
			if err = s.repo.DeletePostLike(dislike.UserID, dislike.PostID); err != nil {
				log.Printf("service: %s", err)
				return err
			}
			if err = s.repo.SetPostDislike(dislike); err != nil {
				log.Printf("service: %s", err)
				return err
			}
		} else if errors.Is(err, sql.ErrNoRows) {
			if err = s.repo.SetPostDislike(dislike); err != nil {
				log.Printf("service: %s", err)
				return err
			}
		}
	}
	if err := s.repo.UpdatePostVote(dislike.PostID); err != nil {
		log.Printf("service: %s", err)
		return err
	}
	return nil
}

func (s *DislikeService) SetCommentDislike(dislike models.DisLike) error {
	if err := s.repo.CheckCommentDislike(dislike.UserID, dislike.CommentID); err == nil {
		if err = s.repo.DeleteCommentDislike(dislike.UserID, dislike.CommentID); err != nil {
			log.Printf("service: %s", err)
			return err
		}
	} else if errors.Is(err, sql.ErrNoRows) {
		if err = s.repo.CheckCommentLike(dislike.UserID, dislike.CommentID); err == nil {
			if err = s.repo.DeleteCommentLike(dislike.UserID, dislike.CommentID); err != nil {
				log.Printf("service: %s", err)
				return err
			}
			if err = s.repo.SetCommentDislike(dislike); err != nil {
				log.Printf("service: %s", err)
				return err
			}
		} else if errors.Is(err, sql.ErrNoRows) {
			if err = s.repo.SetCommentDislike(dislike); err != nil {
				log.Printf("service: %s", err)
				return err
			}
		}
	}
	if err := s.repo.UpdateCommentVote(dislike.CommentID); err != nil {
		log.Printf("service: %s", err)
		return err
	}
	return nil
}
