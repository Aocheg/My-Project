package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("error")
	}
	inputFile := os.Args[1]
	outputFIle := os.Args[2]

	val, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
	}

	text := string(val)

	lines := strings.Split(text, "\n")

	var result []string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			result = append(result, "")
		}
		line = Ahb(line)
		line = Pa(line)
		result = append(result, line)
	}
	output := strings.Join(result, "\n")

	err = os.WriteFile(outputFIle, []byte(output), 0644)
}
