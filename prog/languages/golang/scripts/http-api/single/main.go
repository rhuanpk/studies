package main

import (
	"net/http"
	"strings"
	"time"

	"dev/internal/client"
	"dev/internal/consts"
	"dev/pkg/logger"
	"dev/pkg/router"

	_ "dev/internal/process/inits"
)

func main() {
	// set variables
	const (
		banner = `»--> Golang HTTP API <--«`
		delay  = 3
	)
	done := make(chan bool, 1)
	requests := make(chan bool, 1)
	log := logger.NewLogger("main")

	// trace logs
	println(banner + "\n" + strings.Repeat("-", len(banner)))
	log.Short()
	log.Printf(`listening on "%v%v"`, consts.ServerHost.Host, consts.ServerPort.Parse())

	// start background server
	go func() {
		router.Setup(&done, &requests, router.Schemes)
		log.Fatalln(http.ListenAndServe(consts.ServerPort.Parse(), nil))
	}()
	<-done // wait setup done

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
		time.Sleep(time.Second * delay)
	}
}
