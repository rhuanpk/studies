package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("./templates/*")
	server.GET(
		"/",
		func(context *gin.Context) {
			context.HTML(
				http.StatusOK,
				"index.tmpl",
				map[string]string{
					"tittle": "home page",
				},
			)
		},
	)
	server.GET(
		"/about",
		func(context *gin.Context) {
			context.HTML(
				http.StatusOK,
				"index.tmpl",
				map[string]string{
					"tittle": "about page",
				},
			)
		},
	)
	server.GET(
		"/contact",
		func(context *gin.Context) {
			context.HTML(
				http.StatusOK,
				"index.tmpl",
				map[string]string{
					"tittle": "contact page",
				},
			)
		},
	)
	server.Run(":8888")
}
