package models

type Post struct {
	ID       int
	AuthorID int
	Title    string
	Category string
	Content  string
	Author   string
	Date     string
	Likes    int
	Dislikes int
}
