package main

import (
	"log"
	"net/http"

	"dev/internal/http/handler"
	"dev/internal/storage/database"
)

func main() {
	if err := database.Setup(); err != nil {
		log.Fatalln("error in database setup:", err)
	}
	handler.Setup()
	log.Println("http server listnig on :9999")
	log.Fatalln(
		"error in start http server:",
		http.ListenAndServe(":9999", nil),
	)
}
