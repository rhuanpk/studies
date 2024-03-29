package test

import (
	"net/http"

	"dev/pkg/logger"
)

func test(response http.ResponseWriter, _ *http.Request) {
	// set variables
	log := logger.NewLogger("test")

	// reply request
	response.WriteHeader(http.StatusOK)

	// trace logs
	log.Short("server")
	log.Println("testing ok")
}
