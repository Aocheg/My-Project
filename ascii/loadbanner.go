package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error reading file")
	}

	if len(file) == 0 {
		return nil, errors.New("empty file")
	}

	splitfile := strings.Split(string(file), "\n")
	asciiMap := make(map[rune][]string)

	for i := 32; i < 127; i++ {
		char := rune(i)
		start := (i - 32) * 9

		if start+9 > len(splitfile) {
			return nil, fmt.Errorf("invalid input")
		}
		completelines := splitfile[start+1 : start+9]
		asciiMap[char] = completelines
	}
	return asciiMap, nil
}
func SplitInput(input string) []string {
	return strings.Split(input, "\\n")
}
func RenderLine(input string, banner map[rune][]string) []string {
	output := []string{}

	for row := 0; row < 8; row++ {
		var result strings.Builder
		for _, char := range input {
			newChar := banner[char][row]
			result.WriteString(newChar)
		}
		output = append(output, result.String())
	}
	return output
}
func ValidateInput(input string) (rune, error) {
	if len(input) == 0 {
		return 0, nil
	}
	for _, char := range input {
		if char < 32 || char > 127 {
			return 0, fmt.Errorf("unsupported input%q:", char)
		}

	}
	return 0, nil
}
func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	onlyNewLines := true
	for _, char := range input {
		if char != '\n' {
			onlyNewLines = false
		}
	}
	if onlyNewLines {
		return input
	}
	var result string
	parts := strings.Split(input, "\n")
	for index, words := range parts {
		if words == "" {
			if index != len(parts)-1 {
				result += "\n"
			}
			continue
		}
		output := RenderLine(words, banner)
		for _, lines := range output {
			result += lines + "\n"
		}
	}
	return result
}
