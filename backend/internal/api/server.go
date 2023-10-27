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
	address        string
	userService    *service.UserService
	companyService *service.CompanyService
	jobService     *service.JobService
	userDB         entity.UserDB
	companyDB      entity.CompanyDB
	invitationDB   entity.InvitationDB
	jobDB          entity.JobDB
}

// NewAPIServer creates a new instance of APIServer with the given address.
func NewAPIServer(
	address string,
	bucketDB entity.BucketDB,
	userDB entity.UserDB,
	companyDB entity.CompanyDB,
	invitationDB entity.InvitationDB,
	jobDB entity.JobDB,

) *APIServer {
	return &APIServer{
		address:        address,
		userService:    service.NewUserService(bucketDB, userDB),
		companyService: service.NewCompanyService(bucketDB, companyDB),
		jobService:     service.NewJobService(bucketDB, jobDB),
		userDB:         userDB,
		companyDB:      companyDB,
		invitationDB:   invitationDB,
		jobDB:          jobDB,
	}
}

// Run starts the HTTP server and listens for incoming requests.
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.Use(CORSMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	s.initUserRoutes(apiRouter)
	s.initInvitaionRoutes(apiRouter)
	s.initCompanyRoutes(apiRouter)

	s.initAuthorizationRoutes(apiRouter.PathPrefix("/authorization").Subrouter())

	logrus.Println("JSON API Server is running on port", s.address)
	logrus.Fatal(http.ListenAndServe(s.address, router))
}
