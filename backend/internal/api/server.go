package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shift/internal/entity"
	"shift/internal/service"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer represents an HTTP server for the JSON API.
type APIServer struct {
	address     string
	userDB      entity.UserDB
	userService *service.UserService
}

// NewAPIServer creates a new instance of APIServer with the given address.
func NewAPIServer(address string, userDB entity.UserDB) *APIServer {
	return &APIServer{
		address:     address,
		userDB:      userDB,
		userService: service.NewUserService(userDB),
	}
}

// apiFunc represents a function that handles API requests.
type apiFunc func(http.ResponseWriter, *http.Request) error

// apiError represents an API error response.
type apiError struct {
	Error string `json:"error"`
}

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

type PermissionError struct {
	Message string
}

func (e PermissionError) Error() string {
	return e.Message
}

type BadRequestError struct {
	Message string
}

func (e BadRequestError) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string
}

func (e InternalServerError) Error() string {
	return e.Message
}

// Run starts the HTTP server and listens for incoming requests.
func (s *APIServer) Run() {
	logrus.SetLevel(logrus.TraceLevel)

	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))

	router.HandleFunc("/admin/users", makeHTTPHandleFunc(s.handleUsers))
	router.HandleFunc("/admin/users/create", makeHTTPHandleFunc(s.handleCreateUser))

	router.HandleFunc("/admin/users/{id}", makeHTTPHandleFunc(s.handleGetUserByID))
	router.HandleFunc("/admin/users/delete/{id}", makeHTTPHandleFunc(s.handleDeleteUser))

	log.Println("JSON API Server is running on port", s.address)
	http.ListenAndServe(s.address, router)
}

// IsNotFoundError checks if an error is a not found error.
func IsNotFoundError(err error) bool {
	_, isNotFound := err.(NotFoundError)
	return isNotFound
}

// IsPermissionError checks if an error is a permission error.
func IsPermissionError(err error) bool {
	_, isPermissionDenied := err.(PermissionError)
	return isPermissionDenied
}

func IsBadRequestError(err error) bool {
	_, isBadRequestError := err.(BadRequestError)
	return isBadRequestError
}

func IsInternalServerError(err error) bool {
	_, isInternalServerError := err.(InternalServerError)
	return isInternalServerError
}

// makeHTTPHandleFunc creates an HTTP request handler function for the provided apiFunc.
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{}) // Use JSON format for structured logs
		err := f(w, r)
		if err != nil {
			switch {
			case IsNotFoundError(err):
				WriteJSONResponse(w, http.StatusNotFound, apiError{Error: err.Error()})
			case IsPermissionError(err):
				WriteJSONResponse(w, http.StatusForbidden, apiError{Error: err.Error()})
			case IsBadRequestError(err):
				WriteJSONResponse(w, http.StatusBadRequest, apiError{Error: err.Error()})
			case IsInternalServerError(err):
				WriteJSONResponse(w, http.StatusInternalServerError, apiError{Error: err.Error()})
			default:
				// Log the internal error without exposing details to the client
				logger.Error(err)
				WriteJSONResponse(w, http.StatusInternalServerError, apiError{Error: "Internal server error"})
			}
		}
	}
}

// WriteJSONResponse writes a JSON response with the given status code and value.
func WriteJSONResponse(w http.ResponseWriter, status int, value interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

// handleUsers handles requests related to user accounts.
func (s *APIServer) handleUsers(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetUsers(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateUser(w, r)
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

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("Create user handler running")

	req := new(entity.CreateUserRequest)
	if err := req.FromFormData(r); err != nil {
		return BadRequestError{Message: err.Error()}
	}

	res, err := s.userService.CreateUser(req)
	if err != nil {
		return InternalServerError{Message: err.Error()}
	}

	return WriteJSONResponse(w, http.StatusOK, res)
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
