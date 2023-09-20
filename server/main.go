package main

import (
	"fmt"
	"main/server"
	"net/http"
)

func main() {
	fmt.Println("Server is running on port 8080")
	server.Banner()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	// switch to https
	http.ListenAndServe(":8080", nil)

}
