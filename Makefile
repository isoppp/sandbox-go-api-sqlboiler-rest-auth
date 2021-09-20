export PORT := 8081
export _MIGRATION_DIR := db/migrations
export _MIGRATION_DB_PATH := postgresql://postgres:postgres@localhost:5433/sandbox?sslmode=disable

run:
	go run ./cmd/api

dev:
	go run ./cmd/api -dev

devw:
	reflex -r '\.go$$' -s make dev

postgres:
	docker-compose up -d

sqlboiler:
	go generate

migrateup:
	migrate -path $(_MIGRATION_DIR) -database $(_MIGRATION_DB_PATH) -verbose up

migrateup1:
	migrate -path $(_MIGRATION_DIR) -database $(_MIGRATION_DB_PATH) -verbose up 1

migratedown:
	migrate -path $(_MIGRATION_DIR) -database $(_MIGRATION_DB_PATH) -verbose down

migratedown1:
	migrate -path $(_MIGRATION_DIR) -database $(_MIGRATION_DB_PATH) -verbose down 1

migrateredo:
	make migratedown1 && make migrateup1

test:
	go test -v -cover ./...

.PHONY: run dev devw postgres sqlboiler migratedown migratedown1 migrateup migrateup1 migrateredo test
