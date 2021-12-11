
test:
	@echo "[go test] running tests and collecting coverage metrics"
	@go test -v -tags all_tests -race -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	@echo "[golangci-lint] linting sources"
	@golangci-lint run ./...

.PHONY: build
go-build: lint test
	@echo "  >  Building binary..."
	@go build ./cmd/app/main.go