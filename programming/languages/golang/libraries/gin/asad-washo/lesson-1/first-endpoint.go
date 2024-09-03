package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.GET(
		"/",
		func(context *gin.Context) {
			context.String(http.StatusOK, "hello world!")
		},
	)
	server.Run(":8888")
}
