// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	if len(os.Args) != 3 {
// 		fmt.Println("Error Handling files: Usage: go run . input.txt output.txt")
// 	}
// 	inputfile := os.Args[1]
// 	outputfile := os.Args[2]

// 	input, err := os.ReadFile(inputfile)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	text := NumberToDecimal(words)

// 	err = os.WriteFile(outputfile, []byte(input), 0644)

// 	return

// }

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error Handling files: Usage: go run . input.txt output.txt")
		return
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]

	input, err := os.ReadFile(inputfile)
	if err != nil {
		log.Fatal(err)
	}

	content := string(input)
	words := strings.Fields(content)

	result := NumberToDecimal(words)

	output := strings.Join(result, " ")

	err = os.WriteFile(outputfile, []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
