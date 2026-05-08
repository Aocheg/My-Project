package main

import (
	"fmt"
)

func main() {

	input, banner, err := GetInputAndBanner()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print nothing for empty input
	if input == "" {
		return
	}

	// Normalize input text
	input = NormalizeInput(input)

	// Load banner file
	asciiMap, err := LoadBanner(banner)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Validate characters
	if err := ValidateInput(input, asciiMap); err != nil {
		fmt.Println(err)
		return
	}

	// Render ASCII art
	RenderText(input, asciiMap)
}