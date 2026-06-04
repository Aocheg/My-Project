// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func ValidateInput(s string) (rune, error) {
// 	if len(s) == 0 {
// 		return 0, fmt.Errorf("invalid character")
// 	}
// 	for _, char := range s {
// 		if char < 32 || char > 126 {
// 			return char, errors.New("invalid character")
// 		}
// 	}
// 	return 0, nil
// }

package main

import "errors"

func ValidateInput(s string) (rune, error) {

	// Empty string is valid
	if s == "" {
		return 0, nil
	}

	for _, char := range s {

		// Valid ASCII range: 32 -> 126
		if char < 32 || char > 126 {
			return char, errors.New("invalid character")
		}
	}

	return 0, nil
}
