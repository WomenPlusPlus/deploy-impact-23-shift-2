package api

import (
	"net/http"
	"shift/internal/entity"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (s *APIServer) initAssociationrRoutes(router *mux.Router) {
	router = router.PathPrefix("/associations").Subrouter()

	router.Path("").
		Handler(makeHTTPHandleFunc(s.handleCreateAssociation)).
		Methods(http.MethodGet)
}

func (s *APIServer) handleCreateAssociation(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("Create association handler is running")

	req := new(entity.CreateAssociationRequest)
	if err := req.FromFormData(r); err != nil {
		return BadRequestError{Message: err.Error()}
	}

	res, err := s.associationService.CreateAssociation(req)
	if err != nil {
		return InternalServerError{Message: err.Error()}
	}

	return WriteJSONResponse(w, http.StatusOK, res)
}

func (s *APIServer) handleListAssociations(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("List associations handler is running")

	res, err := s.associationService.ListAssociations()
	if err != nil {
		return InternalServerError{Message: err.Error()}
	}

	return WriteJSONResponse(w, http.StatusOK, res)
}
