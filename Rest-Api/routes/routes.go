package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // GET, POST, PUT, DELETE
	server.GET("/events/:id", getEvent) // /events/1, /event/4
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
}
