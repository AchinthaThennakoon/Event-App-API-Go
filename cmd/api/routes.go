package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{

		v1.GET("/events", app.getAllEvents)
		v1.GET("/event/:id", app.getEvent)

		// v1.GET("/events/:id/attendees/:userId", app.GetByEventAndAttendee)
		v1.GET("/events/:id/attendees", app.GetAttendeesByEventID)
		v1.GET("/events/attendees/:userId", app.GetEventsByAttendee)

		v1.POST("/auth/register", app.registerUser)
		v1.POST("/auth/login", app.loginUser)
	}

	authGroup := v1.Group("/")
	authGroup.Use(app.AuthMiddleware())
	{
		authGroup.POST("/events", app.createEvent)
		authGroup.PUT("/events/:id", app.updateEvent)
		authGroup.DELETE("/events/:id", app.deleteEvent)

		authGroup.POST("/events/:id/attendees/:userId", app.addAttendeeToEvent)
		authGroup.DELETE("/events/:id/attendees/:userId", app.removeAttendeeFromEvent)

	}

	return g
}
