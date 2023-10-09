# Makefile for Go Project

# Declare phony targets to avoid conflicts with file names.
.PHONY: all build run docker-build docker-run clean

# Name of the Go binary (executable) and the port your application will run on.
BINARY_NAME=shift
PORT=8080

# Default target. It will build the project by default.
all: build

# Target to build the Go binary.
build:
	# Use 'go build' to compile the Go code into an executable binary.
	@go build -o bin/$(BINARY_NAME)

# Target to run the Go application.
run: build
	# Run the compiled binary.
	@./bin/$(BINARY_NAME)

# Target to build a Docker image of the Go application.
docker-build:
	# Use Docker to build an image tagged with $(BINARY_NAME).
	docker build -t $(BINARY_NAME) .

# Target to run the Go application in a Docker container.
docker-run:
	# Run a Docker container, mapping the specified port inside the container to the same port on your host.
	docker run -p $(PORT):$(PORT) $(BINARY_NAME)

# Target to clean up generated files.
clean:
	# Remove the compiled binary.
	rm -f bin/$(BINARY_NAME)
