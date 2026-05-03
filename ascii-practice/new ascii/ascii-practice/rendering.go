// // package main

// // import (
// // 	"fmt"
// // 	"strings"
// // )

// // func RenderText(input string, asciiMap map[rune][]string) {

// // 	lines := strings.Split(input, "\n")

// // 	for i, word := range lines {
// // 		if word == "" {
// // 			fmt.Println()
// // 			continue
// // 		}

// // 		output := make([]string, 8)

// // 		// for row := 0; row < 8; row++ {
// // 		for _, ch := range word {
// // 			validChar, ok := asciiMap[ch]

// // 			if !ok {
// // 				fmt.Println("Error: Unsupported Character")
// // 				return
// // 			}
// // 			for row := 0; row < 8; row++ {
// // 				output[row] += validChar[row]
// // 			}

// // 		}
// // 		for row := 0; row < 8; row++ {
// // 			fmt.Println(output[row])

// // 		}
// // 		if i < len(lines)-1 && lines[i+1] != "" {
// // 			fmt.Println()
// // 		}
// // 	}
// // }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func RenderText(input string, asciiMap map[rune][]string) {

// 	lines := strings.Split(input, "\n")

// 	for i, word := range lines {

// 		// Handle empty lines correctly
// 		if word == "" {
// 			if len(lines) > 1 {
// 				fmt.Println("\n")
// 			}

// 			continue
// 		}

// 		output := make([]string, 8)

// 		for _, ch := range word {

// 			validChar, ok := asciiMap[ch]
// 			if !ok {
// 				fmt.Println("Error: Unsupported Character")
// 				return
// 			}

// 			for row := 0; row < 8; row++ {
// 				output[row] += validChar[row]
// 			}
// 		}

// 		for row := 0; row < 8; row++ {
// 			fmt.Println(output[row])
// 		}

// 		// Add spacing only between non-empty words
// 		if i < len(lines)-1 && lines[i+1] != "" {
// 			fmt.Println()
// 		}
// 	}
// }

package main

import (
	"fmt"
	"strings"
)

func RenderText(input string, asciiMap map[rune][]string) {

	lines := strings.Split(input, "\n")

	for i, word := range lines {

		// Handle empty lines
		if word == "" {

			// Do not print extra newline
			// after the final empty split
			if i != len(lines)-1 {
				fmt.Println()
			}

			continue
		}

		output := make([]string, 8)

		for _, ch := range word {

			validChar, ok := asciiMap[ch]
			if !ok {
				fmt.Println("Error: Unsupported Character")
				return
			}

			for row := 0; row < 8; row++ {
				output[row] += validChar[row]
			}
		}

		for row := 0; row < 8; row++ {
			fmt.Println(output[row])
		}
	}
}
