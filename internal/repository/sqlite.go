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
	userTable = `CREATE TABLE IF NOT EXISTS user (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE,
		username TEXT UNIQUE,
		password TEXT,
		token TEXT DEFAULT NULL,
		token_duration DATETIME DEFAULT NULL
		);`
	postTable = `CREATE TABLE IF NOT EXISTS post(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		author_id INTEGER,
		like INTEGER DEFAULT 0,
		dislike INTEGER DEFAULT 0,
		title TEXT,
		category TEXT,
		content TEXT,
		author	TEXT, 
		date text,
		FOREIGN KEY (author_id) REFERENCES user (id)
	);`
	commentTable = `CREATE TABLE IF NOT EXISTS comment(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		post_id INTEGER,
		like INTEGER DEFAULT 0,
		dislike INTEGER DEFAULT 0,
		text TEXT,
		author	TEXT, 
		date TEXT,
		FOREIGN KEY (post_id) REFERENCES post (id)
	);`
	likeTable = `CREATE TABLE IF NOT EXISTS like(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		post_id INTEGER DEFAULT 0,
		comment_id INTEGER DEFAULT 0,
		active	INTEGER DEFAULT 0,
		FOREIGN KEY (post_id) REFERENCES post (id)
	);`
	dislikeTable = `CREATE TABLE IF NOT EXISTS dislike(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		post_id INTEGER DEFAULT 0,
		comment_id INTEGER DEFAULT 0,
		active	INTEGER DEFAULT 0,
		FOREIGN KEY (post_id) REFERENCES post (id)
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
	tables := []string{userTable, postTable, commentTable, likeTable, dislikeTable}
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			return fmt.Errorf(path+"create tables: %w", err)
		}
	}
	return nil
}
