package routes

import (
	"example.com/restapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // GET, POST, PUT, DELETE
	server.GET("/events/:id", getEvent) // /events/1, /event/4

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events", updateEvent)
	authenticated.DELETE("/events", deleteEvent)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}
