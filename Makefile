export PORT := 8081

.PHONY: dev
dev:
	@go run ./cmd/api
