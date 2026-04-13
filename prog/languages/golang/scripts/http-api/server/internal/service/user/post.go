package user

import (
	"dev/internal/model"
	repository "dev/internal/storage/repository/user"
)

// Create is the service for create a user in database.
func Create(user model.User) error {
	return repository.Create(user)
}
