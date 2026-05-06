package main

import (
	"fmt"
)

func ValidateArt(input string) (rune, error) {
	if len(input) == 0 {
		return 0, nil
	}
	for _, ch := range input {
		if ch < 32 || ch > 126 {
			return ch, fmt.Errorf("unsupported input%q", ch)
		}
	}
	return 0, nil
}
