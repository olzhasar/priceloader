test:
	go test -v ./...

build:
	go build -o priceloader ./cmd/...

.PHONY: test build
