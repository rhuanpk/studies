package model

import "time"

// User is abstraction of user table.
type User struct {
	ID        *int       `json:"-"`
	UUID      *string    `json:"uuid,omitempty"`
	Name      *string    `json:"name,omitempty" sql:"name"`
	Username  *string    `json:"username,omitempty" sql:"username"`
	Email     *string    `json:"email,omitempty" sql:"email"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
