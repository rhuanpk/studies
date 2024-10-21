package user

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"dev/internal/model"
	service "dev/internal/service/user"
	"dev/pkg"
)

// Update is the controller for update a user in database.
func Update(res http.ResponseWriter, req *http.Request) {
	var user model.User
	uuid := req.PathValue("uuid")

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		log.Println("error in json decode:", err)
		http.Error(res, "cannot deserialize body", http.StatusUnprocessableEntity)
		return
	}

	if err := service.Update(uuid, user); err != nil {
		log.Println("error in update user:", err)

		message := http.StatusText(http.StatusInternalServerError)
		code := http.StatusInternalServerError

		if errors.Is(err, pkg.ErrDBUpdate) {
			message = "error in update user"
			if pkg.RegEx(`(?i)unique constraint failed`, err.Error()) {
				message = "user already exists"
				code = http.StatusConflict
			}
		}

		http.Error(res, message, code)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}
