export PORT := 8081

.PHONY: run
run:
	@go run ./cmd/api

# To use this command, run `go install github.com/cespare/reflex@latest`
.PHONY: dev
dev:
	@reflex -r '\.go$$' -s make run

