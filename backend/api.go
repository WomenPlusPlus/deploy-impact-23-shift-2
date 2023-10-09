package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// APIServer represents an HTTP server for the JSON API.
type APIServer struct {
	address string
	userDB  UserDB
}

// NewAPIServer creates a new instance of APIServer with the given address.
func NewAPIServer(address string, userDB UserDB) *APIServer {
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
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

// Run starts the HTTP server and listens for incoming requests.
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/users", makeHTTPHandleFunc(s.handleUser))
	router.HandleFunc("/users/{id}", makeHTTPHandleFunc(s.handleGetUserByID))

	log.Println("JSON API Server is running on port:", s.address)

	http.ListenAndServe(s.address, router)
}

// handleUser handles requests related to user accounts.
func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
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
	userRequest := new(CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(userRequest); err != nil {
		return err
	}

	user := NewUser(
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
