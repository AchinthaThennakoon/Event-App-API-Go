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

func (m EventModel) GetAll() (any, any) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, owner_id, name, description, date, location FROM events"

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []*Event{}

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.OwnerID, &event.Name, &event.Description, &event.Date, &event.Location)
		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (m EventModel) GetByID(id int) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, owner_id, name, description, date, location FROM events WHERE id = $1"

	var event Event
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&event.ID, &event.OwnerID, &event.Name, &event.Description, &event.Date, &event.Location)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (m EventModel) Delete(id int) any {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM events WHERE id = $1"

	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (m EventModel) Update(event *Event) any {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "UPDATE events SET owner_id = $1, name = $2, description = $3, date = $4, location = $5 WHERE id = $6"

	_, err := m.DB.ExecContext(ctx, query, event.OwnerID, event.Name, event.Description, event.Date, event.Location, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m EventModel) IsEventExists(id int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT COUNT(*) FROM events WHERE id = $1"

	var count int
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
