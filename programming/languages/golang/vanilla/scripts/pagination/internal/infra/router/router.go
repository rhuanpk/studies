package router

import (
	"log"

	"dev/internal/core/resource"

	"github.com/gin-gonic/gin"
)

// Init record the endpoints and start the server.
func Init() {
	router := gin.Default()
	router.GET(
		"/resource",
		resource.Resource,
	)
	log.Fatalln(router.Run(":9999"))
}
