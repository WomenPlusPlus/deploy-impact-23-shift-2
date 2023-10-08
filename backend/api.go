package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// WriteJSON writes a JSON response with the given status code and value.
func WriteJSON(w http.ResponseWriter, status int, value interface{}) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)
}

// apiFunc represents a function that handles API requests.
type apiFunc func(http.ResponseWriter, *http.Request) error

// apiError represents an API error response.
type apiError struct {
	Error string
}

// makeHTTPHandleFunc creates an HTTP request handler function for the provided apiFunc.
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

// APIServer represents an HTTP server for the JSON API.
type APIServer struct {
	address string
}

// NewAPIServer creates a new instance of APIServer with the given address.
func NewAPIServer(address string) *APIServer {
	return &APIServer{
		address: address,
	}
}

// Run starts the HTTP server and listens for incoming requests.
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/users", makeHTTPHandleFunc(s.handleUser))

	log.Println("JSON API Server is running on port:", s.address)

	http.ListenAndServe(s.address, router)
}

// handleUser handles requests related to user accounts.
func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetUser(w, r)
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
func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	account := NewUser("Test", "User", "Mrs.", "test_user@testusers.test", "Online", "https://placehold.co/400", "Candidate")
	return WriteJSON(w, http.StatusOK, account)
}

// handleCreateUser handles POST requests to create a new user account.
func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// handleDeleteUser handles DELETE requests to delete a user account.
func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// handleUpdateUser handles PUT requests to update a user account.
func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
