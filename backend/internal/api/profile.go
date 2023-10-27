package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"shift/internal/entity"
	cauth "shift/pkg/auth"
	"strings"
)

func (s *APIServer) initProfileRoutes(router *mux.Router) {
	router.Path("/me").
		Handler(makeHTTPHandleFunc(s.handleGetUserProfile)).
		Methods(http.MethodGet)

	router.Path("/setup").
		Handler(makeHTTPHandleFunc(s.handleGetSetupInfo)).
		Methods(http.MethodGet)

	router.Path("/setup").
		Handler(makeHTTPHandleFunc(s.handleSetupProfile)).
		Methods(http.MethodPost)
}

func (s *APIServer) handleGetUserProfile(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("Get user profile handler")
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	claims, err := s.parseToken(r.Context(), token)
	if err != nil {
		WriteErrorResponse(w, http.StatusUnauthorized, fmt.Sprintf("Could not parse token: %s!", err))
		return nil
	}
	email := claims.CustomClaims.(*cauth.CustomClaims).Email

	user, err := s.userService.GetProfileByEmail(email)
	if err != nil {
		return WriteJSONResponse(w, http.StatusUnauthorized, fmt.Sprintf("unauthorized email %s", email))
	}
	return WriteJSONResponse(w, http.StatusOK, user)
}

func (s *APIServer) handleGetSetupInfo(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("Get setup info handler")
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	claims, err := s.parseToken(r.Context(), token)
	if err != nil {
		WriteErrorResponse(w, http.StatusUnauthorized, fmt.Sprintf("Could not parse token: %s!", err))
		return nil
	}
	email := claims.CustomClaims.(*cauth.CustomClaims).Email

	user, err := s.userService.GetProfileSetupByEmail(email)
	if err != nil {
		return InternalServerError{Message: err.Error()}
	}
	return WriteJSONResponse(w, http.StatusOK, user)
}

func (s *APIServer) handleSetupProfile(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("Setup profile handler")
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	claims, err := s.parseToken(r.Context(), token)
	if err != nil {
		WriteErrorResponse(w, http.StatusUnauthorized, fmt.Sprintf("Could not parse token: %s!", err))
		return nil
	}
	email := claims.CustomClaims.(*cauth.CustomClaims).Email

	req := new(entity.CreateUserRequest)
	if err := req.FromFormData(r); err != nil {
		return BadRequestError{Message: err.Error()}
	}
	req.Email = email

	res, err := s.userService.SetupProfile(req)
	if err != nil {
		return InternalServerError{Message: err.Error()}
	}

	return WriteJSONResponse(w, http.StatusOK, res)
}
