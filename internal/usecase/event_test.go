package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	gomock "go.uber.org/mock/gomock"

	"test-go-clickhouse-middle/internal/entity"
	"test-go-clickhouse-middle/internal/usecase"
)

var testEvent = entity.Event{
	EventID:   1,
	EventType: "login",
	UserID:    1,
	EventTime: time.Now(),
	Payload:   "{\"some_field\":\"some_value\"}",
}

func TestEventUseCase_InsertEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockEventRepository(ctrl)
	eventUseCase := usecase.New(mockRepo)

	ctx := context.Background()

	// Successful
	mockRepo.EXPECT().InsertEvent(ctx, testEvent).Return(nil)
	err := eventUseCase.InsertEvent(ctx, testEvent)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Error in the repository
	mockRepo.EXPECT().InsertEvent(ctx, testEvent).Return(errors.New("mock error"))
	err = eventUseCase.InsertEvent(ctx, testEvent)
	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestEventUseCase_InsertEvent_RepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockEventRepository(ctrl)
	eventUseCase := usecase.New(mockRepo)

	ctx := context.Background()
	testEvent := entity.Event{
		EventID:   1,
		EventType: "login",
		UserID:    1,
		EventTime: time.Now(),
		Payload:   "{\"some_field\":\"some_value\"}",
	}

	mockRepo.EXPECT().InsertEvent(ctx, testEvent).Return(errors.New("mock error"))
	err := eventUseCase.InsertEvent(ctx, testEvent)
	if err == nil {
		t.Error("expected an error from repository, got nil")
	}
}
