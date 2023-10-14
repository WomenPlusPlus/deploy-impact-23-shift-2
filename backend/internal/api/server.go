package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shift/internal/entity"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer represents an HTTP server for the JSON API.
type APIServer struct {
	address string
	userDB  entity.UserDB
}

// NewAPIServer creates a new instance of APIServer with the given address.
func NewAPIServer(address string, userDB entity.UserDB) *APIServer {
	return &APIServer{
		address: address,
		userDB:  userDB,
	}
}

// apiFunc represents a function that handles API requests.
type apiFunc func(http.ResponseWriter, *http.Request) error

// apiError represents an API error response.
type apiError struct {
	Error string `json:"error"`
}

// Run starts the HTTP server and listens for incoming requests.
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))

	router.HandleFunc("/admin/users", makeHTTPHandleFunc(s.handleUsers))
	router.HandleFunc("/admin/users/{id}", makeHTTPHandleFunc(s.handleGetUserByID))

	log.Println("JSON API Server is running on port", s.address)
	http.ListenAndServe(s.address, router)
}

// IsNotFoundError checks if an error is a not found error.
func IsNotFoundError(err error) bool {
	return false
}

// IsPermissionError checks if an error is a permission error.
func IsPermissionError(err error) bool {
	// Implement your custom logic to check for permission errors
	return false
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
				WriteJSONResponse(w, http.StatusNotFound, apiError{Error: "Resource not found"})
			case IsPermissionError(err):
				WriteJSONResponse(w, http.StatusForbidden, apiError{Error: "Permission denied"})
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
	return fmt.Errorf("method not allowed %s", r.Method)
}

// handleGetUser handles GET requests for user account information.
func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := s.userDB.GetUsers()
	fmt.Println(users)
	if err != nil {
		return err
	}
	return WriteJSONResponse(w, http.StatusOK, users)
}

func (s *APIServer) handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return fmt.Errorf("invalid id given %s", idStr)
	}

	user, err := s.userDB.GetUserByID(id)

	if err != nil {
		return err
	}

	return WriteJSONResponse(w, http.StatusOK, user)
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	userRequest := new(entity.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(userRequest); err != nil {
		return err
	}

	user := entity.NewUser(
		userRequest.FirstName,
		userRequest.LastName,
		userRequest.PreferredName,
		userRequest.Email,
		userRequest.State,
		userRequest.ImageUrl,
		userRequest.Role,
	)

	if err := s.userDB.CreateUser(user); err != nil {
		return err
	}

	return WriteJSONResponse(w, http.StatusOK, user)
}

// func (s *APIServer) handleAdminInvites(w http.ResponseWriter, r *http.Request) error {
// 	if r.Method == "GET" {
// 		return WriteJSONResponse(w, http.StatusOK, "admin invites")
// 	}
// 	return fmt.Errorf("method not allowed %s", r.Method)
// }

// func (s *APIServer) handleAdminAssociations(w http.ResponseWriter, r *http.Request) error {
// 	if r.Method == "GET" {
// 		return WriteJSONResponse(w, http.StatusOK, "admin associations")
// 	}
// 	return fmt.Errorf("method not allowed %s", r.Method)
// }

// handleDeleteUser handles DELETE requests to delete a user account.
// func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }

// handleUpdateUser handles PUT requests to update a user account.
// func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }
