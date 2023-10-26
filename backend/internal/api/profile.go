package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"shift/internal/entity"
)

func (s *APIServer) initProfileRoutes(router *mux.Router) {
	router.Path("/me").
		Handler(s.AuthenticationMiddleware(makeHTTPHandleFunc(s.handleGetUserProfile))).
		Methods(http.MethodGet)
}

func (s *APIServer) handleGetUserProfile(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("Get user profile handler")
	email, ok := r.Context().Value(entity.ContextKeyEmail).(string)
	if !ok {
		WriteErrorResponse(w, http.StatusUnauthorized, "Email not provided!")
		return nil
	}

	user, err := s.userService.GetProfileByEmail(email)
	if err != nil {
		WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
		return nil
	}

	return WriteJSONResponse(w, http.StatusOK, user)
}
