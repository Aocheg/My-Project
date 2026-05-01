package main

import (
	"fmt"
)

func main() {
	// input = strings.TrimSuffix(input, "\n")

	input, err := GetInputArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	if input == "" {
		return
	}
	// if input == "\n" {
	// 	fmt.Println()
	// 	return
	// }

	content, err := ReadBannerFile("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	lines := ParseBanner(content)
	asciiMap := BuildAsciiMap(lines)
	input = NormalizeInput(input)

	if err := ValidateInput(input, asciiMap); err != nil {
		fmt.Println(err)
		return
	}
	RenderText(input, asciiMap)
}
