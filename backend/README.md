# SHIFT2 backend

# Setting Up a Go Development Environment with Docker and Make

This guide will help you set up a Go development environment using Visual Studio Code, Docker, and Make. It includes instructions for installing and configuring Go, Docker, creating a Makefile, and setting up a PostgreSQL container.

## Table of Contents

- [Install and Configure Go in VSCode](#install-and-configure-go-in-vscode)
- [Install Docker](#install-docker)
- [Create a Makefile](#create-a-makefile)
- [Install PostgreSQL with Docker](#install-postgresql-with-docker)
- [Run a Go Program using Make](#run-a-go-program-using-make)

## Install and Configure Go in VSCode

1. Install Go on your system by following the official installation guide: [https://golang.org/doc/install](https://golang.org/doc/install)

2. Open Visual Studio Code (VSCode).

3. Install the Go extension for VSCode by navigating to the Extensions view (`Ctrl+Shift+X`), searching for "Go," and clicking the Install button for the official Go extension.

4. Create a new Go project or open an existing one in VSCode.

5. VSCode should automatically detect your Go installation and offer to install any necessary tools. Follow the prompts to install the required tools like `gopls` (Go Language Server) for code navigation and autocompletion.

6. Configure your Go workspace settings in VSCode, if needed. You can access workspace settings by clicking on the gear icon in the bottom-right corner and selecting "Settings." Customize settings like the Go `GOROOT` and `GOPATH` as per your project requirements.

## Install Docker

1. Install Docker on your system by following the official installation guide for your platform: [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)

## Install PostgreSQL with Docker

1. Pull the PostgreSQL Docker image and run a container with the following commands:

```
docker pull postgres
docker run --name shift-postgres -e POSTGRES_PASSWORD=shift2023 -d -p 5432:5432 postgres
```

2. Your PostgreSQL database is now running in a Docker container. You can connect to it from your Go application using the appropriate connection string.

## Run a Go Program using Make

1. To build your Go application, open a terminal and navigate to your project directory, then run:

```
make build
```

2. To run your Go application:

```
make run
```

3. To build a Docker image of your Go application:

```
make docker-build
```

4. To run your Go application in a Docker container:

```
make docker-run
```

This will start your Go application in a container and expose it on the specified port (8080 in this example).

That's it! You've successfully set up a Go development environment with Docker and Make.