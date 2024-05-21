// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "time"

// Event
type Event struct {
	EventID   int       `json:"eventID"       example:"1"`
	EventType string    `json:"eventType"  example:"login"`
	UserID    int       `json:"UserID"     example:"1"`
	EventTime time.Time `json:"EventTime"  example:"2023-04-09 13:00:00"`
	Payload   string    `json:"Payload"  example:"{\"some_field\":\"some_value\"}"`
}
