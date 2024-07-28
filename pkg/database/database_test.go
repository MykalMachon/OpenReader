package database

import (
	"testing"
)

func TestInit(t *testing.T) {
	db, err := Init(":memory:")

	if err != nil {
		t.Errorf("failed to init database: got error %v", err)
	}

	// Check if "settings" table exists
	_, err = db.Exec("SELECT 1 FROM settings LIMIT 1")
	if err != nil {
		t.Errorf("failed to find 'settings' table: got error %v", err)
	}

	// Check if "books" table exists
	_, err = db.Exec("SELECT 1 FROM books LIMIT 1")
	if err != nil {
		t.Errorf("failed to find 'books' table: got error %v", err)
	}
}
