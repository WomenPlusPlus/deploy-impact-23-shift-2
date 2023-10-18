package api

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, status int, value interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

func WriteErrorResponse(w http.ResponseWriter, status int, value string) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(apiError{value}); err != nil {
		logrus.Errorf("could not send response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(status)
	}
}
