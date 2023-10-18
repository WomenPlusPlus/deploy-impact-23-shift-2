package api

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		logrus.Tracef("Running authentication middleware: token=%s", token)
		if token != "" {
			logrus.Tracef("Authenticated user: token=%s", token)
			next.ServeHTTP(w, r)
		} else {
			WriteErrorResponse(w, http.StatusForbidden, "Authentication failed!")
		}
	})
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string `json:"error"`
}

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

type PermissionError struct {
	Message string
}

func (e PermissionError) Error() string {
	return e.Message
}

type BadRequestError struct {
	Message string
}

func (e BadRequestError) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string
}

func (e InternalServerError) Error() string {
	return e.Message
}

func IsNotFoundError(err error) bool {
	_, isNotFound := err.(NotFoundError)
	return isNotFound
}

func IsPermissionError(err error) bool {
	_, isPermissionDenied := err.(PermissionError)
	return isPermissionDenied
}

func IsBadRequestError(err error) bool {
	_, isBadRequestError := err.(BadRequestError)
	return isBadRequestError
}

func IsInternalServerError(err error) bool {
	_, isInternalServerError := err.(InternalServerError)
	return isInternalServerError
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{}) // Use JSON format for structured logs
		err := f(w, r)
		if err != nil {
			switch {
			case IsNotFoundError(err):
				WriteErrorResponse(w, http.StatusNotFound, err.Error())
			case IsPermissionError(err):
				WriteErrorResponse(w, http.StatusForbidden, err.Error())
			case IsBadRequestError(err):
				WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			case IsInternalServerError(err):
				WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			default:
				// Log the internal error without exposing details to the client
				logger.Error(err)
				WriteErrorResponse(w, http.StatusInternalServerError, "Internal server error")
			}
		}
	}
}
