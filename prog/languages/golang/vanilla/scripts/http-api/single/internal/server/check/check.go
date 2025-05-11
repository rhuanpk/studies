package check

import (
	"net/http"

	"dev/pkg/logger"
)

func check(response http.ResponseWriter, _ *http.Request) {
	// set variables
	log := logger.NewLogger("check")

	// reply request
	response.WriteHeader(http.StatusOK)

	// trace logs
	log.Short("server")
	log.Println("health check ok")
}
