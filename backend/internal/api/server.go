package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shift/internal/db"
	"shift/internal/entity"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer represents an HTTP server for the JSON API.
type APIServer struct {
	address string
	//userDB  entity.UserDB
	storage db.Storage
}

// NewAPIServer creates a new instance of APIServer with the given address.
func NewAPIServer(address string, storage db.Storage) *APIServer {
	return &APIServer{
		address: address,
		storage: storage,
		//userDB:  userDB,
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
	fmt.Println("In Run")
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))

	router.HandleFunc("/users", makeHTTPHandleFunc(s.handleUsers))

	router.HandleFunc("/companies", makeHTTPHandleFunc(s.handleCompanies))
	router.HandleFunc("/companies/{id}", makeHTTPHandleFunc(s.handleGetCompanyByID))

	router.HandleFunc("/joblistings", makeHTTPHandleFunc(s.handleJobListings))
	router.HandleFunc("/joblistings/{id}", makeHTTPHandleFunc(s.handleGetJobListingByID))

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
	fmt.Println("In makeHTTPHandleFunc")

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

// handleUsers handles requests related to user accounts.
func (s *APIServer) handleUsers(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("In handleUsers")

	if r.Method == "GET" {
		fmt.Println("In handleUsers GET")

		return s.handleGetUsers(w, r)
	}
	if r.Method == "POST" {
		fmt.Println("In handleUsers POST")

		return s.handleCreateUser(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

// handleGetUser handles GET requests for user account information.
func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("In handleGetUsers")

	users, err := s.storage.GetUsers()
	fmt.Println(users)
	if err != nil {
		return err
	}
	return WriteJSONResponse(w, http.StatusOK, users)
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("In handleCreateUser")
	createUserRequest := new(entity.CreateUserRequest)
	fmt.Println("In handleCreateUser 2 ")
	if err := json.NewDecoder(r.Body).Decode(createUserRequest); err != nil {
		fmt.Println("In handleCreateUser 3 ")
		return err
	}
	fmt.Println("In handleCreateUser 4 ")

	user := entity.NewUser(
		createUserRequest.FirstName,
		createUserRequest.LastName,
		createUserRequest.PreferredName,
		createUserRequest.Email,
		createUserRequest.State,
		createUserRequest.ImageUrl,
		createUserRequest.Role,
	)

	if err := s.storage.CreateUser(user); err != nil {
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

// WriteJSONResponse writes a JSON response with the given status code and value.
func WriteJSONResponse(w http.ResponseWriter, status int, value interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

// func (s *APIServer) handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
// 	idStr := mux.Vars(r)["id"]

// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		return fmt.Errorf("invalid id given %s", idStr)
// 	}

// 	user, err := s.userDB.GetUserByID(id)

// 	if err != nil {
// 		return err
// 	}

// 	return WriteJSONResponse(w, http.StatusOK, user)
// }

// handleDeleteUser handles DELETE requests to delete a user account.
// func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }

// handleUpdateUser handles PUT requests to update a user account.
// func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }

// handleCompanies handles requests related to comapny.
func (s *APIServer) handleCompanies(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("In handleCompanies")

	fmt.Println("in handleCompanies ")
	if r.Method == "GET" {
		fmt.Println("in handleCompaniesGET")
		return s.handleGetCompanies(w, r)
	}

	if r.Method == "POST" {
		fmt.Println("in handleCompanies POST")
		return s.handleCreateCompany(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

// handleGetCompany handles GET requests for companies information.
func (s *APIServer) handleGetCompanies(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("In handleGetCompanies")
	users, err := s.storage.GetCompanies()

	if err != nil {
		return err
	}

	return WriteJSONResponse(w, http.StatusOK, users)
}

func (s *APIServer) handleCreateCompany(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("in handleCreateCompany ")

	// companyRequest := new(entity.CreateCompanyRequest)
	// fmt.Println("in handleCreateCompany 2")

	// if err := json.NewDecoder(r.Body).Decode(companyRequest); err != nil {
	// 	fmt.Println("in handleCreateCompany 3")

	// 	return err
	// }
	// fmt.Println("in handleCreateCompany 4")

	// company := entity.NewCompany(
	// 	companyRequest.CompanyName,
	// 	companyRequest.Email,
	// )
	company := entity.NewCompany(
		"CompanyName",
		"Email",
	)

	if err := s.storage.CreateCompany(company); err != nil {
		return err
	}

	return WriteJSONResponse(w, http.StatusOK, company)
}

// func (s *APIServer) handleCreateCompany(w http.ResponseWriter, r *http.Request) error {
// 	fmt.Println("in handleCreateCompany ")
// 	companyRequest := new(entity.CreateCompanyRequest)
// 	fmt.Println("in handleCreateCompany 1 ")
// 	b, err := io.ReadAll(r.Body)
// 	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
// 	if err != nil {
// 		fmt.Println("in handleCreateCompany 3")
// 		log.Fatalln(err)
// 	}
// 	fmt.Println("in handleCreateCompany 4")

// 	fmt.Println(string(b))
// 	// if err := json.NewDecoder(r.Body).Decode(userRequest); err != nil {
// 	// 	fmt.Println("in handleCreateUser 2")
// 	// 	return err
// 	// }
// 	fmt.Println("in handleCreateCompany 5")

// 	company := entity.NewCompany(
// 		companyRequest.CompanyName,
// 		companyRequest.Email,
// 		//...
// 	)
// 	if err := s.storage.CreateCompany(company); err != nil {
// 		return err
// 	}

// 	fmt.Println("company created")

// 	return WriteJSONResponse(w, http.StatusOK, company)
// }

func (s *APIServer) handleGetCompanyByID(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("handleGetCompanyByID")
	if r.Method == "GET" {
		fmt.Println("In handleGetCompanyByID get")
		id, err := getID(r)
		fmt.Println(id)
		if err != nil {
			return err
		}

		account, err := s.storage.GetCompanyByID(id)
		if err != nil {
			return err
		}

		return WriteJSONResponse(w, http.StatusOK, account)
	}

	if r.Method == "DELETE" {
		fmt.Println("In handleGetCompanyByID DELETE")
		return s.handleDeleteCompany(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}
	return id, nil
}

func (s *APIServer) handleDeleteCompany(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("In handleDeleteCompany")
	id, err := getID(r)
	if err != nil {
		return err
	}

	if err := s.storage.DeleteCompany(id); err != nil {
		return err
	}

	return WriteJSONResponse(w, http.StatusOK, map[string]int{"deleted": id})
}

//JobListings

func (s *APIServer) handleJobListings(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("In handleJobListings")

	if r.Method == "GET" {
		fmt.Println("In handleJobListings GET")

		return s.handleGetJobListings(w, r)
	}
	if r.Method == "POST" {
		fmt.Println("In handleJobListings POST")

		return s.handleCreateJobListing(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetJobListings(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("In handleGetJobListings")
	users, err := s.storage.GetJobListings()

	if err != nil {
		return err
	}

	return WriteJSONResponse(w, http.StatusOK, users)
}

func (s *APIServer) handleCreateJobListing(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("in handleCreateJobListing ")

	// companyRequest := new(entity.CreateCompanyRequest)
	// fmt.Println("in handleCreateCompany 2")

	// if err := json.NewDecoder(r.Body).Decode(companyRequest); err != nil {
	// 	fmt.Println("in handleCreateCompany 3")

	// 	return err
	// }
	// fmt.Println("in handleCreateCompany 4")

	// company := entity.NewCompany(
	// 	companyRequest.CompanyName,
	// 	companyRequest.Email,
	// )
	jl := entity.NewJobListing(
		"Job",
		"description",
	)

	if err := s.storage.CreateJobListing(jl); err != nil {
		return err
	}

	return WriteJSONResponse(w, http.StatusOK, jl)
}

func (s *APIServer) handleGetJobListingByID(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("handleGetJobListingByID")
	if r.Method == "GET" {
		fmt.Println("In handleGetJobListingByID get")
		id, err := getID(r)
		fmt.Println(id)
		if err != nil {
			return err
		}

		jl, err := s.storage.GetJobListingByID(id)
		if err != nil {
			return err
		}

		return WriteJSONResponse(w, http.StatusOK, jl)
	}

	if r.Method == "DELETE" {
		fmt.Println("In handleGetJobListingByID DELETE")
		return s.handleDeleteJobListing(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleDeleteJobListing(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("In handleDeleteJobListing")
	id, err := getID(r)
	if err != nil {
		return err
	}

	if err := s.storage.DeleteJoblisting(id); err != nil {
		return err
	}

	return WriteJSONResponse(w, http.StatusOK, map[string]int{"deleted": id})
}
