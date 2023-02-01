package repository

import (
	"fmt"

	"github.com/muratovdias/my-forum.2.0/models"
	"gorm.io/gorm"
)

type PostVoteRepo struct {
	db *gorm.DB
}

func NewPostVoteRepo(db *gorm.DB) *PostVoteRepo {
	return &PostVoteRepo{
		db: db,
	}
}

func (v *PostVoteRepo) SetVotePost(p models.UserPostVote) error {
	row := v.db.Create(&p)
	if row.Error != nil {
		return fmt.Errorf("set vote post: %w", row.Error)
	}
	return nil
}

func (v *PostVoteRepo) GetPostVoteByID(userID, postID int) (models.UserPostVote, error) {
	var userVote models.UserPostVote
	row := v.db.Where("user_id = ? AND post_id = ?", userID, postID).Find(&userVote)
	if row.Error != nil {
		return models.UserPostVote{}, fmt.Errorf("get vote by ID: %w", row.Error)
	}
	return userVote, nil
}

func (v *PostVoteRepo) ManipulationPostVote(p models.UserPostVote) error {
	userVote, err := v.GetPostVoteByID(p.UserID, p.PostID)
	if err != nil {
		return fmt.Errorf("getVoteByID:%w", err)
	}
	if (userVote == models.UserPostVote{}) {
		err := v.db.Transaction(func(tx *gorm.DB) error {
			err = v.SetVotePost(p)
			if err != nil {
				return fmt.Errorf("set post vote:%w", err)
			}
			if p.Vote {
				tx.Table("posts").Where("author_id=? AND id=?", p.UserID, p.PostID).UpdateColumn("likes", gorm.Expr("likes+ ?", 1))
			}
			return nil
		})
		if err != nil {
			return err
		}
		// если такая запись есть, опираясь на голос обновляем запись. Удаляем или же меняем голос на противоположный
	} else {
		if userVote.Vote == p.Vote {
			err := v.DeletePostVote(p)
			if err != nil {
				return fmt.Errorf("delete: %w", err)
			}
		} else {
			if userVote.Vote {
				v.db.Model(&models.UserPostVote{}).Where("user_id=? AND post_id=?", p.UserID, p.PostID).Update("vote", false)
			} else {
				v.db.Model(&models.UserPostVote{}).Where("user_id=? AND post_id=?", p.UserID, p.PostID).Update("vote", true)
			}
		}
	}

	return nil
}

func (v *PostVoteRepo) DeletePostVote(p models.UserPostVote) error {
	row := v.db.Where("user_id=? AND post_id=?", p.UserID, p.PostID).Delete(&models.UserPostVote{})
	if row.Error != nil {
		return fmt.Errorf("delete post vote:%w", row.Error)
	}
	return nil
}
