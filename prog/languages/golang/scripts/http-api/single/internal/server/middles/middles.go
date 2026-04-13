package middles

import (
	"net/http"

	"dev/pkg/logger"
)

func middles(response http.ResponseWriter, _ *http.Request) {
	// set variables
	log := logger.NewLogger("middles")

	// reply request
	response.WriteHeader(http.StatusOK)

	// trace logs
	log.Short("server")
	log.Println("middlewares test ok")
}
