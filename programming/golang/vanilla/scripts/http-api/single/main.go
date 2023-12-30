package main

import (
	"net/http"
	"strings"
	"time"

	"backend/internal/api"
	"backend/internal/client"
	"backend/internal/consts"
	"backend/internal/server"
	"backend/pkg/logger"
)

func main() {
	// set variavles
	const banner = `»--> Golang HTTP API With Standard Library <--«`
	var requests = make(chan *http.Request, 1)
	log := logger.NewLogger(logger.Main)

	// trace logs
	println(banner + "\n" + strings.Repeat("-", len(banner)))
	log.Short()
	log.Printf(`listening on "%v%v"`, consts.ServerHost, consts.ServerPort.Parse())

	// start background server
	go func() {
		http.HandleFunc(
			consts.APIEndpoint.Parse(true),
			api.Middleware(&requests, server.Handler),
		)
		log.Fatalln(http.ListenAndServe(consts.ServerPort.Parse(), nil))
	}()
	time.Sleep(time.Second)

	// client do request
	client.Request()

	// stand by
	for {
		select {
		case _ = <-requests:
			log.Short()
			log.Println("> waiting requests <")
		default:
			log.Println("> waiting requests <")
		}
		time.Sleep(time.Second * 3)
	}
}
