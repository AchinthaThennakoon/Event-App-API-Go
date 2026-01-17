package database

import "database/sql"

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	ID        int    `json:"id"`
	EventID   int    `json:"eventId"`
	UserID    int    `json:"userId"`
}