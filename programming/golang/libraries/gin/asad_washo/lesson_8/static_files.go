package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	// necessário startar o servidor na pasta absoluta do arquivo para que os caminhos relativos fiquem coerentes.
	server.Static("/", "./")
	server.Run(":8888")
}
