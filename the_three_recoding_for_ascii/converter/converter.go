package main

import (
	"strings"
)

func StringToArt(input string) string {
	if input == "" {
		return ""
	}

	digits := map[rune][]string{
		'0': {
			" ___ ",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},
		'1': {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'2': {
			" ___ ",
			"    |",
			" ___|",
			"|    ",
			"|___ ",
		},
		'3': {
			"____ ",
			"    |",
			" ___|",
			"    |",
			"____|",
		},
		'4': {
			"|   |",
			"|___|",
			"    |",
			"    |",
			"    |",
		},
		'5': {
			" ____",
			"|    ",
			"|___ ",
			"    |",
			"____|",
		},
		'6': {
			" ___ ",
			"|    ",
			"|___ ",
			"|   |",
			"|___|",
		},
		'7': {
			"____ ",
			"    |",
			"   / ",
			"  /  ",
			" /   ",
		},
		'8': {
			" ___ ",
			"|   |",
			"|___|",
			"|   |",
			"|___|",
		},
		'9': {
			" ___ ",
			"|   |",
			"|___|",
			"    |",
			" ___|",
		},
	}

	lines := strings.Split(input, "\n")
	var result []string

	for _, line := range lines {
		asciiLines := make([]string, 5)

		for _, ch := range line {
			pattern, ok := digits[ch]
			if !ok {
				return ""
			}

			for i := 0; i < 5; i++ {
				asciiLines[i] += pattern[i]
			}
		}

		result = append(result, asciiLines...)
	}

	return strings.Join(result, "\n") + "\n"
}
