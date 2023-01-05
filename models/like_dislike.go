package models

type Like struct {
	ID        int
	PostID    int
	UserID    int
	CommentID int
	Active    string
}

type DisLike struct {
	ID        int
	PostID    int
	UserID    int
	CommentID int
	Active    string
}
