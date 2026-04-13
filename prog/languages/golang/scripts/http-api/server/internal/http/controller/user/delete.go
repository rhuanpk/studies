package user

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	service "dev/internal/service/user"
	"dev/pkg"
)

// Delete is the controller for delete a user in database.
func Delete(res http.ResponseWriter, req *http.Request) {
	uuid := req.PathValue("uuid")

	if err := service.Delete(uuid); err != nil {
		log.Println("error in delete user:", err)

		message := http.StatusText(http.StatusInternalServerError)
		code := http.StatusInternalServerError

		if errors.Is(err, pkg.ErrDBDelete) {
			message = "error in delete user"
		}

		if errors.Is(err, sql.ErrNoRows) {
			message = "user not found"
			code = http.StatusNotFound
		}

		http.Error(res, message, code)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}
