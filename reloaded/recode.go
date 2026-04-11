package main

import (
	"fmt"

	"strconv"
	"strings"
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
func CapitalizeFirstWords(text string) string {
	words := strings.Fields(text)

	var result []string
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(cap") {

			clean := strings.Trim(words[i], "()")
			parts := strings.Split(clean, ",")

			n := 1
			if len(parts) > 1 {
				numStr := strings.TrimSpace(parts[1])
				num, _ := strconv.Atoi(numStr)
				n = num
			}

			for j := 1; j <= n && len(result)-j >= 0; j++ {
				word := result[len(result)-j]
				result[len(result)-j] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
			}
			continue
		}
		result = append(result, words[i])

	}
	return strings.Join(result, " ")

}

func main() {
	fmt.Println(NumberToDecimal([]string{"1E", "(hex)", "files", "wered", "added"}))
	fmt.Println(NumberToDecimal([]string{"10", "(bin)", "years"}))
	fmt.Println(CapitalizeFirstWords("go is very cool  (cap,3)"))
}
