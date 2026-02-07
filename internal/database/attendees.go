package database

import (
	"context"
	"database/sql"
	"time"
)

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	ID      int `json:"id"`
	EventID int `json:"eventId"`
	UserID  int `json:"userId"`
}

func (m *AttendeeModel) Insert(attendee *Attendee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO attendees (event_id, user_id) VALUES ($1, $2) RETURNING id"

	return m.DB.QueryRowContext(ctx, query, attendee.EventID, attendee.UserID).Scan(&attendee.ID)
}

func (m *AttendeeModel) IsAttendeeExists(eventID, userID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT COUNT(*) FROM attendees WHERE event_id = $1 AND user_id = $2"
	var count int
	err := m.DB.QueryRowContext(ctx, query, eventID, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetByAttendeeAndEvent retrieves an attendee by user ID and event ID

// GetAttendeesByEventID retrieves all attendees for a specific event ID
