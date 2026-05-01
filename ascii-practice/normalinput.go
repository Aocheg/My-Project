package main

import "strings"

func NormalizeInput(input string) string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.TrimSuffix(input, "\n")

	return input
}
