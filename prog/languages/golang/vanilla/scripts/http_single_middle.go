package main

import (
	"log"
	"net/http"
)

func middleware(param any, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Println("middleware step")
		if param != nil {
			log.Println("param:", param)
		}
		next.ServeHTTP(res, req)
	})
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /hello", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("hello"))
	})
	log.Fatalln(http.ListenAndServe(":9999", middleware(nil, router)))
}
