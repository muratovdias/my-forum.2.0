package models

type Comment struct {
	ID       int
	UserID   int
	PostID   int
	Text     string
	Author   string
	Date     string
}
