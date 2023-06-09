.PHONY: build
build:
	go build -v ./cmd/subscriber
	go build -v ./cmd/publisher
	docker compose up

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: clean
clean:
	rm -f publisher
	rm -f subscriber

.PHONY: fclean
fclean:
	make clean
	docker compose down
	docker image prune
	docker volume prune

.DEFAULT_GOAL := build