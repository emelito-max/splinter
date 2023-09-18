package server

import "fmt"

// Print the banner of the program
func Banner() {
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
	lines := SplitLines(art)
	for _, line := range lines {
		fmt.Println(line)
	}

}

func SplitLines(art string) []string {
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
