package router

import (
	"net/http"

	"dev/pkg/logger"
	"dev/pkg/router/middle"
)

// Setup set's up the api endpoints.
func Setup(done *chan bool, requests *chan bool, schemes []Scheme) {
	// set variables
	log := logger.NewLogger("router")

	// trace logs
	log.Short()

	// iterate over all scheme to setup them
	for _, scheme := range schemes {
		log.Println(scheme.Method, "->", scheme.Route)
		http.HandleFunc(
			scheme.Parse(),
			middle.Wrapper(
				requests,
				scheme.Handler,
				scheme.Middlewares...,
			),
		)
	}

	*requests <- true // send done signal to requests channel
	*done <- true     // send done signal to main function procced
}
