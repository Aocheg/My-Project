package main

import (
	"fmt"
	"strings"
)

func RenderLine(input string, asciiMap map[rune][]string) {

	lines := strings.Split(input, "\n")

	for i, word := range lines {

		// Ignore the final empty split
		if word == "" {

			if i != len(lines)-1 {
				fmt.Println()
			}

			continue
		}

		output := make([]string, 8)

		for _, ch := range word {

			art := asciiMap[ch]

			for row := 0; row < 8; row++ {
				output[row] += art[row]
			}
		}

		for row := 0; row < 8; row++ {
			fmt.Println(output[row])
		}
	}
}