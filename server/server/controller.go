package server

import (
	json "encoding/json"
	"fmt"
	"io"
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
	sessionIds []string
	listening  bool

	requests  []string
	responses []string
}

// Struct that represents the json object that will be received from the client
type ClientRequest struct {
	Init      string `json:"init"`
	SessionId string `json:"sessionId"`
}

func NewController() *Controller {

	controller := &Controller{}

	return controller
}

func (c *Controller) Run() {

}

func (c *Controller) HandleSession(sessionId string) {

}

// HandleRequest Extract session information from request and route to the appropriate session handler
// Example: clientRequest.SessionId := r.Header.Get("SessionID")
// Example: sessionHandler := GetSessionHandler(clientRequest.SessionId)
// Perform any necessary operations on the session and return the response
// Example: sessionHandler.DoSomething(w, r)
func (c *Controller) HandleRequest(request *http.Request, response http.ResponseWriter) {

	go func() {

		var clientRequest ClientRequest

		// get session id from request
		err := json.NewDecoder(request.Body).Decode(&clientRequest)
		if err != nil {
			fmt.Println("ERROR: Could not decode request body")
			return
		}

		body, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Println("ERROR: Could not read request body")
			return
		}

		err = json.Unmarshal(body, &clientRequest)
		if err != nil {
			fmt.Println("ERROR: Could not unmarshal request body")
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("Session ID:", clientRequest.SessionId)

		// check if session exists
		// if session exists, update session
		// if session does not exist, create session

		// add session id to pool of session ids
		c.sessionIds = append(c.sessionIds, clientRequest.SessionId)

		response.Header().Set("Connection", "keep-alive")
		response.Header().Set("SessionID", clientRequest.SessionId)
		response.Header().Set("Content-Type", "application/json")
		fmt.Println("INFO: Session created successfully to client with session id:", clientRequest.SessionId)

		// send response to client
		_, err = fmt.Fprintf(response, "Hello World") // need explaination
		if err != nil {
			return
		}

		// create a new session
		session := NewSessionHandler()
		var connectionID, erro = strconv.Atoi(clientRequest.SessionId)
		if erro != nil {
			return
		}
		session.SetID(connectionID)

	}()
}

func (c *Controller) HandleResponse(response string) {

}

// CreateSession Create a goroutine for each client
func (c *Controller) CreateSession(sessionId string) {
	go func() {

	}()
}

func (c *Controller) DeleteSession() {
}

func (c *Controller) UpdateSession() {

}
