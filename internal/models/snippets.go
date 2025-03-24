package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Type to wrap a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// To insert a snippet into the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// To get a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// to return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
