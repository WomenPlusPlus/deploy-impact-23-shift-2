package api

import (
	"net/http"
	"shift/internal/entity/invitation"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (s *APIServer) initInvitaionRoutes(router *mux.Router) {
	router = router.PathPrefix("/invitations").Subrouter()

	router.Path("/create").
		Handler(makeHTTPHandleFunc(s.handleCreateInvitation)).
		Methods(http.MethodPost)

	router.Path("").
		Handler(makeHTTPHandleFunc(s.handleListInvitations)).
		Methods(http.MethodPost)

	// router.Use(AuthenticationMiddleware)
	// router.Use(AuthorizationMiddleware(ContextKeyKind, entity.UserKindAdmin))
}

func (s *APIServer) handleCreateInvitation(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("Create invitation handler is running")

	req := new(invitation.CreateInvitationRequest)
	if err := req.FromFormData(r); err != nil {
		return BadRequestError{Message: err.Error()}
	}

	res, err := s.invitationService.CreateInvitation(req)
	if err != nil {
		return InternalServerError{Message: err.Error()}
	}

	return WriteJSONResponse(w, http.StatusOK, res)
}

func (s *APIServer) handleListInvitations(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("List invitations handler running")
	// res, err := s.invitationService.ListInvitations()
	return nil
}
