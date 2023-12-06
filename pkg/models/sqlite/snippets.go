package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/robwestbrook/snippetbox/pkg/models"
)

// Define a SnippetModel Type to wrap a
// sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

//
// Insert a new Snippet into the database
// A method of the SnippetModel type
// Takes in:
//	- snippet title
//	- snippet content
//	- snippet expires date
// Returns:
//	- int			snippet ID
//	- error		any function errors
//
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	
	// Create an SQLite prepared statement
	stmt := `
		INSERT INTO snippets(title, content, expires)
		VALUES(?, ?, DATETIME('now', '+' || ? || ' DAYS'))
		`
	
	// Execute the statement and check for errors
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Get the ID of the created snippet and check for errors
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	// Convert ID to type int64 and return with no error
	return int(id), nil

}

// Get a specific snippet by ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Return 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}