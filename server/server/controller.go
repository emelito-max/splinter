package server

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
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Run() {
}

func (c *Controller) HandleSession(sessionID string) {

}

func (c *Controller) HandleRequest(request string) {

}

func (c *Controller) HandleResponse(response string) {

}

// CreateSession Create a goroutine for each client
func (c *Controller) CreateSession() {
}

func (c *Controller) DeleteSession() {
}

func (c *Controller) UpdateSession() {

}
