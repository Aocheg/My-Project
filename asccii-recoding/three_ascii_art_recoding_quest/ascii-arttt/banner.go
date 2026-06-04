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
		return nil, errors.New("Error reading file")
	}
	if len(file) == 0 {
		return nil, fmt.Errorf("empty file")
	}
	splitfile := strings.Split(string(file), "\n")
	asciiMap := make(map[rune][]string)
	for i := 32; i < 127; i++ {
		char := rune(i)
		start := (i - 32) * 9

		if start+9 > len(splitfile) {
			return nil, errors.New("invalid input")
		}
		completeLines := splitfile[start+1 : start+9]
		asciiMap[char] = completeLines
	}
	return asciiMap, nil

}
