package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"shift/internal/api"
	"shift/internal/db"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("Setup log level")

	if err := godotenv.Overload(); err != nil {
		logrus.Fatalf("Error loading the .env file: %v", err)
	}
	logrus.Trace("Setup environment variables")

	ctx := context.Background()

	postgresDB := db.NewPostgresDB()
	postgresDB.Init()
	logrus.Tracef("PostgreSQL DB initialized: %T", postgresDB)

	bucketDB := db.NewGoogleBucketDB(ctx)
	logrus.Tracef("GCP Bucket initialized: %T", bucketDB)

	server := api.NewAPIServer(":"+os.Getenv("API_PORT"), bucketDB, postgresDB)
	logrus.Tracef("Server initialized: %T", server)
	server.Run()
}
