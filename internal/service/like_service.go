package service

import (
	"database/sql"
	"errors"
	"log"

	"forum/internal/repository"
	"forum/models"
)

type LikeService struct {
	repo repository.Like
}

func NewLikeService(repo repository.Like) *LikeService {
	return &LikeService{
		repo: repo,
	}
}

func (s *LikeService) SetPostLike(like models.Like) error {
	if err := s.repo.CheckPostLike(like.UserID, like.PostID); err == nil {
		if err = s.repo.DeletePostLike(like.UserID, like.PostID); err != nil {
			log.Printf("service: %s", err)
			return err
		}
	} else if errors.Is(err, sql.ErrNoRows) {
		if err = s.repo.CheckPostDislike(like.UserID, like.PostID); err == nil {
			if err = s.repo.DeletePostDislike(like.UserID, like.PostID); err != nil {
				log.Printf("service: %s", err)
				return err
			}
			if err = s.repo.SetPostLike(like); err != nil {
				log.Printf("service: %s", err)
				return err
			}
		} else if errors.Is(err, sql.ErrNoRows) {
			if err = s.repo.SetPostLike(like); err != nil {
				log.Printf("service: %s", err)
				return err
			}
		}
	}
	if err := s.repo.UpdatePostVote(like.PostID); err != nil {
		log.Printf("service: %s", err)
		return err
	}
	return nil
}

func (s *LikeService) SetCommentLike(like models.Like) error {
	if err := s.repo.CheckCommentLike(like.UserID, like.CommentID); err == nil {
		if err = s.repo.DeleteCommentLike(like.UserID, like.CommentID); err != nil {
			log.Printf("service: %s", err)
			return err
		}
	} else if errors.Is(err, sql.ErrNoRows) {
		if err = s.repo.CheckCommentDislike(like.UserID, like.CommentID); err == nil {
			if err = s.repo.DeleteCommentDislike(like.UserID, like.CommentID); err != nil {
				log.Printf("service: %s", err)
				return err
			}
			if err = s.repo.SetCommentLike(like); err != nil {
				log.Printf("service: %s", err)
				return err
			}
		} else if errors.Is(err, sql.ErrNoRows) {
			if err = s.repo.SetCommentLike(like); err != nil {
				log.Printf("service: %s", err)
				return err
			}
		}
	}
	if err := s.repo.UpdateCommentVote(like.CommentID); err != nil {
		log.Printf("service: %s", err)
		return err
	}
	return nil
}
