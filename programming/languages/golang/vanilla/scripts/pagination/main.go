package main

import (
	"dev/internal/infra/database"
	"dev/internal/infra/router"
)

func main() {
	// curl -fsSL 'localhost:9999/resource?page_number=1&page_size=10'
	defer database.DB.Close()
	router.Init()
}
