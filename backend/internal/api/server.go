package api

import (
	"context"
	"net/http"
	"shift/internal/entity"
	"shift/internal/service"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/gorilla/mux"
	"github.com/heptiolabs/healthcheck"

	"github.com/sirupsen/logrus"
)

// APIServer represents an HTTP server for the JSON API.
type APIServer struct {
	address       string
	userDB        entity.UserDB
	associationDB entity.AssociationDB
	invitationDB  entity.InvitationDB
	bucketDb      entity.BucketDB
	userService   *service.UserService
}

// NewAPIServer creates a new instance of APIServer with the given address.
func NewAPIServer(
	address string,
	bucketDb entity.BucketDB,
	userDB entity.UserDB,
) *APIServer {
	return &APIServer{
		address:     address,
		userDB:      userDB,
		bucketDb:    bucketDb,
		userService: service.NewUserService(bucketDb, userDB),
	}
}

// Run starts the HTTP server and listens for incoming requests.
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.Use(CORSMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))

	http.Handle("/live", health.NewHandler(
		health.NewChecker(
			health.WithCheck(health.Check{
				Name: "database",
				Check: func(ctx context.Context) error {
					deadline, _ := ctx.Deadline()
					timeout := time.Since(deadline)
					return healthcheck.DatabasePingCheck(nil, timeout)()
				},
			}),
		)))

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	s.initUserRoutes(apiRouter)
	s.initInvitaionRoutes(apiRouter)

	// TODO: temporary, only to demonstrate the authorization abilities - delete it and the handlers later.
	s.initAuthorizationRoutes(apiRouter.PathPrefix("/authorization").Subrouter())

	logrus.Println("JSON API Server is running on port", s.address)
	logrus.Fatal(http.ListenAndServe(s.address, router))
}
