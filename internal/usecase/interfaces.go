// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"test-go-clickhouse-middle/internal/entity"
)

type (
	// Event
	Event interface {
		InsertEvent(ctx context.Context, e entity.Event) error
	}

	// EventRepository
	EventRepository interface {
		InsertEvent(ctx context.Context, e entity.Event) error
	}
)
