package api

import (
	"context"
	"net/http"
	"shift/internal/db"
	"shift/internal/entity"
	"shift/internal/service"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer represents an HTTP server for the JSON API.
type APIServer struct {
	address string
	// invitationDB       entity.InvitationDB
	bucketDb           entity.BucketDB
	userService        *service.UserService
	associationService *service.AssociationService
}

// NewAPIServer creates a new instance of APIServer with the given address.
func NewAPIServer(
	address string,
) *APIServer {
	ctx := context.Background()

	bucketDB := db.NewGoogleBucketDB(ctx)
	logrus.Tracef("GCP Bucket initialized: %T", bucketDB)

	postgresDB := db.NewPostgresDB()
	logrus.Tracef("PostgreSQL DB initialized: %T", postgresDB)

	return &APIServer{
		address:            address,
		bucketDb:           bucketDB,
		userService:        service.NewUserService(bucketDB, postgresDB),
		associationService: service.NewAssociationService(bucketDB, postgresDB),
	}
}

// Run starts the HTTP server and listens for incoming requests.
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.Use(CORSMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	s.initUserRoutes(apiRouter)
	s.initAssociationRoutes(apiRouter)

	// TODO: temporary, only to demonstrate the authorization abilities - delete it and the handlers later.
	s.initAuthorizationRoutes(apiRouter.PathPrefix("/authorization").Subrouter())

	logrus.Println("JSON API Server is running on port", s.address)
	logrus.Fatal(http.ListenAndServe(s.address, router))
}
