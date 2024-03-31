package router

import (
	"fmt"
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
		if scheme.Method == "" {
			scheme.Method = "ANY"
		}
		for index, parse := range scheme.Parses() {
			log.Printf(
				`%s%`+fmt.Sprint((4+(7-len(scheme.Method))))+`s%s`,
				scheme.Method, " -> ", scheme.Routes[index],
			)
			http.HandleFunc(
				parse,
				middle.Wrapper(
					requests,
					scheme.Handler,
					scheme.Middlewares...,
				),
			)
		}
	}

	*requests <- true // send done signal to requests channel
	*done <- true     // send done signal to main function procced
}
