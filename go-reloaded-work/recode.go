package main

import (
	// "fmt"
	"strconv"
)

func NumberToDecimal(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" {
			if len(result) > 0 {
				if val, err := strconv.ParseInt(result[len(result)-1], 16, 64); err == nil {
					result[len(result)-1] = strconv.FormatInt(val, 10)
				}
			}

		} else if words[i] == "(bin)" {
			if len(result) > 0 {
				if val2, err := strconv.ParseInt(result[len(result)-1], 2, 64); err == nil {
					result[len(result)-1] = strconv.FormatInt(val2, 10)
				}
			}
		} else {
			result = append(result, words[i])

		}

	}
	return result
}

// func main() {
// 	fmt.Println(NumberToDecimal([]string{"1E", "(hex)", "files", "wered", "added"}))
// 	fmt.Println(NumberToDecimal([]string{"10", "(bin)", "years"}))
// }
