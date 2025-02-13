package migrations

import (
	"database/sql"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// Migrate applies the migration
func Migrate(db *sql.DB) error {
	// Example: Create a table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS example (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);`)
	return err
}

// Rollback reverts the migration
func Rollback(db *sql.DB) error {
	// Example: Drop the table
	_, err := db.Exec(`DROP TABLE IF EXISTS example;`)
	return err
}
