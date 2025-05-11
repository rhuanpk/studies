package user

import (
	"context"
	"fmt"
	"strings"

	"dev/internal/storage/database"
	"dev/pkg"
)

// Delete is the repository for Delete a user in database.
func Delete(uuid string) error {
	if err := database.DB.Ping(); err != nil {
		return fmt.Errorf("%w: %w", pkg.ErrDBConnection, err)
	}

	tx, err := database.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("%w: %w", pkg.ErrDBTransaction, err)
	}

	user, err := Get(uuid, true)
	if err != nil {
		return err
	}

	query := `
		UPDATE users
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE uuid = ?
	`
	if user.DeletedAt != nil {
		query = strings.ReplaceAll(query, `CURRENT_TIMESTAMP`, `NULL`)
	}

	if _, err := tx.Exec(query, uuid); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("%w: %w", pkg.ErrDBTransaction, err)
		}
		return fmt.Errorf("%w: %w", pkg.ErrDBDelete, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%w: %w", pkg.ErrDBTransaction, err)
	}

	return nil
}
