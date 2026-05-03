// package main

// import "strings"

// func ParseBanner(content string) []string {
// 	content = strings.ReplaceAll(content, "\r\n", "\n")
// 	return strings.Split(content, "\n")
// }

package main

import "strings"

func ParseBanner(content string) []string {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	return strings.Split(content, "\n")
}
