package user

import (
	"context"
	"fmt"

	"dev/internal/model"
	"dev/internal/storage/database"
	"dev/pkg"
)

// Create is the repository for create a user in database.
func Create(user model.User) error {
	if err := database.DB.Ping(); err != nil {
		return fmt.Errorf("%w: %w", pkg.ErrDBConnection, err)
	}

	tx, err := database.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("%w: %w", pkg.ErrDBTransaction, err)
	}

	if _, err := tx.Exec(`
		INSERT INTO users (name, username, email)
		VALUES (?, ?, ?)
	`,
		user.Name, user.Username, user.Email,
	); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("%w: %w", pkg.ErrDBTransaction, err)
		}
		return fmt.Errorf("%w: %w", pkg.ErrDBInsert, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%w: %w", pkg.ErrDBTransaction, err)
	}

	return nil
}
