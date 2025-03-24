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

	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))` // Using placeholders to avoid SQL injection

	result, err := m.DB.Exec(stmt, title, content, expires) // Returns a sql.Result type, which contains some basic information about what happened when the statement was executed
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId() // Getting the id of the newly inserted record
	// rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// To get a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// to return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
