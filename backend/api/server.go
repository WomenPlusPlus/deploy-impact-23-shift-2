package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shift/entity"
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

// WriteJSON writes a JSON response with the given status code and value.
func WriteJSON(w http.ResponseWriter, status int, value interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
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
				WriteJSON(w, http.StatusNotFound, apiError{Error: "Resource not found"})
			case IsPermissionError(err):
				WriteJSON(w, http.StatusForbidden, apiError{Error: "Permission denied"})
			default:
				// Log the internal error without exposing details to the client
				logger.Error(err)
				WriteJSON(w, http.StatusInternalServerError, apiError{Error: "Internal server error"})
			}
		}
	}
}

// IsNotFoundError checks if an error is a not found error.
func IsNotFoundError(err error) bool {
	// Implement your custom logic to check for not found errors
	return false
}

// IsPermissionError checks if an error is a permission error.
func IsPermissionError(err error) bool {
	// Implement your custom logic to check for permission errors
	return false
}

// Run starts the HTTP server and listens for incoming requests.
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	// Users
	router.HandleFunc("/users", makeHTTPHandleFunc(s.handleUsers)).
		Methods(http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/users/{id}", makeHTTPHandleFunc(s.handleGetUserByID)).
		Methods(http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions)
	// Admin
	// router.HandleFunc("/admin/invites", makeHTTPHandleFunc(s.handleAdminInvites))
	// router.HandleFunc("/admin/associations", makeHTTPHandleFunc(s.handleAdminAssociations))
	// router.HandleFunc("/admin/companies", makeHTTPHandleFunc(s.handleAdmin))
	// router.HandleFunc("/admin/training", makeHTTPHandleFunc(s.handleAdmin))
	// router.HandleFunc("/admin/help", makeHTTPHandleFunc(s.handleAdmin))
	// Associations Admin
	// router.HandleFunc("/associations-admin/dashboard", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/associations-admin/my-association", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/associations-admin/my-association/view", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/associations-admin/my-association/edit", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/associations-admin/my-association/initiatives", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/associations-admin/my-association/users", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/associations-admin/my-profile/details", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/associations-admin/my-profile/preferences", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/associations-admin/invites", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/associations-admin/requests", makeHTTPHandleFunc(s.handleUsers))
	// // Associations User
	// router.HandleFunc("/associations-user/dashboard", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/association-user/my-association", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/association-user/my-profile/details", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/association-user/my-profile/preferences", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/association-user/invites", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/association-user/task-center", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/association-user/task-center/todo-lists", makeHTTPHandleFunc(s.handleUsers))
	// // Candidates
	// router.HandleFunc("/candidates/search/matching", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/candidates/search/jobs", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/candidates/search/companies", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/candidates/search/saved", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/candidates/my-profile/details", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/candidates/my-profile/data-privacy", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/candidates/my-profile/preferences", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/candidates/task-center/alerts", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/candidates/task-center/todo-lists", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/candidates/extras/personality-tests", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/candidates/extras/cognitive-tests", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/candidates/extras/tutorial", makeHTTPHandleFunc(s.handleUsers))
	// // Company Admin
	// router.HandleFunc("/company-admin/profile/details", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-admin/profile/preferences", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-admin/my-company", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-admin/job-listings/my-listings", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-admin/job-listings/company-listings", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-admin/job-listings/candidates", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-admin/search/candidates", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-admin/search/saved", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-admin/task-center/alerts", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-admin/task-center/todo-lists", makeHTTPHandleFunc(s.handleUsers))
	// // Company User
	// router.HandleFunc("/company-user/profile/details", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-user/profile/preferences", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-user/my-company", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-user/job-listings/company-listings", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-user/job-listings/my-listings", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-user/job-listings/candidates", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-user/search/candidates", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-user/search/saved", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-user/task-center/alerts", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/company-user/task-center/todo-lists", makeHTTPHandleFunc(s.handleUsers))

	// Test paths
	// router.HandleFunc("/users", makeHTTPHandleFunc(s.handleUsers))
	// router.HandleFunc("/users/{id}", makeHTTPHandleFunc(s.handleGetUserByID))

	log.Println("JSON API Server is running on port:", s.address)

	http.ListenAndServe(s.address, router)
}

// func (s *APIServer) handleAdminInvites(w http.ResponseWriter, r *http.Request) error {
// 	if r.Method == "GET" {
// 		return WriteJSON(w, http.StatusOK, "admin invites")
// 	}
// 	return fmt.Errorf("method not allowed %s", r.Method)
// }

// func (s *APIServer) handleAdminAssociations(w http.ResponseWriter, r *http.Request) error {
// 	if r.Method == "GET" {
// 		return WriteJSON(w, http.StatusOK, "admin associations")
// 	}
// 	return fmt.Errorf("method not allowed %s", r.Method)
// }

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

	return WriteJSON(w, http.StatusOK, users)
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

	return WriteJSON(w, http.StatusOK, user)
}

// handleCreateUser handles POST requests to create a new user account.
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
	return WriteJSON(w, http.StatusOK, user)
}

// handleDeleteUser handles DELETE requests to delete a user account.
func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// handleUpdateUser handles PUT requests to update a user account.
// func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }
