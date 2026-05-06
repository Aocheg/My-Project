package main

import (
	"fmt"
	"os"
)

func GetInputAndBanner() (string, string, error) {
	banner := "standard.txt"

	if len(os.Args) < 2 || len(os.Args) > 3 {
		return "", "", fmt.Errorf("Usage: go run . \"text\" [banner]")
	}
	input := os.Args[1]

	if len(os.Args) == 3 {
		switch os.Args[2] {
		case "standard":
			banner = "standard.txt"
		case "shadow":
			banner = "shadow.txt"
		case "thinkertoy":
			banner = "thinkertoy.txt"
		default:
			return "", "", fmt.Errorf("invalid banner type")
		}
	}
	return input, banner, nil
}
