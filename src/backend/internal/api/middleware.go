package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"net/http"
	"shift/internal/entity"
	cauth "shift/pkg/auth"
	"strings"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allMethods := []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodHead,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(allMethods, ","))
		w.Header().Set("Access-Control-Allow-Headers", "*")

		next.ServeHTTP(w, r)
	})
}

func (s *APIServer) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		ctx := r.Context()

		logrus.Tracef("Running authentication middleware: token=%s", token)
		claims, err := s.parseToken(ctx, token)
		if err != nil {
			WriteErrorResponse(w, http.StatusUnauthorized, fmt.Sprintf("Authentication failed: %s!", err))
			return
		}
		logrus.Tracef("Authenticated user: token=%s, claims=%#v | %#v", token, claims, claims.CustomClaims)

		email := claims.CustomClaims.(*cauth.CustomClaims).Email

		user, err := s.userService.GetUserRecordByEmail(email)
		if err != nil {
			logrus.Tracef("Authentication failed: invalid user: %v", err)
			WriteErrorResponse(w, http.StatusUnauthorized, "Authentication failed: invalid user!")
			return
		}
		logrus.Tracef("Authenticated valid user: user=%v", user)

		ctx = context.WithValue(ctx, entity.ContextKeyKind, user.Kind)
		ctx = context.WithValue(ctx, entity.ContextKeyRole, user.Role)
		ctx = context.WithValue(ctx, entity.ContextKeyUserId, user.ID)
		ctx = context.WithValue(ctx, entity.ContextKeyEmail, user.Email)
		ctx = context.WithValue(ctx, entity.ContextKeyToken, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *APIServer) parseToken(ctx context.Context, token string) (*validator.ValidatedClaims, error) {
	tokenClaims, err := s.jwtValidator.ValidateToken(ctx, token)
	if err != nil {
		logrus.Tracef("Authentication failed: %v", err)
		return nil, errors.New("invalid token")
	}
	claims, ok := tokenClaims.(*validator.ValidatedClaims)
	if !ok {
		logrus.Tracef("Authentication failed: invalid claims: %v", err)
		return nil, errors.New("invalid claims")
	}
	return claims, nil
}

func AuthorizationMiddleware(key entity.ContextKey, values ...string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return AuthorizationHandler(next, key, values...)
	}
}

func AuthorizationHandler(next http.Handler, key entity.ContextKey, values ...string) http.Handler {
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

func AuthorizationCase(next http.Handler, key entity.ContextKey, values ...string) func(http.ResponseWriter, *http.Request) bool {
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

func AuthorizationDefault(next http.Handler) func(http.ResponseWriter, *http.Request) bool {
	return func(w http.ResponseWriter, r *http.Request) bool {
		logrus.Tracef("Authorized user: default")
		next.ServeHTTP(w, r)
		return true
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

type ResponseMessage struct {
	Message string
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
