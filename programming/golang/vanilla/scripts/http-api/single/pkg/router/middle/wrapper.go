package middle

import "net/http"

// Wrapper is the default middleware.
func Wrapper(requests *chan bool, handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	// function pre handler to signal requests channel
	handler = func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(
			func(response http.ResponseWriter, request *http.Request) {
				if !(len(*requests) == cap(*requests)) {
					*requests <- true
				}
				next(response, request)
			},
		)
	}(handler)

	// iterate over all middlewares to build the single wrapped function
	for index := len(middlewares) - 1; index > -1; index-- {
		handler = middlewares[index](handler)
	}
	return handler
}
