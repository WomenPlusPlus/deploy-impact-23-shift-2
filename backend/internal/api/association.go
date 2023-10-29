package api

import (
	"net/http"
	"shift/internal/entity"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (s *APIServer) initAssociationRoutes(router *mux.Router) {
	router = router.PathPrefix("/associations").Subrouter()

	router.Path("").
		Handler(makeHTTPHandleFunc(s.handleListAssociations)).
		Methods(http.MethodGet)

	router.Path("").
		Handler(makeHTTPHandleFunc(s.handleCreateAssociation)).
		Methods(http.MethodPost)

	router.Path("/{id}").
		Handler(makeHTTPHandleFunc(s.handleViewAssociation)).
		Methods(http.MethodGet)

	router.Path("/{id}").
		Handler(makeHTTPHandleFunc(s.handleDeleteAssociation)).
		Methods(http.MethodDelete)

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

func (s *APIServer) handleViewAssociation(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("View user handler running")

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "invalid value for parameter id")
		return nil
	}

	assoc, err := s.associationService.GetAssociationById(id)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return nil
	}

	return WriteJSONResponse(w, http.StatusOK, assoc)
}

func (s *APIServer) handleDeleteAssociation(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	if err := s.associationService.DeleteAssociation(id); err != nil {
		return WriteJSONResponse(w, http.StatusInternalServerError, apiError{Error: err.Error()})
	}

	return WriteJSONResponse(w, http.StatusOK, ResponseMessage{"association deleted successfully"})
}
