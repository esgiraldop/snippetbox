package models

import (
	"database/sql"
	"errors"
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
	stmt := `SELECT id, title, content, 		created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id) // Returns a pointer to a sql.Row object which holds the result from the database

	s := &Snippet{} // Initialize a pointer to a new zeroed Snippet struct

	// row.Scan copies the values from each field in sql.Row to the corresponding field in the Snippet struct. Notice the arguments to row.Scan are *pointers* to the place I want to copy the data into, and the number of arguments must be exactly the same as the number of columns returned in the statement
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // Using "errors.Is" is safer than using "==" since an error can be wrapped to add additional information. If the error happens to be wrapped, then an entirely new error value is created, so "==" doesn't work anymore
			// If no row is found, return a custom error
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// to return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
