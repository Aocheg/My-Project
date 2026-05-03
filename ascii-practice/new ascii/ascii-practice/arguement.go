// package main

// import (
// 	"fmt"
// 	"os"
// )

// func GetInputArgs() (string, error) {

// 	if len(os.Args) != 2 {
// 		return "", fmt.Errorf("usage: go run . \"Hello\"")
// 	}
// 	input := os.Args[1]

// 	// if input == "" {
// 	// 	return "", fmt.Errorf("error: empty input")
// 	// }
// 	return input, nil
// }

package main

import (
	"fmt"
	"os"
)

func GetInputArgs() (string, error) {

	if len(os.Args) != 2 {
		return "", fmt.Errorf("usage: go run . \"Hello\"")
	}

	input := os.Args[1]

	return input, nil
}
