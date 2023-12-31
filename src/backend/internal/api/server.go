package api

import (
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"log"
	"net/http"
	"os"
	"shift/internal/entity"
	"shift/internal/service"
	cauth "shift/pkg/auth"
	golocation "shift/pkg/go-location"

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
	companyService     *service.CompanyService
	jobService         *service.JobService
}

// NewAPIServer creates a new instance of APIServer with the given address.
func NewAPIServer(
	address string,
	bucketDB entity.BucketDB,
	userDB entity.UserDB,
	associationDB entity.AssociationDB,
	invitationDB entity.InvitationDB,
	companyDB entity.CompanyDB,
	jobDB entity.JobDB,
) *APIServer {
	jwtValidator, err := cauth.JwtValidator()
	if err != nil {
		log.Fatalf("initializing jwt validator: %v", err)
	}

	associationService := service.NewAssociationService(bucketDB, associationDB)
	invitationService := service.NewInvitationService(bucketDB, invitationDB)
	userService := service.NewUserService(bucketDB, userDB)
	companyService := service.NewCompanyService(bucketDB, companyDB)
	jobService := service.NewJobService(bucketDB, jobDB)

	associationService.Inject(userService)
	userService.Inject(invitationService, associationService)
	companyService.Inject(userService)
	jobService.Inject(userService, companyService)

	return &APIServer{
		address:            address,
		jwtValidator:       jwtValidator,
		userService:        userService,
		associationService: associationService,
		invitationService:  invitationService,
		companyService:     companyService,
		jobService:         jobService,
	}
}

// Run starts the HTTP server and listens for incoming requests.
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.Use(CORSMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))

	if os.Getenv("LOAD_LOCATIONS_API") == "true" {
		logrus.Infof("Loading locations API locally")
		golocation.InitRoutes(router)
	}

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	s.initUserRoutes(apiRouter)
	s.initProfileRoutes(apiRouter)
	s.initAssociationRoutes(apiRouter)
	s.initInvitationRoutes(apiRouter)
	s.initCompanyRoutes(apiRouter)
	s.initJobRoutes(apiRouter)

	router.PathPrefix("").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}).
		Methods(http.MethodOptions)

	logrus.Println("JSON API Server is running on port", s.address)
	logrus.Fatal(http.ListenAndServe(s.address, router))
}
