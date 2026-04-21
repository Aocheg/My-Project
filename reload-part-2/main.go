package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("usage: go run . input.txt output.txt")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("erro reading File")
	}
	content := string(input)

	text := strings.Split(content, "\n")

	var result []string

	for _, lines := range text {
		if strings.TrimSpace(lines) == "" {
			result = append(result, "")
			continue
		}
		lines = PuncArticle(lines)
		result = append(result, lines)
	}
	output := strings.Join(result, "\n")

	err = os.WriteFile(outputFile, []byte(output), 0644)
}
