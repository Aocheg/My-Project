// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func RenderLine(input string, asciiMap map[rune][]string) {

// 	lines := strings.Split(input, "\n")

// 	for i, word := range lines {

// 		// Ignore the final empty split
// 		if word == "" {

// 			if i != len(lines)-1 {
// 				fmt.Println()
// 			}

// 			continue
// 		}

// 		output := make([]string, 8)

// 		for _, ch := range word {

// 			art := asciiMap[ch]

// 			for row := 0; row < 8; row++ {
// 				output[row] += art[row]
// 			}
// 		}

// 		for row := 0; row < 8; row++ {
// 			fmt.Println(output[row])
// 		}
// 	}
// }

package main

func RenderLine(input string, asciiMap map[rune][]string) []string {

	// Always return 8 lines
	output := make([]string, 8)

	// Empty input returns 8 empty strings
	if input == "" {
		return output
	}

	for _, ch := range input {

		art, ok := asciiMap[ch]
		if !ok {
			continue
		}

		for row := 0; row < 8; row++ {
			output[row] += art[row]
		}
	}

	return output
}
