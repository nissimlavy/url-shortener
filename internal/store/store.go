package store

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Example: Creating a table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS url_mappings (short_url TEXT PRIMARY KEY, original_url TEXT)")
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	return &Store{db: db}, nil
}

func (s *Store) CreateURLMapping(shortURL, originalURL string) error {
	_, err := s.db.Exec("INSERT INTO url_mappings (short_url, original_url) VALUES (?, ?)", shortURL, originalURL)
	if err != nil {
		log.Printf("Failed to insert URL mapping: %v", err)
		return err
	}
	return nil
}

func (s *Store) GetOriginalURL(shortURL string) (string, error) {
	var originalURL string
	row := s.db.QueryRow("SELECT original_url FROM url_mappings WHERE short_url = ?", shortURL)
	err := row.Scan(&originalURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}
