package models

type UserCommentVote struct {
	ID        int
	CommentID int
	UserID    int
	Vote      bool
}
