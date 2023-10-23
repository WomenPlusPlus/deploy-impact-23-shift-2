package api

import (
	"net/http"
	"shift/internal/entity"
	"shift/internal/service"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer represents an HTTP server for the JSON API.
type APIServer struct {
<<<<<<< Updated upstream
	address       string
	userDB        entity.UserDB
	associationDB entity.AssociationDB
	invitationDB  entity.InvitationDB
	bucketDb      entity.BucketDB
	userService   *service.UserService
=======
	address string
	// invitationDB       entity.InvitationDB
	bucketDb           entity.BucketDB
	userService        *service.UserService
	associationService *service.AssociationService
>>>>>>> Stashed changes
}

// NewAPIServer creates a new instance of APIServer with the given address.
func NewAPIServer(
	address string,
	bucketDb entity.BucketDB,
	userDB entity.UserDB,
) *APIServer {
	return &APIServer{
<<<<<<< Updated upstream
		address:     address,
		userDB:      userDB,
		bucketDb:    bucketDb,
		userService: service.NewUserService(bucketDb, userDB),
=======
		address:            address,
		bucketDb:           bucketDB,
		userService:        service.NewUserService(bucketDB, postgresDB),
		associationService: service.NewAssociationService(bucketDB, postgresDB),
>>>>>>> Stashed changes
	}
}

// Run starts the HTTP server and listens for incoming requests.
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.Use(CORSMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	s.initUserRoutes(apiRouter)
<<<<<<< Updated upstream
	s.initInvitaionRoutes(apiRouter)
=======
	s.initAssociationRoutes(apiRouter)
>>>>>>> Stashed changes

	// TODO: temporary, only to demonstrate the authorization abilities - delete it and the handlers later.
	s.initAuthorizationRoutes(apiRouter.PathPrefix("/authorization").Subrouter())

	logrus.Println("JSON API Server is running on port", s.address)
	logrus.Fatal(http.ListenAndServe(s.address, router))
}
