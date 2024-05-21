// Package clickhouse implements ClickHouse connection.
package clickhouse

import (
	"database/sql"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type ClickHouse struct {
	*sql.DB
}

// New
func New(addr []string, database, username, password string) (*ClickHouse, error) {
	ch := &ClickHouse{}
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: addr,
		Auth: clickhouse.Auth{
			Database: database,
			Username: username,
		},
		Protocol: clickhouse.HTTP,
	})
	err := conn.Ping()
	if err != nil {
		return nil, err
	}
	ch.DB = conn
	return ch, nil
}

// Close
func (ch *ClickHouse) Close() {
	ch.DB.Close()
}
