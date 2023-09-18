package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

/*
   Listen for incoming http requests from the clients
*/

func main() {
	fmt.Println("Client is running on port 8081")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.ListenAndServe(":8081", nil)

	ping := true

	for {

		sendRequest()
		time.Sleep(5 * time.Second)

		if !ping {
			break
		}
	}
}

// send request to server to establish a http connection
// include a session id in the header request to identify the client
// send a message to the server
// receive a message from the server
func sendRequest() {

	serverAddr := os.Getenv("SERVER_ADDR")

	// create a new request
	req, err := http.NewRequest("Get", serverAddr, nil)
	if err != nil {
		fmt.Printf("ERROR: Failed to create new request %s \n", err)
		return
	}

	req.Header.Add("session_id", "123456789")

	// Create an Http client
	client := http.Client{
		Timeout: time.Second * 10,
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("ERROR: Failed to send request %s \n", err)
		return
	}
	defer resp.Body.Close()

	// Process the response
	fmt.Println("Response Status:", resp.Status)

}
