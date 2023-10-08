package main

import (
	"log"
)

func main() {
	userDB, err := NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := userDB.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", userDB)
	server.Run()
}
