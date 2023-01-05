package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
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

type ConfigDB struct {
	Driver string
	Path   string
	Name   string
}

func NewConfDB() *ConfigDB {
	return &ConfigDB{
		Driver: "sqlite3",
		Name:   "forum.db",
	}
}

func InitDB(c *ConfigDB) (*sql.DB, error) {
	db, err := sql.Open(c.Driver, c.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTables(db *sql.DB) error {
	tables := []string{userTable, postTable, commentTable, likeTable, dislikeTable}
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			return fmt.Errorf(path+"create tables: %w", err)
		}
	}
	return nil
}
