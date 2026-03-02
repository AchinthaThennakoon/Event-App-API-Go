package main

import (
	"fmt"
	"net/http"
	"rest-api-gin-go/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) createEvent(c *gin.Context) {
	var event database.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := app.GetUserFromContext(c)
	if user == nil || user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	event.OwnerID = user.ID
	if err := app.models.Events.Insert(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, event)
}

func (app *application) getAllEvents(c *gin.Context) {
	events, err := app.models.Events.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}

	c.JSON(http.StatusOK, events)
}

func (app *application) getEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := app.models.Events.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (app *application) updateEvent(c *gin.Context) {
	// var event database.Event
	// if err := c.ShouldBindJSON(&event); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// if err := app.models.Events.Update(&event); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
	// 	return
	// }
	// c.JSON(http.StatusOK, event)
}

func (app *application) deleteEvent(c *gin.Context) {
	//delete event logic here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	user := app.GetUserFromContext(c)
	if user == nil || user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	existingEvent, err := app.models.Events.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}
	if existingEvent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	fmt.Printf("Existing Event Owner ID: %d, User ID: %d\n", existingEvent.OwnerID, user.ID)

	if existingEvent.OwnerID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the creator of this event"})
		return
	}

	if err := app.models.Events.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}

func (app *application) addAttendeeToEvent(c *gin.Context) {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	//retrieve attendee logic here
	IsUserExists, err := app.models.Users.IsUserExists(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}
	if !IsUserExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	//retrieve event
	isEventExists, err := app.models.Events.IsEventExists(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}
	if !isEventExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	//check if user is already an attendee of the event
	isAttendeeExists, err := app.models.Attendees.IsAttendeeExists(eventID, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check attendee"})
		return
	}
	if isAttendeeExists {
		c.JSON(http.StatusConflict, gin.H{"error": "User is already an attendee of the event"})
		return
	}

	var newAttendee database.Attendee
	newAttendee.EventID = eventID
	newAttendee.UserID = userId
	if err := app.models.Attendees.Insert(&newAttendee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add attendee to event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Attendee added successfully"})
}

// get attendee by user ID and event ID
func (app *application) GetByEventAndAttendee(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	attendee, err := app.models.Attendees.GetByEventAndAttendee(eventID, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve attendee"})
		return
	}
	if attendee == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attendee not found"})
		return
	}
	c.JSON(http.StatusOK, attendee)
}

// get attendees of an event
func (app *application) GetAttendeesByEventID(c *gin.Context) {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	attendees, err := app.models.Attendees.GetAttendeesByEventID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve attendees"})
		return
	}
	c.JSON(http.StatusOK, attendees)
}

func (app *application) removeAttendeeFromEvent(c *gin.Context) {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	//check if attendee exists
	isAttendeeExists, err := app.models.Attendees.IsAttendeeExists(eventID, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check attendee"})
		return
	}
	if !isAttendeeExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attendee not found"})
		return
	}

	if err := app.models.Attendees.Delete(eventID, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove attendee from event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Attendee removed successfully"})
}

func (app *application) GetEventsByAttendee(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	events, err := app.models.Attendees.GetEventsByAttendee(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}
	c.JSON(http.StatusOK, events)
}
