test:
	go test -v ./...

build:
	go build -o bin/ ./cmd/...

.PHONY: test build
