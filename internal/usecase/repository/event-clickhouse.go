package repository

import (
	"context"

	"test-go-clickhouse-middle/internal/entity"
	"test-go-clickhouse-middle/pkg/clickhouse"
)

// EventRepository
type EventRepository struct {
	*clickhouse.ClickHouse
}

// New -.
func New(ch *clickhouse.ClickHouse) *EventRepository {
	return &EventRepository{ch}
}

// InsertEvent
func (r *EventRepository) InsertEvent(ctx context.Context, e entity.Event) error {
	query := "INSERT INTO events (eventType, userID, eventTime, payload) VALUES ($1, $2, $3, $4)"
	_, err := r.Exec(query, e.EventType, e.UserID, e.EventTime, e.Payload)
	return err
}
