package handler

import (
	"net/http"

	"dev/internal/http/controller/user"
)

// Setup setups the api endpoints.
func Setup() {
	// user
	{
		http.HandleFunc("GET /user", user.List)
		http.HandleFunc("GET /user/{uuid}", user.Get)
		http.HandleFunc("POST /user", user.Create)
		http.HandleFunc("PATCH /user/{uuid}", user.Update)
		http.HandleFunc("DELETE /user/{uuid}", user.Delete)
	}
}
