package main

import (
	"fmt"
	"net/http"
	"splinter/server"
)

func main() {
	fmt.Println("Server is running on port 8080")
	server.Banner()

	// create a new controller
	controller := server.NewController()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
		controller.HandleRequest(r, w)
	})
	// switch to https
	http.ListenAndServe(":8080", nil)

}
