package repository

import (
	"fmt"

	"github.com/muratovdias/my-forum.2.0/models"
	"gorm.io/gorm"
)

type CommentVoteRepo struct {
	db *gorm.DB
}

func NewCommentVoteRepo(db *gorm.DB) *CommentVoteRepo {
	return &CommentVoteRepo{
		db: db,
	}
}

func (v *CommentVoteRepo) SetVoteComment(c models.UserCommentVote) error {
	row := v.db.Create(&c)
	if row.Error != nil {
		return fmt.Errorf("set vote post: %w", row.Error)
	}
	return nil
}

func (v *CommentVoteRepo) GetCommentVoteByID(userID, commentID int) (models.UserCommentVote, error) {
	var userVote models.UserCommentVote
	row := v.db.Where("user_id = ? AND comment_id = ?", userID, commentID).Find(&userVote)
	if row.Error != nil {
		return models.UserCommentVote{}, fmt.Errorf("get vote by ID: %w", row.Error)
	}
	return userVote, nil
}

func (v *CommentVoteRepo) ManipulationCommentVote(c models.UserCommentVote) error {
	userVote, err := v.GetCommentVoteByID(c.UserID, c.CommentID)
	if err != nil {
		return fmt.Errorf("getVoteByID:%w", err)
	}
	if (userVote == models.UserCommentVote{}) {
		err = v.SetVoteComment(c)
		if err != nil {
			return fmt.Errorf("set comment vote:%w", err)
		}
	} else {
		if userVote.Vote == c.Vote {
			err := v.DeleteCommentVote(c)
			if err != nil {
				return fmt.Errorf("delete: %w", err)
			}
		} else {
			if userVote.Vote {
				v.db.Model(&models.UserCommentVote{}).Where("user_id=? AND comment_id=?", c.UserID, c.CommentID).Update("vote", false)
			} else {
				v.db.Model(&models.UserCommentVote{}).Where("user_id=? AND comment_id=?", c.UserID, c.CommentID).Update("vote", true)
			}
		}
	}

	return nil
}

func (v *CommentVoteRepo) DeleteCommentVote(c models.UserCommentVote) error {
	row := v.db.Where("user_id=? AND comment_id=?", c.UserID, c.CommentID).Delete(&models.UserCommentVote{})
	if row.Error != nil {
		return fmt.Errorf("delete post vote:%w", row.Error)
	}
	return nil
}
