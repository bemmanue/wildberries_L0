.PHONY: build
build:
	go build -v ./cmd/subscriber
	go build -v ./cmd/publisher

.PHONY: test
test:
	go test -v -race -timeout 30s ./...


.DEFAULT_GOAL := build