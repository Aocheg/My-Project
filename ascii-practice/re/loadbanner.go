package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file")
	}

	if len(data) == 0 {
		return nil, errors.New("invalid data")
	}

	splitfile := strings.Split(string(data), "\n")
	asciiMap := make(map[rune][]string)

	for i := 32; i < 127; i++ {
		char := rune(i)
		start := (i - 32) * 9

		if start+9 > len(splitfile) {
			return nil, errors.New("invalid banner format")
		}

		completelines := splitfile[start+1 : start+9]
		asciiMap[char] = completelines

	}
	return asciiMap, nil
}

func RenderLine(input string, banner map[rune][]string) []string {
	output := []string{}

	for row := 0; row < 8; row++ {
		var result strings.Builder
		for _, char := range input {
			new := banner[char][row]
			result.WriteString(new)
		}
		output = append(output, result.String())
	}
	return output
}
