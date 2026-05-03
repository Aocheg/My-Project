package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func ParseArgs(input string) string {

// 	if input == "" {
// 		return ""
// 	}
// 	if input == "\n" {
// 		fmt.Println()
// 	}

// 	args := os.Args

// 	fmt.Println(len(args))

// 	for i, val := range args {
// 		fmt.Println(i, val)
// 	}
// 	return ""
// }

// func main() {

// 	if len(os.Args) != 2 {
// 		fmt.Println("Usage: go run . \"Hello\"")
// 		return
// 	}
// 	input := os.Args[1]
// 	// fmt.Println(input)

// 	data, err := os.ReadFile("standard.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	content := string(data)
// 	// fmt.Println(content)

// 	content = strings.ReplaceAll(content, "\r\n", "\n")

// 	blocks := strings.Split(content, "\n")
// 	// fmt.Println((blocks))

// 	asciiMap := map[rune][]string{}

// 	char := ' '

// 	for i := 0; i+8 <= len(blocks); i += 9 {
// 		charBlock := blocks[i : i+8]

// 		asciiMap[char] = charBlock
// 		char++

// 	}
// 	// for _, line := range asciiMap['A'] {
// 	// 	fmt.Println(line)
// 	// }

// 	lines := strings.Split(input, "\\n")

// 	for _, word := range lines {
// 		for row := 0; row < 8; row++ {
// 			for _, ch := range word {
// 				validChar, ok := asciiMap[ch]

// 				if !ok {
// 					fmt.Println("Error: Unsupported Character")
// 					return
// 				}
// 				fmt.Print(validChar[row])
// 			}
// 			fmt.Println()
// 		}

// 	}

// 	// fmt.Println(asciiMap['A'])

// }
