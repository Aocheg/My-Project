package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {

	// Read banner file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Convert file content to string
	content := string(data)

	// Handle Windows line endings
	content = strings.ReplaceAll(content, "\r\n", "\n")

	// Split file into lines
	lines := strings.Split(content, "\n")

	asciiMap := make(map[rune][]string)

	// ASCII printable chars start from space
	char := rune(32)

	// Each character:
	// 8 lines of art
	// 1 empty separator line
	for i := 0; i+8 <= len(lines) && char <= 126; i += 9 {

		// Copy exactly 8 lines
		asciiMap[char] = append([]string{}, lines[i:i+8]...)

		char++
	}

	// Validate total characters loaded
	if len(asciiMap) != 95 {
		return nil, fmt.Errorf("invalid banner file format")
	}

	return asciiMap, nil
}

// ascii-art
// -mod.go
// -banner.go
// -render.go --map
// -split.go
// -validate.go
// -generate.go
// -main.go
// -standard.txt
// -shadow.txt
// -thinkertoy.txt
