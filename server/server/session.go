package server

import "net/http"

type SessionHandler struct {
	// ...

	ID string
}

func NewSessionHandler() *SessionHandler {
	sessionHandler := &SessionHandler{}

	return sessionHandler
}

func (s *SessionHandler) GetID() string {
	return s.ID
}

func (s *SessionHandler) ExecFunctionality(w http.ResponseWriter, r *http.Request) {
	// Handle request and execute some functionality
}
