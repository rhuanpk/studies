package middle

import (
	"net/http"

	"dev/pkg/logger"
)

// Bar is the Bar middlware test.
func Bar(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			log := logger.NewLogger("bar")
			log.Short("middle")
			log.Println("bar logged")
			next(response, request)
		},
	)
}
