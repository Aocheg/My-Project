package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file")
	}

	if len(file) == 0 {
		return nil, errors.New("empty file")
	}

	splitfile := strings.Split(string(file), "\n")
	asciiMap := make(map[rune][]string)

	for i := 32; i < 127; i++ {
		char := rune(i)
		start := (i - 32) * 9

		if start+9 > len(splitfile) {
			return nil, errors.New("invalid banner format")
		}

		completelines := splitfile[start+1 : start+9]
		asciiMap[char] = completelines

	}
	return asciiMap, nil
}
