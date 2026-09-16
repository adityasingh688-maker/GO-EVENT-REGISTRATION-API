package routes

import (
	"example.com/api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouts(server *gin.Engine) {
	server.GET("/events", getEvents)    //GET, POST, PUT, PACH, DELETE
	server.GET("/events/:id", getEvent) // /events/1, /events/5   // dynamic path handler or routs

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authorization)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
