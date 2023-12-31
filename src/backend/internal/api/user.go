package api

import (
	"net/http"
	"shift/internal/entity"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (s *APIServer) initUserRoutes(router *mux.Router) {
	router = router.PathPrefix("/users").Subrouter()

	router.Path("").
		Handler(AuthorizationHandler(makeHTTPHandleFunc(s.handleCreateUser), entity.ContextKeyKind, entity.UserKindAdmin)).
		Methods(http.MethodPost)

	router.Path("").
		Handler(
			AuthorizationSwitch(
				AuthorizationCase(makeHTTPHandleFunc(s.handleListUsersForAdmins), entity.ContextKeyKind, entity.UserKindAdmin),
				AuthorizationDefault(makeHTTPHandleFunc(s.handleListUsersForNonAdmins)),
			),
		).
		Methods(http.MethodGet)

	router.Path("/{id}").
		Handler(makeHTTPHandleFunc(s.handleViewUser)).
		Methods(http.MethodGet)

	router.Path("/{id}").
		Handler(makeHTTPHandleFunc(s.handleEditUser)).
		Methods(http.MethodPut)

	router.Path("/{id}").
		Handler(AuthorizationHandler(makeHTTPHandleFunc(s.handleDeleteUser), entity.ContextKeyKind, entity.UserKindAdmin)).
		Methods(http.MethodDelete)

	router.Use(s.AuthenticationMiddleware)
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

func (s *APIServer) handleListUsersForAdmins(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("List users handler running")

	res, err := s.userService.ListUsers(true)
	if err != nil {
		return InternalServerError{Message: err.Error()}
	}

	return WriteJSONResponse(w, http.StatusOK, res)
}

func (s *APIServer) handleListUsersForNonAdmins(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("List users handler running")

	res, err := s.userService.ListUsers(false)
	if err != nil {
		return InternalServerError{Message: err.Error()}
	}

	return WriteJSONResponse(w, http.StatusOK, res)
}

func (s *APIServer) handleViewUser(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("View user handler running")

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "invalid value for parameter id")
		return nil
	}

	user, err := s.userService.GetUserById(id)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return nil
	}

	return WriteJSONResponse(w, http.StatusOK, user)
}

func (s *APIServer) handleEditUser(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("Edit user handler running")

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "invalid value for parameter id")
		return nil
	}

	if canAccess, err := UserResourcePermission(r.Context(), id); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return nil
	} else if !canAccess {
		WriteErrorResponse(w, http.StatusForbidden, "insufficient permissions")
		return nil
	}

	req := new(entity.EditUserRequest)
	if err := req.FromFormData(id, r); err != nil {
		return BadRequestError{Message: err.Error()}
	}

	res, err := s.userService.EditUser(id, req)
	if err != nil {
		return InternalServerError{Message: err.Error()}
	}

	return WriteJSONResponse(w, http.StatusOK, res)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	if err := s.userService.DeleteUserById(id); err != nil {
		return WriteJSONResponse(w, http.StatusInternalServerError, apiError{Error: err.Error()})
	}

	return WriteJSONResponse(w, http.StatusOK, ResponseMessage{"user deleted successfully"})
}
