package routes

import (
	"project/REST-API/middelwares"

	"github.com/gin-gonic/gin"
)

func ResgisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) //GET, POST,PUT,PATCH,DELETE
	server.GET("/events/:id", getEventById)
	// server.POST("/events", middelwares.Authenticate, createEvent)

	authenticate := server.Group("/")
	authenticate.Use(middelwares.Authenticate)
	authenticate.POST("/events", createEvent)
	authenticate.PUT("/events/:id", updateEventById)
	authenticate.DELETE("/events/:id", deleteEventById)
	authenticate.POST("/events/:id/register", register)
	authenticate.DELETE("/events/:id/register", cancelRegister)

	server.POST("/signup", signUp)
	server.POST("/login", login)

}
