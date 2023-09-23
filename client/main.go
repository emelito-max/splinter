package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var sessionID = generateId()

func main() {
	fmt.Println("Client is running on port 8081")

	ping := true

	for {

		sendRequest()
		fmt.Println("INFO: Sending GET request to server")
		initConnection()
		fmt.Println("INFO: Sending POST request to server")
		time.Sleep(5 * time.Second)

		if !ping {
			break
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.ListenAndServe(":8081", nil)

}

// send request to server to establish a http connection
// include a session id in the header request to identify the client
// send a message to the server
// receive a message from the server
func sendRequest() {

	//serverAddr := os.Getenv("SERVER_ADDR")
	serverAddr := "http://192.168.0.2:8080"

	// create a new request
	req, err := http.NewRequest("Get", serverAddr, nil)
	if err != nil {
		fmt.Printf("ERROR: Failed to create new request %s \n", err)
		return
	}

	req.Header.Add("session_id", string(rune(sessionID)))

	// Create an Http client
	client := http.Client{
		Timeout: time.Second * 1,
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("ERROR: Failed to send request %s \n", err)
		return
	}

	fmt.Println("INFO: Request sent successfully")
	defer resp.Body.Close()

	// Process the response
	fmt.Println("Response Status:", resp.Status)

}

// send post request to server to fetch command
// include session id in the header request to identify the client
// send a message in the body of the request in json to indicate a request for a command
// receive a message from the server in json format
func initConnection() {

	//serverAddr := os.Getenv("SERVER_ADDR")
	serverAddr := "http://192.168.0.2:8080"

	initRequest := map[string]string{
		"init": "true",
	}
	jsonData, err := json.Marshal(initRequest)
	if err != nil {
		fmt.Printf("ERROR: Failed to marshal json %s \n", err)
		panic(err)
		return
	}

	// create a new request
	req, err := http.NewRequest("Post", serverAddr, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("ERROR: Failed to create new request %s \n", err)
		panic(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("ERROR: Failed to send request %s \n", err)
		panic(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ERROR: Failed to read response body %s \n", err)
		panic(err)
		return
	}
	fmt.Println("Response Body:", string(body))

}

func generateId() string {

	// generate a random session id
	rand.Seed(time.Now().UnixNano())
	sessionID := rand.Intn(1000000000)
	sessionIDtest := string(rune(sessionID))
	fmt.Println("Session ID:", sessionIDtest)

	return sessionIDtest
}
