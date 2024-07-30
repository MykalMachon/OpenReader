package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Author struct {
	Id   int
	Name string
}

type Book struct {
	id            int
	name          string
	author_id     string
	file_location string
	ingested_at   time.Time
	status        string
	pages         int
}

func Init(file string) (*sql.DB, error) {
	// create or open database
	db, err := sql.Open("sqlite3", file)

	if err != nil {
		fmt.Println("failed to init sqlite", err)
		return nil, err
	}

	// set the proper pragmas
	_, err = db.Exec(`PRAGMA journal_mode = WAL;`)

	if err != nil {
		fmt.Println("failed to set the journal mode to WAL", err)
		return nil, err
	}

	_, err = db.Exec(`PRAGMA foreign_keys = ON;`)
	if err != nil {
		fmt.Println("failed to set enforce foreign_keys on database", err)
	}

	// create basic tables
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS authors (
			id INTEGER PRIMARY KEY, 
			name TEXT
		);
	`)

	if err != nil {
		fmt.Println("failed to create a authors table", err)
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY,
			name TEXT,
			author_id INTEGER,
			file_location TEXT,
			ingested_at TEXT DEFAULT CURRENT_TIMESTAMP,
			status TEXT,
			pages INTEGER,
			FOREIGN KEY (author_id) REFERENCES authors(id)
		);
	`)

	if err != nil {
		fmt.Println("failed to create books table", err)
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS settings (
			id INTEGER PRIMARY KEY,
			key TEXT,
			value TEXT
		);
	`)

	if err != nil {
		fmt.Println("failed to create settings table", err)
		return nil, err
	}

	return db, nil
}
