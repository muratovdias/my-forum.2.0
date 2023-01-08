package repository

import (
	"database/sql"
	"fmt"

	"github.com/muratovdias/my-forum.2.0/models"
)

type CommentRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (r *CommentRepo) CreateComment(comment models.Comment) error {
	query := `INSERT INTO comment (user_id, post_id, text, author, date) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, comment.UserID, comment.PostID, comment.Text, comment.Author, comment.Date)
	if err != nil {
		fmt.Printf("repo: %s\n", err)
		return fmt.Errorf(path+"create comment: %w", err)
	}
	return nil
}

func (r *CommentRepo) CheckCommentExists(id string) error {
	query = `SELECT id FROM comment WHERE id =$1`
	row := r.db.QueryRow(query, id)
	var commentID int
	if err := row.Scan(&commentID); err != nil {
		fmt.Println("Check comment")
		return err
	}
	return nil
}

func (r *CommentRepo) GetCommentByPostID(id int) (*[]models.Comment, error) {
	query := `SELECT * FROM comment WHERE post_id =$1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf(path+"get post comment: %w", err)
	}
	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Likes, &comment.Dislikes, &comment.Text, &comment.Author, &comment.Date); err != nil {
			fmt.Printf("repo: %s\n", err)
			return nil, fmt.Errorf(path+"scan comment: %w", err)
		}
		comments = append(comments, comment)
	}
	return &comments, nil
}
