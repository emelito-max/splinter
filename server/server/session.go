package server

import "net/http"

type SessionHandler struct {
	// ...
	ID       int
	lifetime int
}

func NewSessionHandler() *SessionHandler {
	sessionHandler := &SessionHandler{}

	return sessionHandler
}

func (s *SessionHandler) SetID(sessionID int) {
	s.ID = sessionID
	return
}

func (s *SessionHandler) GetID() int {
	return s.ID
}

func (s *SessionHandler) ExecFunctionality(w http.ResponseWriter, r *http.Request) {
	// Handle request and execute some functionality
}
