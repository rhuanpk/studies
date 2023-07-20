package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET(
		"/gogenericmap",
		func(context *gin.Context) {
			context.JSON(
				http.StatusOK,
				map[string]any{
					"zero": true,
				},
			)
		},
	)
	server.GET(
		"/gingenericmap",
		func(context *gin.Context) {
			context.JSON(
				http.StatusOK,
				gin.H{
					"one": false,
				},
			)
		},
	)
	server.Run(":8888")
}
