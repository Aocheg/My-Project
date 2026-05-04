package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {

	// Empty input
	if input == "" {
		return ""
	}

	// Input contains ONLY newlines
	onlyNewlines := true
	for _, ch := range input {
		if ch != '\n' {
			onlyNewlines = false
			break
		}
	}

	// Preserve exact newline count
	if onlyNewlines {
		return input
	}

	var result string

	// Split text into segments
	parts := strings.Split(input, "\n")

	for index, word := range parts {

		// Empty segment between words
		if word == "" {

			// Ignore trailing empty split created by strings.Split
			if index != len(parts)-1 {
				result += "\n"
			}

			continue
		}

		// Generate ASCII rows using your tested RenderLine
		output := RenderLine(word, banner)

		// Add all 8 rows
		for _, line := range output {
			result += line + "\n"
		}
	}

	return result
}
