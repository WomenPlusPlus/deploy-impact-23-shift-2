package main

import (
	"shift/api"
	"shift/db"
)

func main() {
	userDB := db.NewPostgresDB()

	server := api.NewAPIServer(":8080", userDB)
	server.Run()
}
