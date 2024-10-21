package user

import (
	"fmt"

	"dev/internal/model"
	"dev/internal/storage/database"
	"dev/pkg"
)

// List is the repository for list users in database.
func List(scoped bool) (*[]model.User, error) {
	var users []model.User

	if err := database.DB.Ping(); err != nil {
		return nil, fmt.Errorf("%w: %w", pkg.ErrDBConnection, err)
	}

	query := `
		SELECT ` + columns + `
		FROM users
	`

	rows, err := database.DB.Query(
		map[bool]string{
			true:  query,
			false: query + `WHERE deleted_at IS NULL`,
		}[scoped],
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", pkg.ErrDBSelect, err)
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID, &user.UUID,
			&user.Name, &user.Username, &user.Email,
			&user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
		); err != nil {
			return nil, fmt.Errorf("%w: %w", pkg.ErrDBScan, err)
		}
		users = append(users, user)
	}

	return &users, nil
}

// Get is the repository for get a user in database.
func Get(uuid string, scoped bool) (*model.User, error) {
	var user model.User

	if err := database.DB.Ping(); err != nil {
		return nil, fmt.Errorf("%w: %w", pkg.ErrDBConnection, err)
	}

	query := `
		SELECT ` + columns + `
		FROM users
		WHERE uuid = ?
	`

	row := database.DB.QueryRow(
		map[bool]string{
			true:  query,
			false: query + `AND deleted_at IS NULL`,
		}[scoped], uuid,
	)

	if err := row.Scan(
		&user.ID, &user.UUID,
		&user.Name, &user.Username, &user.Email,
		&user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	); err != nil {
		return nil, fmt.Errorf("%w: %w", pkg.ErrDBScan, err)
	}

	return &user, nil
}
