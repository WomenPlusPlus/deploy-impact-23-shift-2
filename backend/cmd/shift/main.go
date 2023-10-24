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
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("Setup log level")

	ctx := context.Background()

	bucketDB := db.NewGoogleBucketDB(ctx)
	logrus.Tracef("GCP Bucket initialized: %T", bucketDB)

	postgresDB := db.NewPostgresDB()
	logrus.Tracef("PostgreSQL DB initialized: %T", postgresDB)

	if err := godotenv.Overload(); err != nil {
		logrus.Fatalf("Error loading the .env file: %v", err)
	}
	logrus.Trace("Setup environment variables")

	server := api.NewAPIServer(":"+os.Getenv("API_PORT"), bucketDB, postgresDB, postgresDB, postgresDB)
	logrus.Tracef("Server initialized: %T", server)
	server.Run()
}
