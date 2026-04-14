package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AlpHexBin(words string) string {
	text := strings.Fields(words)

	var result []string

	for _, val := range text {
		if strings.HasPrefix(val, "(") && strings.HasSuffix(val, ")") {
			newVal := val[1 : len(val)-1]
			line := strings.Split(newVal, ",")

			count := 1
			line2 := strings.TrimSpace(line[0])

			if len(line) > 1 {
				line3, err := strconv.Atoi(strings.TrimSpace(line[1]))
				if err == nil && line3 > count {
					count = line3
				}
			}
			start := max(len(result)-count, 0)

			for j := start; j < len(result); j++ {
				if line2 == "up" {
					result[j] = strings.ToUpper(result[j])
				}
				if line2 == "low" {
					result[j] = strings.ToLower(result[j])
				}
				if line2 == "hex" {
					val, err := strconv.ParseInt(result[j], 16, 64)
					if err != nil {
						fmt.Println(err)
					}
					result[j] = strconv.FormatInt(val, 10)
				}
				if line2 == "bin" {
					val, err := strconv.ParseInt(result[j], 2, 64)
					if err != nil {
						fmt.Println(err)
					}
					result[j] = strconv.FormatInt(val, 10)
				}
			}
		} else {
			result = append(result, val)
		}

	}
	return strings.Join(result, " ")
}
func main() {
	fmt.Println(AlpHexBin("GO (low) reloaded go (up) and she 10 (bin) had 1E (hex)"))
}
