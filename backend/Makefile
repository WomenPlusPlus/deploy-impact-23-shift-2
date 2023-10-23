.PHONY: all build run docker-build docker-run clean

BINARY_NAME=shift
PORT=8080
DATA_PATH=data
VERSION=0.0.2
POSTGRESQL_URL=postgres://postgres:shift2023@0.0.0.0:5432/postgres?sslmode=disable

build:
	@go build -o bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)
run: build
	@go run ./cmd/$(BINARY_NAME)
docker-build:
	@docker build -t $(BINARY_NAME) .
docker-run:
	docker run -p $(PORT):$(PORT) $(BINARY_NAME)

docker-migrate:
	migrate -database ${POSTGRESQL_URL} -path internal/db/migrations up

docker-run-db:
	mkdir -p $(DATA_PATH) && docker run -d --rm --name shift2-db -v $(DATA_PATH):/var/lib/postgresql -p 5432:5432 -e POSTGRES_PASSWORD=shift2023 postgres
docker-pull:
	docker pull shift2/server:$(VERSION)
docker-push:
	docker push shift2/server:$(VERSION)
clean:
	rm -f bin/$(BINARY_NAME)