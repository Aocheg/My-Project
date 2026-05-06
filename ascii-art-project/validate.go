package main

import (
	"fmt"
)

func ValidateInput(input string, asciiMap map[rune][]string) error {

	for _, ch := range input {

		if ch == '\n' {
			continue
		}

		if _, ok := asciiMap[ch]; !ok {
			return fmt.Errorf("unsupported character: %q", ch)
		}
	}
	return nil
}
