// // package main

// // import "strings"

// // func NormalizeInput(input string) string {
// // 	input = strings.ReplaceAll(input, "\r\n", "\n")
// // 	input = strings.TrimSuffix(input, "\n")

// // 	return input
// // }

// package main

// import "strings"

// func NormalizeInput(input string) string {

// 	// Convert Windows newlines to Unix style
// 	input = strings.ReplaceAll(input, "\r\n", "\n")

// 	// Do NOT remove all trailing newlines.
// 	// Only remove extra carriage returns.
// 	input = strings.TrimRight(input, "\r")

// 	return input
// }

package main

import "strings"

func NormalizeInput(input string) string {

	input = strings.ReplaceAll(input, "\r\n", "\n")

	// Convert literal \n into real newline
	input = strings.ReplaceAll(input, `\n`, "\n")

	return input
} // package main

// // import "strings"

// // func NormalizeInput(input string) string {
// // 	input = strings.ReplaceAll(input, "\r\n", "\n")
// // 	input = strings.TrimSuffix(input, "\n")

// // 	return input
// // }

// package main

// import "strings"

// func NormalizeInput(input string) string {

// 	// Convert Windows newlines to Unix style
// 	input = strings.ReplaceAll(input, "\r\n", "\n")

// 	// Do NOT remove all trailing newlines.
// 	// Only remove extra carriage returns.
// 	input = strings.TrimRight(input, "\r")

// 	return input
// }
