package middle

import "net/http"

// CORS is the middleware to allow CORS policy.
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			request.Header.Set("Access-Control-Allow-Origin", "*")
			request.Header.Set("Access-Control-Allow-Methods", "*")
			request.Header.Set("Access-Control-Allow-Headers", "*")
			if request.Method == http.MethodOptions {
				response.WriteHeader(http.StatusContinue)
			}
			next(response, request)
		},
	)
}
