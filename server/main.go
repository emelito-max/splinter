package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is running on port 8080")
	banner()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.ListenAndServe(":8080", nil)

}

// Print the banner of the program
func banner() {
	art := `            
			   _         _
			  | (_)     | |           
	 ___ _ __ | |_ _ __ | |_ ___ _ __ 
	/ __| '_ \| | | '_ \| __/ _ \ '__|
	\__ \ |_) | | | | | | ||  __/ |   
	|___/ .__/|_|_|_| |_|\__\___|_|   
		| |                           
		|_|                           
`
	lines := splitLines(art)
	for _, line := range lines {
		fmt.Println(line)
	}

}

func splitLines(art string) []string {
	var lines []string
	currentLine := ""
	for _, char := range art {
		if char == '\n' {
			lines = append(lines, currentLine)
			currentLine = ""
		} else {
			currentLine += string(char)
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return lines
}
