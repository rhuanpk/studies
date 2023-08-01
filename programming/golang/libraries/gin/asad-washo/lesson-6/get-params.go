package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET(
		// necessário definir o nome dos parâmetros que serão pegos da URL para que possa usalos dentro da função com *gin.Context.Param().
		"/user/:name/identifier/:id",
		func(context *gin.Context) {
			name := context.Param("name")
			id := context.Param("id")
			context.String(http.StatusOK, name)
			context.String(http.StatusOK, id)
		},
	)
	server.Run(":8888")
}
