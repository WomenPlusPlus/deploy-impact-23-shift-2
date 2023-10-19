package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"net/http"
)

type ContextKey string

var (
	ContextKeyKind  ContextKey = "X-Kind"
	ContextKeyRole  ContextKey = "X-Role"
	ContextKeyToken ContextKey = "X-Token"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		ctx := context.WithValue(r.Context(), ContextKeyKind, r.Header.Get(string(ContextKeyKind)))
		ctx = context.WithValue(ctx, ContextKeyRole, r.Header.Get(string(ContextKeyRole)))
		ctx = context.WithValue(ctx, ContextKeyToken, token)

		logrus.Tracef("Running authentication middleware: token=%s", token)
		if token != "" {
			logrus.Tracef("Authenticated user: token=%s", token)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			WriteErrorResponse(w, http.StatusForbidden, "Authentication failed!")
		}
	})
}

func AuthorizationMiddleware(key ContextKey, values ...string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return AuthorizationHandler(next, key, values...)
	}
}

func AuthorizationHandler(next http.Handler, key ContextKey, values ...string) http.Handler {
	handler := AuthorizationCase(next, key, values...)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !handler(w, r) {
			WriteErrorResponse(w, http.StatusForbidden, "Authorization failed: not enough permissions!")
		}
	})
}

func AuthorizationSwitch(handlers ...func(http.ResponseWriter, *http.Request) bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range handlers {
			if handler(w, r) {
				return
			}
		}
		WriteErrorResponse(w, http.StatusForbidden, "Authorization failed: not enough permissions!")
	})
}

func AuthorizationCase(next http.Handler, key ContextKey, values ...string) func(http.ResponseWriter, *http.Request) bool {
	return func(w http.ResponseWriter, r *http.Request) bool {
		value, ok := r.Context().Value(key).(string)
		if !ok {
			logrus.Tracef("Invalid authorization key: key=%s, value=%v", key, value)
			return false
		}
		logrus.Tracef("Running authorization middleware: key=%s, value=%s", key, value)
		if slices.Contains(values, value) {
			logrus.Tracef("Authorized user: key=%s, value=%s", key, value)
			next.ServeHTTP(w, r)
			return true
		}
		logrus.Tracef("Not authorized user: key=%s, value=%s", key, value)
		return false
	}
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