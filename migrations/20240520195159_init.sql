-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS default.events (
    eventID Int64 DEFAULT rand(),
    eventType String,
    userID Int64,
    eventTime DateTime,
    payload String
) ENGINE = MergeTree
ORDER BY (eventID, eventTime);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS default.events;
-- +goose StatementEnd
