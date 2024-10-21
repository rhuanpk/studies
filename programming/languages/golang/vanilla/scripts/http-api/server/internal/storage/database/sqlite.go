package database

import (
	"database/sql"
	"fmt"

	"dev/pkg"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// DB is the global database connection.
var DB *sql.DB

// Setup setups and initialize database connection.
func Setup() error {
	var err error

	DB, err = sql.Open("sqlite3", "tmp.db")
	if err != nil {
		return fmt.Errorf("%w: %w", pkg.ErrDBConnection, err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("%w: %w", pkg.ErrDBConnection, err)
	}

	if _, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			uuid TEXT(36) DEFAULT (LOWER(
				HEX(RANDOMBLOB(4)) || '-' ||
				HEX(RANDOMBLOB(2)) || '-' ||
				'4' || SUBSTR(HEX(RANDOMBLOB(2)), 2) || '-' ||
				SUBSTR('ab89', 1 + (ABS(RANDOM()) % 4), 1) ||
				SUBSTR(HEX(RANDOMBLOB(2)), 2) || '-'
				|| HEX(RANDOMBLOB(6))
			)) NOT NULL UNIQUE,
			name TEXT,
			username TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL UNIQUE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME,
			deleted_at DATETIME
		)
	`); err != nil {
		return fmt.Errorf("%w: %w", pkg.ErrDBMigration, err)
	}

	return nil
}
