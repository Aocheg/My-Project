package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {

	// Empty input must return empty string
	if input == "" {
		return ""
	}
	// Convert literal \n into real newlines
	input = strings.ReplaceAll(input, "\\n", "\n")

	// Input contains ONLY newlines
	onlyNewLines := true
	for _, char := range input {
		if char != '\n' {
			onlyNewLines = false
			break
		}
	}
	// Return the exact newline sequence
	if onlyNewLines {
		return input
	}
	var result string
	parts := strings.Split(input, "\n")
	for index, words := range parts {

		// Empty segment handling
		if words == "" {

			// Leading newline OR middle blank line
			if index != len(parts)-1 {
				result += "\n"
				continue
			}
			// Trailing newline must create 8 blank lines
			for i := 0; i < 8; i++ {
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
