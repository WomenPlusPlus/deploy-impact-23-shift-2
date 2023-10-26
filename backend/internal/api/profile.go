package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	cauth "shift/pkg/auth"
	"strings"
)

func (s *APIServer) initProfileRoutes(router *mux.Router) {
	router.Path("/me").
		Handler(makeHTTPHandleFunc(s.handleGetUserProfile)).
		Methods(http.MethodGet)
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
