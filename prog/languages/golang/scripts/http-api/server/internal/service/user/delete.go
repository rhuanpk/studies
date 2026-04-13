package user

import repository "dev/internal/storage/repository/user"

// Delete is the service for delete a user in database.
func Delete(uuid string) error {
	return repository.Delete(uuid)
}
