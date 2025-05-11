package user

import (
	"dev/internal/model"
	repository "dev/internal/storage/repository/user"
)

// List is the service for list users in database.
func List(scoped bool) (*[]model.User, error) {
	return repository.List(scoped)
}

// Get is the service for get a user in database.
func Get(uuid string, scoped bool) (*model.User, error) {
	return repository.Get(uuid, scoped)
}
