.PHONY: all build run docker-build docker-run clean

BINARY_NAME=shift
PORT=8080

build:
	@go build -o bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)
run: build
	@go run ./cmd/$(BINARY_NAME)
docker-build:
	@docker build -t $(BINARY_NAME) .
docker-run:
	docker run -p $(PORT):$(PORT) $(BINARY_NAME)
clean:
	rm -f bin/$(BINARY_NAME)
