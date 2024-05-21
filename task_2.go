package main

import (
	"database/sql"
	"fmt"
	"log"
	"test-go-clickhouse-middle/config"
	"test-go-clickhouse-middle/internal/entity"
	"test-go-clickhouse-middle/pkg/clickhouse"
	"test-go-clickhouse-middle/pkg/logger"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	l := logger.New(cfg.Log.Level)

	// Repository
	ch, err := clickhouse.New(
		[]string{fmt.Sprintf("%s:%d", cfg.ClickHouse.Host, cfg.ClickHouse.Port)},
		cfg.ClickHouse.DB, cfg.ClickHouse.Username, cfg.ClickHouse.Password)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - clickhouse.New: %w", err))
	}
	defer ch.Close()

	// Insert test data
	insertTestData(ch.DB, l, "test", 10, "2030-01-01 12:00:00", "{\"some_field\":\"some_value\"}")

	// Output events by specified eventType and time range
	queryEventsByTypeAndTime(ch.DB, l, "login", "2023-01-01 12:00:00", "2024-01-01 12:00:00")
}

func insertTestData(conn *sql.DB, l *logger.Logger, eventType string, userID int, eventTime string, payload string) {
	query := "INSERT INTO events_db.events (eventType, userID, eventTime, payload) VALUES ($1, $2, $3, $4)"
	_, err := conn.Exec(query, eventType, userID, eventTime, payload)
	if err != nil {
		l.Error("error inserting test data:", err)
		return
	}
	l.Info("test data inserted successfully")
}

func queryEventsByTypeAndTime(conn *sql.DB, l *logger.Logger, eventType string, startTime string, endTime string) {
	var event entity.Event

	query := "SELECT * FROM events_db.events WHERE eventType = $1 AND eventTime BETWEEN $2 AND $3"
	rows, err := conn.Query(query, eventType, startTime, endTime)
	if err != nil {
		l.Error("error querying data:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&event.EventID, &event.EventType, &event.UserID, &event.EventTime, &event.Payload)
		if err != nil {
			l.Error("error scanning row:", err)
			return
		}
		l.Info("EventID: %d, EventType: %s, UserID: %d, EventTime: %s, Payload: %s\n",
			event.EventID, event.EventType, event.UserID, event.EventTime, event.Payload)
	}
}
