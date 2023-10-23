package main

import (
	"context"
	"os"
	"shift/internal/api"
	"shift/internal/db"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Overload(); err != nil {
		logrus.Fatalf("Error loading the .env file: %v", err)
	}
	logrus.Trace("Setup environment variables")

	if os.Getenv("ENVIRONMENT") == "prod" {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.Info("Setup log level to INFO")
	} else {
		logrus.SetLevel(logrus.TraceLevel)
		logrus.Info("Setup log level to TRACE")
	}

	ctx := context.Background()

	postgresDB := db.NewPostgresDB()
	logrus.Tracef("PostgreSQL DB initialized: %T", postgresDB)

	bucketDB := db.NewGoogleBucketDB(ctx)
	logrus.Tracef("GCP Bucket initialized: %T", bucketDB)

	server := api.NewAPIServer(":"+os.Getenv("API_PORT"), bucketDB, postgresDB)
	logrus.Tracef("Server initialized: %T", server)
	server.Run()
}
