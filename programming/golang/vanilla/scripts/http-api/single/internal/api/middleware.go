package api

import "net/http"

// Middleware is the default middleware.
func Middleware(requests *chan *http.Request, next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		*requests <- request
		next(response, request)
	}
}
