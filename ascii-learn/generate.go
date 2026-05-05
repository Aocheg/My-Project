package main

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	if strings.ReplaceAll(input, "\\n", "") == "" {
		return strings.Repeat("\n", len(input)/2)
	}
	lines := SplitInput(input)
	var output []string
	for _, ch := range lines {
		if ch == "" {
			output = append(output, "")
			continue
		}
		if _, err := ValidateArt(ch); err != nil {
			return err.Error()
		}
		rendered := RenderLine(ch, banner)
		output = append(output, rendered...)
	}
	return strings.Join(output, "\n") + "\n"

}
