package main

import (
	"shift/internal/api"
	"shift/internal/db"
)

func main() {
	userDB := db.NewPostgresDB()

	server := api.NewAPIServer(":8080", userDB)
	server.Run()
}
