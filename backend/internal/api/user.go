package api

import (
	"net/http"
	"shift/internal/entity"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (s *APIServer) initUserRoutes(router *mux.Router) {
	router = router.PathPrefix("/users").Subrouter()

	router.Path("").
		Handler(makeHTTPHandleFunc(s.handleCreateUser)).
		Methods(http.MethodPost)

	router.Path("").
		Handler(makeHTTPHandleFunc(s.handleListUsers)).
		Methods(http.MethodGet)

	router.Use(AuthenticationMiddleware)
	router.Use(AuthorizationMiddleware(ContextKeyKind, entity.UserKindAdmin))
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

func (s *APIServer) handleListUsers(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("List users handler running")

	res, err := s.userService.ListUsers()
	if err != nil {
		return InternalServerError{Message: err.Error()}
	}

	return WriteJSONResponse(w, http.StatusOK, res)
}
