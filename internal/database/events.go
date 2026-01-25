package database

import (
	"context"
	"database/sql"
	"time"
)

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

func (m EventModel) Insert(event *Event) any {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO events (owner_id, name, description, date, location) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	return m.DB.QueryRowContext(ctx, query, event.OwnerID, event.Name, event.Description, event.Date, event.Location).Scan(&event.ID)
}
