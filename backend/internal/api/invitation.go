package api

// func (s *APIServer) initInvitaionRoutes(router *mux.Router) {
// 	router = router.PathPrefix("/invitations").Subrouter()

// 	router.Path("").
// 		Handler(makeHTTPHandleFunc(s.handleCreateInvitation)).
// 		Methods(http.MethodPost)

// 	router.Use(AuthenticationMiddleware)
// 	router.Use(AuthorizationMiddleware(ContextKeyKind, entity.UserKindAdmin))
// }

// func (s *APIServer) handleCreateInvitation(w http.ResponseWriter, r *http.Request) error {
// 	invRequest := new(entity.CreateInvitationRequest)
// 	if err := json.NewDecoder(r.Body).Decode(invRequest); err != nil {
// 		return err
// 	}

// 	inv := entity.NewInvitation(
// 		invRequest.Kind,
// 		invRequest.Email,
// 		invRequest.Subject,
// 		invRequest.Message,
// 	)

// 	if _, err := s.invi.CreateInvitation(inv); err != nil {
// 		return WriteJSONResponse(w, http.StatusNotFound, apiError{Error: err.Error()})
// 	}

// 	return WriteJSONResponse(w, http.StatusOK, inv)

// }
