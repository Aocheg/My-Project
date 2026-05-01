package main

import (
	"fmt"
	"os"
	"strings"
)

func ParseArgs(input string) string {

	args := os.Args

	fmt.Println(len(args))

	for i, val := range args {
		fmt.Println(i, val)
	}
	return ""
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"Hello\"")
		return
	}
	input := os.Args[1]
	fmt.Println(input)

	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	content := string(data)
	// fmt.Println(content)

	content = strings.ReplaceAll(content, "\r\n", "\n")

	blocks := strings.Split(content, "\n")
	// fmt.Println((blocks))

	clean := []string{}


	for i := 0; i+8 <= len(clean); i += 9 {
		charBlock := clean[i : i+8]
		fmt.Println(charBlock)
	}

	// fmt.Println(ParseArgs("Hello"))

}
