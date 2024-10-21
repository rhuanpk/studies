package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	service "dev/internal/service/user"
	"dev/pkg"
)

// List is the controller for list users in database.
func List(res http.ResponseWriter, req *http.Request) {
	var (
		scoped bool
		err    error
	)

	if req.URL.Query().Has("deleted") {
		deleted := req.URL.Query().Get("deleted")
		scoped, err = strconv.ParseBool(deleted)
		if err != nil {
			log.Println("error in parse deleted query param:", err)
			http.Error(res, "invalid deleted query param", http.StatusBadRequest)
			return
		}
	}

	users, err := service.List(scoped)
	if err != nil {
		log.Println("error in list users:", err)

		message := http.StatusText(http.StatusInternalServerError)

		if errors.Is(err, pkg.ErrDBSelect) {
			message = "error in list users"
		}

		http.Error(res, message, http.StatusInternalServerError)
		return
	}

	if len(*users) < 1 {
		http.Error(res, "users not found", http.StatusNotFound)
		return
	}

	payload, err := json.Marshal(users)
	if err != nil {
		log.Println("error in json marshal:", err)
		http.Error(res, "cannot serialize body", http.StatusInternalServerError)
		return
	}

	if _, err := res.Write(payload); err != nil {
		log.Println("error in write response:", err)
	}
}

// Get is the controller for get a user in database.
func Get(res http.ResponseWriter, req *http.Request) {
	var (
		scoped bool
		err    error
	)
	uuid := req.PathValue("uuid")

	if req.URL.Query().Has("deleted") {
		deleted := req.URL.Query().Get("deleted")
		scoped, err = strconv.ParseBool(deleted)
		if err != nil {
			log.Println("error in parse deleted query param:", err)
			http.Error(res, "invalid deleted query param", http.StatusBadRequest)
			return
		}
	}

	user, err := service.Get(uuid, scoped)
	if err != nil {
		log.Println("error in get user:", err)

		message := http.StatusText(http.StatusInternalServerError)
		code := http.StatusInternalServerError

		if errors.Is(err, sql.ErrNoRows) {
			message = "user not found"
			code = http.StatusNotFound
		}

		http.Error(res, message, code)
		return
	}

	payload, err := json.Marshal(user)
	if err != nil {
		log.Println("error in json marshal:", err)
		http.Error(res, "cannot serialize body", http.StatusInternalServerError)
		return
	}

	if _, err := res.Write(payload); err != nil {
		log.Println("error in write response:", err)
	}
}
