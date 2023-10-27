FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg

RUN CGO_ENABLED=0 GOOS=linux go build -o /shift ./cmd/shift

EXPOSE 8080

# Run
CMD ["/shift"]
