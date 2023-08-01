package scripts

import "github.com/gin-gonic/gin"

func endpointsGroup() {
	server := gin.Default()
	users := server.Group("/usuarios")
	users.GET(
		"/first",
		func(ctx *gin.Context) {},
	)
	users.GET(
		"/second",
		func(ctx *gin.Context) {},
	)
}
