package main

import (
	"os"
	"shift/internal/api"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("Setup log level")

	if err := godotenv.Overload(); err != nil {
		logrus.Fatalf("Error loading the .env file: %v", err)
	}
	logrus.Trace("Setup environment variables")

	server := api.NewAPIServer(":" + os.Getenv("API_PORT"))
	logrus.Tracef("Server initialized: %T", server)
	server.Run()
}
