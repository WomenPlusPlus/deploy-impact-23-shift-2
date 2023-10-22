.PHONY: all build run docker-build docker-run clean

BINARY_NAME=shift
PORT=8080
DATA_PATH=data

build:
	@go build -o bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)
run: build
	@go run ./cmd/$(BINARY_NAME)
docker-build:
	@docker build -t $(BINARY_NAME) .
docker-run:
	docker run -p $(PORT):$(PORT) $(BINARY_NAME)
docker-run-db:
	mkdir -p $(DATA_PATH) && docker run -d --rm --name shift2-db -v $(DATA_PATH):/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=shift2023 postgres
clean:
	rm -f bin/$(BINARY_NAME)
