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

// Create is the controller for create a user in database.
func Create(res http.ResponseWriter, req *http.Request) {
	var user model.User

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		log.Println("error in json decode:", err)
		http.Error(res, "cannot deserialize body", http.StatusUnprocessableEntity)
		return
	}

	if err := service.Create(user); err != nil {
		log.Println("error in create user:", err)

		message := http.StatusText(http.StatusInternalServerError)
		code := http.StatusInternalServerError

		if errors.Is(err, pkg.ErrDBInsert) {
			message = "error in insert user"
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
