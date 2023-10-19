package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"shift/internal/entity"
	"shift/internal/service"
)

// APIServer represents an HTTP server for the JSON API.
type APIServer struct {
	address       string
	userDB        entity.UserDB
	associationDB entity.AssociationDB
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

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	s.initUserRoutes(apiRouter)

	// TODO: temporary, only to demonstrate the authorization abilities - delete it and the handlers later.
	s.initAuthorizationRoutes(apiRouter.PathPrefix("/authorization").Subrouter())

	router.HandleFunc("/admin/associations", makeHTTPHandleFunc(s.handleCreateAssociation))

	logrus.Println("JSON API Server is running on port", s.address)
	logrus.Fatal(http.ListenAndServe(s.address, router))
}

func (s *APIServer) handleCreateAssociation(w http.ResponseWriter, r *http.Request) error {
	associationRequest := new(entity.CreateAssociationRequest)
	if err := json.NewDecoder(r.Body).Decode(associationRequest); err != nil {
		return err
	}

	ass := entity.NewAssociation(
		associationRequest.Name,
		associationRequest.Logo,
		associationRequest.WebsiteUrl,
		associationRequest.Focus,
	)

	if err := s.associationDB.CreateAssociation(ass); err != nil {
		return WriteJSONResponse(w, http.StatusNotFound, apiError{Error: err.Error()})
	}

	return WriteJSONResponse(w, http.StatusOK, ass)
}
