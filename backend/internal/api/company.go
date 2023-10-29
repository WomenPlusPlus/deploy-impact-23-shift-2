package api

import (
	"net/http"
	"shift/internal/entity"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (s *APIServer) initCompanyRoutes(router *mux.Router) {
	router = router.PathPrefix("/companies").Subrouter()

	router.Path("").
		Handler(makeHTTPHandleFunc(s.handleCreateCompany)).
		Methods(http.MethodPost)

	router.Path("").
		Handler(makeHTTPHandleFunc(s.handleListCompanies)).
		Methods(http.MethodGet)

	router.Path("/{id}").
		Handler(makeHTTPHandleFunc(s.handleViewCompany)).
		Methods(http.MethodGet)

	router.Path("/{id}").
		Handler(makeHTTPHandleFunc(s.handleDeleteCompany)).
		Methods(http.MethodDelete)

	router.Path("/{id}/jobs").
		Handler(makeHTTPHandleFunc(s.handleGetCompanyJobs)).
		Methods(http.MethodGet)

	router.Use(s.AuthenticationMiddleware)
	router.Use(AuthorizationMiddleware(entity.ContextKeyKind, entity.UserKindAdmin))

}

func (s *APIServer) handleCreateCompany(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("Create company handler running")

	req := new(entity.CreateCompanyRequest)
	if err := req.FromFormData(r); err != nil {
		return BadRequestError{Message: err.Error()}
	}

	res, err := s.companyService.CreateCompany(req)
	if err != nil {
		return InternalServerError{Message: err.Error()}
	}

	return WriteJSONResponse(w, http.StatusOK, res)
}

func (s *APIServer) handleListCompanies(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("List companies handler running")

	res, err := s.companyService.ListCompanies()
	if err != nil {
		return InternalServerError{Message: err.Error()}
	}

	return WriteJSONResponse(w, http.StatusOK, res)
}

func (s *APIServer) handleViewCompany(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("View company handler running")

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "invalid value for parameter id")
		return nil
	}

	company, err := s.companyService.GetCompanyById(id)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return nil
	}

	return WriteJSONResponse(w, http.StatusOK, company)
}

func (s *APIServer) handleGetCompanyJobs(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("Get company jobs handler running")

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "invalid value for parameter id")
		return nil
	}

	company, err := s.jobService.GetJobsByCompanyId(id)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return nil
	}

	return WriteJSONResponse(w, http.StatusOK, company)
}

func (s *APIServer) handleDeleteCompany(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	if err := s.companyService.DeleteCompanyById(id); err != nil {
		return WriteJSONResponse(w, http.StatusInternalServerError, apiError{Error: err.Error()})
	}

	return WriteJSONResponse(w, http.StatusOK, ResponseMessage{"company deleted successfully"})
}
