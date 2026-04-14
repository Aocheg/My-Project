package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run. input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	text := string(content)

	words := SplitWords(text)

	words = Process(words)

	result := ""
	for i := 0; i < len(words); i++ {
		if i > 0 {
			if !isPunctuationWord(words[i]) {
				result += " "
			}
		}
		result += words[i]

		if i < len(words)-1 {

		}
	}
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error go writing output file:", err)
		return
	}
}
