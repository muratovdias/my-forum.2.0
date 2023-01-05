package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"forum/models"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) CreateUser(user models.User) error {
	query := "INSERT INTO user (email, username, password) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, user.Email, user.Username, user.Password)
	if err != nil {
		log.Printf("repo: create user: %s", err)
		return fmt.Errorf(path+"create user: %w", err)
	}
	return nil
}

func (r *AuthRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := `SELECT  username, password FROM user WHERE email=$1`
	row := r.db.QueryRow(query, email)
	if err := row.Scan(&user.Username, &user.Password); err != nil {
		return user, fmt.Errorf(path+"get user by email: %w", err)
	}
	return user, nil
}

func (r *AuthRepo) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	query := `SELECT  username, password FROM user WHERE username=$1`
	row := r.db.QueryRow(query, username)
	if err := row.Scan(&user.Username, &user.Password); err != nil {
		return user, fmt.Errorf(path+"get user by username: %w", err)
	}
	return user, nil
}

func (r *AuthRepo) GetUserByToken(token string) (models.User, error) {
	var user models.User
	query := "SELECT * FROM user WHERE token=$1"
	row := r.db.QueryRow(query, token)
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Token, &user.TokenDuration)
	if err != nil {
		return models.User{}, fmt.Errorf(path+"get user by token %w", err)
	}
	return user, nil
}

func (r *AuthRepo) SaveToken(username, token string, duration time.Time) error {
	query := `UPDATE user SET token=$1, token_duration=$2 WHERE username=$3`
	_, err := r.db.Exec(query, token, duration, username)
	if err != nil {
		return fmt.Errorf("ERROR: /repository save token: %w", err)
	}
	return nil
}

func (r *AuthRepo) DeleteToken(token string) error {
	query := `UPDATE user SET token=NULL, token_duration=NULL WHERE token=$1`
	_, err := r.db.Exec(query, token)
	if err != nil {
		return err
	}
	return nil
}
