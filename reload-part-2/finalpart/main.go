package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("usage: go run input.txt output.txt")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
	}

	content := string(input)

	text := strings.Split(content, "\n")

	var result []string

	for _, line := range text {
		if strings.TrimSpace(line) == "" {
			// result = append(result, "")
			continue
		}

		line = PuncArticle(line)
		result = append(result, line)
	}
	output := strings.Join(result, "\n")

	err = os.WriteFile(outputFile, []byte(output), 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
