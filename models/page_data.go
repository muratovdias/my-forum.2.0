package models

type PageData struct {
	Username string
	Posts    []Post
}

type PostData struct {
	Username string
	Post     Post
	Comments []Comment
}
