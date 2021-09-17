export PORT := 8081
export MIGRATION_PATH := db/migrations
export DEV_DATABASE_PATH := postgresql://postgres:postgres@localhost:5433/sandbox?sslmode=disable

run:
	go run ./cmd/api

dev:
	reflex -r '\.go$$' -s make run

postgres:
	docker-compose up -d

sqlboiler:
	go generate

migrateup:
	migrate -path $(MIGRATION_PATH) -database $(DEV_DATABASE_PATH) -verbose up

migrateup1:
	migrate -path $(MIGRATION_PATH) -database $(DEV_DATABASE_PATH) -verbose up 1

migratedown:
	migrate -path $(MIGRATION_PATH) -database $(DEV_DATABASE_PATH) -verbose down

migratedown1:
	migrate -path $(MIGRATION_PATH) -database $(DEV_DATABASE_PATH) -verbose down 1

migrateredo:
	make migratedown1 && make migrateup1

.PHONY: run dev postgres sqlboiler migratedown migratedown1 migrateup migrateup1 migrateredo
