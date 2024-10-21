package user

import (
	"context"
	"fmt"

	"dev/internal/model"
	"dev/internal/storage/database"
	"dev/pkg"

	"github.com/go-hl/dbutil/extract"
)

// Update is the repository for update a user in database.
func Update(uuid string, user model.User) error {
	if err := database.DB.Ping(); err != nil {
		return fmt.Errorf("%w: %w", pkg.ErrDBConnection, err)
	}

	tx, err := database.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("%w: %w", pkg.ErrDBTransaction, err)
	}

	query, values, err := extract.QueryAndValues(
		&user, extract.UpdateQueryType, extract.QuestionMarkQueryPlaceholder, "sql", "users",
	)
	if err != nil {
		return err
	}

	if _, err := tx.Exec(
		query+`, updated_at = CURRENT_TIMESTAMP WHERE uuid = ?`,
		append(values, uuid)...,
	); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("%w: %w", pkg.ErrDBTransaction, err)
		}
		return fmt.Errorf("%w: %w", pkg.ErrDBUpdate, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%w: %w", pkg.ErrDBTransaction, err)
	}

	return nil
}
