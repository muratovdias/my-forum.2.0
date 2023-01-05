package models

type Post struct {
	ID       int
	AuthorID int
	Likes    int
	Dislikes int
	Title    string
	Category string
	Content  string
	Author   string
	Date     string
}
