package api

import (
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"log"
	"net/http"
	"shift/internal/entity"
	"shift/internal/service"
	cauth "shift/pkg/auth"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer represents an HTTP server for the JSON API.
type APIServer struct {
	address            string
	jwtValidator       *validator.Validator
	userService        *service.UserService
	associationService *service.AssociationService
	invitationService  *service.InvitationService
}

// NewAPIServer creates a new instance of APIServer with the given address.
func NewAPIServer(
	address string,
	bucketDB entity.BucketDB,
	userDB entity.UserDB,
	associationDB entity.AssociationDB,
	invitationDB entity.InvitationDB,
) *APIServer {
	jwtValidator, err := cauth.JwtValidator()
	if err != nil {
		log.Fatalf("initializing jwt validator: %v", err)
	}

	return &APIServer{
		address:            address,
		jwtValidator:       jwtValidator,
		userService:        service.NewUserService(bucketDB, userDB, invitationDB, associationDB),
		associationService: service.NewAssociationService(bucketDB, associationDB),
		invitationService:  service.NewInvitationService(bucketDB, invitationDB),
	}
}

// Run starts the HTTP server and listens for incoming requests.
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.Use(CORSMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	s.initUserRoutes(apiRouter)
	s.initProfileRoutes(apiRouter)
	s.initAssociationRoutes(apiRouter)
	s.initInvitationRoutes(apiRouter)

	// TODO: temporary, only to demonstrate the authorization abilities - delete it and the handlers later.
	s.initAuthorizationRoutes(apiRouter.PathPrefix("/authorization").Subrouter())

	router.PathPrefix("").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}).
		Methods(http.MethodOptions)

	logrus.Println("JSON API Server is running on port", s.address)
	logrus.Fatal(http.ListenAndServe(s.address, router))
}
