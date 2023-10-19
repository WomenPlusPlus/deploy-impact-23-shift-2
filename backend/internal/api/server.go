package api

import (
	"fmt"
	"net/http"
	"shift/internal/entity"
	"shift/internal/service"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer represents an HTTP server for the JSON API.
type APIServer struct {
	address            string
	userDB             entity.UserDB
	associationDB      entity.AssociationDB
	bucketDb           entity.BucketDB
	userService        *service.UserService
	associationService *service.AssociationService
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

	router.HandleFunc("/admin/users", makeHTTPHandleFunc(s.handleUsers))
	router.HandleFunc("/admin/users/{id}", makeHTTPHandleFunc(s.handleGetUserByID))
	router.HandleFunc("/admin/users/delete/{id}", makeHTTPHandleFunc(s.handleDeleteUser))

	router.HandleFunc("/admin/associations", makeHTTPHandleFunc(s.handleCreateAssociation))

	logrus.Println("JSON API Server is running on port", s.address)
	logrus.Fatal(http.ListenAndServe(s.address, router))
}

// handleUsers handles requests related to user accounts.
func (s *APIServer) handleUsers(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetUsers(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteUser(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

// handleGetUser handles GET requests for user account information.
func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := s.userDB.GetUsers()
	if err != nil {
		return err
	}
	return WriteJSONResponse(w, http.StatusOK, users)
}

func (s *APIServer) handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	user, err := s.userDB.GetUserByID(id)

	if err != nil {
		return NotFoundError{Message: "User not found"}
	}

	return WriteJSONResponse(w, http.StatusOK, user)
}

// handleDeleteUser handles DELETE requests to delete a user account.
func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	if _, err := s.userDB.GetUserByID(id); err != nil {
		return WriteJSONResponse(w, http.StatusNotFound, apiError{Error: err.Error()})
	}

	return WriteJSONResponse(w, http.StatusOK, "User deleted successfully")
}

// handleUpdateUser handles PUT requests to update a user account.
// func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }