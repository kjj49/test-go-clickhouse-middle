package http

import (
	"net/http"
	"test-go-clickhouse-middle/internal/entity"
	"test-go-clickhouse-middle/internal/usecase"
	"test-go-clickhouse-middle/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type EventRoutes struct {
	u usecase.Event
	l logger.Interface
}

func newEventRoutes(handler *gin.RouterGroup, u usecase.Event, l logger.Interface) {
	r := &EventRoutes{u, l}

	h := handler.Group("event")
	{
		h.POST("", r.insertEvent)
	}
}

type eventRequest struct {
	EventType string `json:"eventType"  binding:"required" example:"login"`
	UserID    int    `json:"UserID"     binding:"required" example:"1"`
	EventTime string `json:"EventTime"  binding:"required" example:"2023-04-09 13:00:00"`
	Payload   string `json:"Payload"  binding:"required" example:"{\"some_field\":\"some_value\"}"`
}

type successResponse struct {
	Status string `json:"status" example:"ok"`
}

// @Summary Insert Event
// @Tags event
// @Description insert event
// @ID create-event
// @Accept  json
// @Produce  json
// @Param input body entity.Event true "event info"
// @Success 200 {object} successResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/event [post]
func (r *EventRoutes) insertEvent(ctx *gin.Context) {
	var event eventRequest
	if err := ctx.BindJSON(&event); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	eventTime, err := time.Parse("2006-01-02 15:04:05", event.EventTime)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid EventTime format")
		return
	}

	err = r.u.InsertEvent(
		ctx.Request.Context(),
		entity.Event{
			EventType: event.EventType,
			UserID:    event.UserID,
			EventTime: eventTime,
			Payload:   event.Payload,
		},
	)
	if err != nil {
		r.l.Error(err, "http - insertEvent")
		newErrorResponse(ctx, http.StatusInternalServerError, "database problems")
		return
	}

	ctx.JSON(http.StatusOK, successResponse{
		Status: "ok",
	})
}
