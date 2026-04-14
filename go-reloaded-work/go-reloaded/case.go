package main

import (
	"strconv"
	"strings"
)

func HandleCase(words []string) []string {
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ")") {

			cmd, count := parseCaseCommand(words[i])

			if cmd == "" {
				continue
			}

			if count == 0 {
				count = 1
			}

			for j := 1; j <= count && i-j >= 0; j++ {
				switch cmd {
				case "up":
					words[i-j] = strings.ToUpper(words[i-j])

				case "low":
					words[i-j] = strings.ToLower(words[i-j])

				case "cap":
					words[i-j] = Capitalize(words[i-j])

				}
			}
			words = removeIndex(words, i)
			i--
		}

	}
	return words
}

func parseCaseCommand(s string) (string, int) {
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")

	cmd := strings.TrimSpace(parts[0])

	if cmd != "up" && cmd != "low" && cmd != "cap" {
		return "", 0
	}

	if len(parts) == 1 {
		return cmd, 0
	}

	n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return cmd, 0
	}
	return cmd, n
}
func Capitalize(words string) string {
	if len(words) == 0 {
		return words
	}
	first := strings.ToUpper(string(words[0]))
	second := strings.ToLower(words[1:])

	return first + second
}
