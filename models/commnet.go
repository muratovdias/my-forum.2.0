package models

type Comment struct {
	ID       int
	UserID   int
	PostID   int
	Likes    int
	Dislikes int
	Text     string
	Author   string
	Date     string
}
