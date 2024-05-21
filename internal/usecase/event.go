package usecase

import (
	"context"
	"fmt"
	"test-go-clickhouse-middle/internal/entity"
)

// EventUseCase
type EventUseCase struct {
	repository EventRepository
}

// New
func New(r EventRepository) *EventUseCase {
	return &EventUseCase{
		repository: r,
	}
}

// InsertEvent - insert event to store.
func (uc *EventUseCase) InsertEvent(ctx context.Context, e entity.Event) error {
	err := uc.repository.InsertEvent(ctx, e)
	if err != nil {
		return fmt.Errorf("EventUseCase - InsertEvent - s.repository.InsertEvent: %w", err)
	}

	return nil
}
