package middle

import (
	"net/http"

	"dev/pkg/logger"
)

// Foo is the Foo middlware test.
func Foo(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			log := logger.NewLogger("foo")
			log.Short("middle")
			log.Println("foo logged")
			next(response, request)
		},
	)
}
