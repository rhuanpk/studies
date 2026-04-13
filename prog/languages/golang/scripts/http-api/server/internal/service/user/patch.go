package user

import (
	"dev/internal/model"
	repository "dev/internal/storage/repository/user"
)

// Update is the service for update a user in database.
func Update(uuid string, user model.User) error {
	return repository.Update(uuid, user)
}
