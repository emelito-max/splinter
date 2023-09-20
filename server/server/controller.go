package server

import (
	"fmt"
	"net/http"
	"strconv"
)

/*
   Controller to handle the sessions to the clients.
   It will be responsible for the following:
       - create goroutine for each client
	   - Handle the sessions to the clients
         (create, delete, update, etc.)
	   - Handle the requests from the clients
	   - Handle the responses to the clients
*/

type Controller struct {
	sessionIDs []string
	listening  bool

	requests  []string
	responses []string
}

func NewController() *Controller {

	controller := &Controller{}

	return controller
}

func (c *Controller) Run() {
}

func (c *Controller) HandleSession(sessionID string) {

}

func (c *Controller) HandleRequest(request *http.Request, response http.ResponseWriter) {
	// Extract session information from request and route to the appropriate session handler
	// Example: sessionID := r.Header.Get("SessionID")
	// Example: sessionHandler := GetSessionHandler(sessionID)

	// Perform any necessary operations on the session and return the response
	// Example: sessionHandler.DoSomething(w, r)

	go func() {
		sessionID := request.Header.Get("session_id")
		c.HandleSession(sessionID)

		// check if session exists
		// if session exists, update session
		// if session does not exist, create session

		fmt.Println("Session ID:", sessionID)

		// add session id to pool of session ids
		c.sessionIDs = append(c.sessionIDs, sessionID)

		response.Header().Set("Connection", "keep-alive")
		fmt.Println("INFO: Session created successfully to client with session id:", sessionID)

		// send response to client
		_, err := fmt.Fprintf(response, "Hello World") // need explaination
		if err != nil {
			return
		}

		// create a new session
		session := NewSessionHandler()
		var connectionID, erro = strconv.Atoi(sessionID)
		if erro != nil {
			return
		}
		session.SetID(connectionID)

	}()
}

func (c *Controller) HandleResponse(response string) {

}

// CreateSession Create a goroutine for each client
func (c *Controller) CreateSession(sessionID string) {
	go func() {

	}()
}

func (c *Controller) DeleteSession() {
}

func (c *Controller) UpdateSession() {

}
