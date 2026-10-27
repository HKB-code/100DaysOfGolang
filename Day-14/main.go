package main

import (
	"project/REST-API/db"
	"project/REST-API/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()

	routes.ResgisterRoutes(server)

	server.Run(":8080") //localHost: 8080
}
