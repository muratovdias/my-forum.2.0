package models

type UserPostVote struct {
	ID     int
	PostID int
	UserID int
	Vote   bool
}
