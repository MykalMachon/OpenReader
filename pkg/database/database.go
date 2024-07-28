package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

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
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY,
			name TEXT,
			author TEXT,
			file_location TEXT,
			ingested_at TEXT DEFAULT CURRENT_TIMESTAMP,
			status TEXT,
			pages INTEGER
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
