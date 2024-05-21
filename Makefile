require: ### download requirements
	go mod download
.PHONY: require

run: ### run application
	go run cmd/app/main.go
.PHONY: run

migrate-create:  ### create new migration
	goose -dir migrations create init sql
.PHONY: migrate-create

migrate-up: ### migration up
	GOOSE_DRIVER=clickhouse GOOSE_MIGRATION_DIR=./migrations goose "http://localhost:8123" up
.PHONY: migrate-up

migrate-down: ### migration down
	GOOSE_DRIVER=clickhouse GOOSE_MIGRATION_DIR=./migrations goose "http://localhost:8123" down
.PHONY: migrate-down

compose-up: ### run docker-compose
	docker-compose up --build -d clickhouse && docker-compose logs -f
.PHONY: compose-up

compose-down: ### down docker-compose
	docker-compose down --remove-orphans
.PHONY: compose-down

swag-init: ### swag init
	swag init -g cmd/app/main.go
.PHONY: swag-init

mock: ### run mockgen
	mockgen -source ./internal/usecase/interfaces.go -package usecase_test > ./internal/usecase/mocks_test.go
.PHONY: mock

test: ### run test
	go test -v -cover -race ./internal/...
.PHONY: test
