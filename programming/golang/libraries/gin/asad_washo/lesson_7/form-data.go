package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.POST(
		"/post",
		func(context *gin.Context) {
			// pega o valor do campo "field" e caso não esteja setado usar "default" como valor padrão (vazio (null) não cai no "default").
			user := context.DefaultPostForm("field", "default")
			context.String(http.StatusOK, "%s!", user)
		},
	)
	server.Run(":8888")
}
