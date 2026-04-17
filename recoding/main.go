package main

import (
	"os"
	"strings"
	 "fmt"
)
func main(){
	if len(os.Args) != 3{
		fmt.Println("Usage: go run input.txt output.txt")
		return
	}
	
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	input, err := os.ReadFile(inputFile)
	if err != nil{
		fmt.Println("Error reading File")
	}

	text := string(input)

	lines := strings.Split(text, "\n")
	
	var result []string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			result = append(result, "")
			continue
		}

		line = AHB(line)
		// line = Punct(line)
		// line = ArtiQuotes(line)
		line = Fix(line)

		result = append(result, line)

		
	}
	output := strings.Join(result, "\n")

	err = os.WriteFile(outputFile, []byte(output), 0644)
	if err != nil{
		fmt.Println("Error Writing file")
	}
}
	