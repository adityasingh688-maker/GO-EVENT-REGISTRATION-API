package main

import (
	"example.com/api/api-test/db"
	"example.com/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRouts(server)

	server.Run(":8080") // localhost:8080
}
