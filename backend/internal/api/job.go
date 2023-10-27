package api

import (
	"net/http"
	"shift/internal/entity"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (s *APIServer) initJobRoutes(router *mux.Router) {
	router = router.PathPrefix("/jobs").Subrouter()

	router.Path("").
		Handler(makeHTTPHandleFunc(s.handleCreateJob)).
		Methods(http.MethodPost)

	router.Path("").
		Handler(makeHTTPHandleFunc(s.handleListJobs)).
		Methods(http.MethodGet)

	router.Path("/{id}").
		Handler(makeHTTPHandleFunc(s.handleViewJob)).
		Methods(http.MethodGet)

	router.Path("/{id}").
		Handler(makeHTTPHandleFunc(s.handleDeleteJob)).
		Methods(http.MethodDelete)

}

func (s *APIServer) handleCreateJob(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("Create job handler running")

	req := new(entity.CreateJobRequest)
	if err := req.FromFormData(r); err != nil {
		return BadRequestError{Message: err.Error()}
	}

	res, err := s.jobService.CreateJob(req)
	if err != nil {
		return InternalServerError{Message: err.Error()}
	}

	return WriteJSONResponse(w, http.StatusOK, res)
}

func (s *APIServer) handleListJobs(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("List jobs handler running")

	res, err := s.jobService.ListJobs()
	if err != nil {
		return InternalServerError{Message: err.Error()}
	}

	return WriteJSONResponse(w, http.StatusOK, res)
}

func (s *APIServer) handleViewJob(w http.ResponseWriter, r *http.Request) error {
	logrus.Debugln("View job handler running")

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "invalid value for parameter id")
		return nil
	}

	job, err := s.jobService.GetJobById(id)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return nil
	}

	return WriteJSONResponse(w, http.StatusOK, job)
}

func (s *APIServer) handleDeleteJob(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	if _, err := s.jobDB.GetJobById(id); err != nil {
		return WriteJSONResponse(w, http.StatusNotFound, apiError{Error: err.Error()})
	}

	return WriteJSONResponse(w, http.StatusOK, "job deleted successfully")
}
