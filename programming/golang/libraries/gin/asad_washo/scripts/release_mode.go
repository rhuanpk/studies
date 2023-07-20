package scripts

import "github.com/gin-gonic/gin"

func releaseMode() {
	// set release mode (debug is default).
	gin.SetMode(gin.ReleaseMode)
	gin.Default().Run()
}
