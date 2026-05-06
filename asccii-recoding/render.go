package main

import "strings"

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
