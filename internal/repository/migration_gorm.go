package repository

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/muratovdias/my-forum.2.0/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	userTable = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(30) UNIQUE,
		username VARCHAR(30) UNIQUE,
		password TEXT,
		token TEXT DEFAULT NULL,
		token_duration TIMESTAMP DEFAULT NULL
		);`
	postTable = `CREATE TABLE IF NOT EXISTS posts(
		id SERIAL PRIMARY KEY,
		author_id INTEGER,
		title VARCHAR(150),
		category VARCHAR(50),
		content TEXT,
		author	VARCHAR(30), 
		date text,
		FOREIGN KEY (author_id) REFERENCES users (id)
	);`
	commentTable = `CREATE TABLE IF NOT EXISTS comments(
		id SERIAL PRIMARY KEY,
		user_id INTEGER,
		post_id INTEGER,
		text TEXT,
		author	VARCHAR(30), 
		date TEXT,
		FOREIGN KEY (post_id) REFERENCES posts (id)
	);`
	userPostVote = `CREATE TABLE IF NOT EXISTS user_post_votes(
		id SERIAL PRIMARY KEY,
		user_id INTEGER,
		post_id INTEGER,
		vote	BOOLEAN,
		FOREIGN KEY (user_id) REFERENCES users (id),
		FOREIGN KEY (post_id) REFERENCES posts (id)
	);`
	userCommentVote = `CREATE TABLE IF NOT EXISTS user_comment_votes(
		id SERIAL PRIMARY KEY,
		user_id INTEGER,
		comment_id INTEGER,
		vote	BOOLEAN,
		FOREIGN KEY (user_id) REFERENCES users (id),
		FOREIGN KEY (comment_id) REFERENCES comments (id)
	);`
)

const path = "repository: "

func InitDB(c *config.ConfigDB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", c.Host, c.User, c.Password, c.DbName, c.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}

func CreateTables(db *gorm.DB) error {
	tables := []string{userTable, postTable, commentTable, userPostVote, userCommentVote}
	for _, table := range tables {
		r := db.Exec(table)
		if r.Error != nil {
			return fmt.Errorf(path+"create tables: %w", r.Error)
		}
	}
	return nil
}
