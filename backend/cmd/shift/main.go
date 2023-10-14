package main

import (
	"shift/internal/api"
	"shift/internal/db"
)

func main() {
	db := db.NewPostgresDB()
	db.Init()

	server := api.NewAPIServer(":8080", db)
	server.Run()
}
