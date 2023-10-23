package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"shift/internal/entity"
)

// TODO: delete later
func (s *APIServer) initAuthorizationRoutes(router *mux.Router) {
	router = router.PathPrefix("/authorization").Subrouter()

	router.Path("").
		Handler(
			AuthorizationSwitch(
				AuthorizationCase(makeHTTPHandleFunc(s.handleAdminAuthorization), ContextKeyKind, entity.UserKindAdmin),
				AuthorizationCase(
					AuthorizationHandler(
						makeHTTPHandleFunc(s.handleEntityAuthorization),
						ContextKeyRole,
						entity.UserRoleAdmin,
					),
					ContextKeyKind,
					entity.UserKindAssociation,
					entity.UserKindCompany,
				),
				AuthorizationCase(makeHTTPHandleFunc(s.handleCandidateAuthorization), ContextKeyKind, entity.UserKindCandidate),
			),
		).
		Methods(http.MethodPost)

	router.Use(s.AuthenticationMiddleware)
}

// TODO: delete later
func (s *APIServer) handleAdminAuthorization(w http.ResponseWriter, _ *http.Request) error {
	return WriteJSONResponse(w, http.StatusOK, struct{ Message string }{"Admin"})
}

// TODO: delete later
func (s *APIServer) handleCandidateAuthorization(w http.ResponseWriter, _ *http.Request) error {
	return WriteJSONResponse(w, http.StatusOK, struct{ Message string }{"Candidate"})
}

// TODO: delete later
func (s *APIServer) handleEntityAuthorization(w http.ResponseWriter, _ *http.Request) error {
	return WriteJSONResponse(w, http.StatusOK, struct{ Message string }{"Other"})
}
