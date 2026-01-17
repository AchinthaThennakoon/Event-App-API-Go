package database

import "database/sql"

type EventModel struct {
	DB *sql.DB
}

type Event struct {
	ID          int    `json:"id"`
	OwnerID     int    `json:"ownerId" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date" binding:"required,datetime=2006-01-02"`
	Location    string `json:"location"`
}