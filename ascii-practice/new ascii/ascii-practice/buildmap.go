// package main

// func BuildAsciiMap(lines []string) map[rune][]string {
// 	asciiMap := make(map[rune][]string)

// 	char := ' '

// 	for i := 0; i+8 <= len(lines); i += 9 {
// 		asciiMap[char] = lines[i : i+8]
// 		char++
// 	}
// 	return asciiMap
// }

package main

func BuildAsciiMap(lines []string) map[rune][]string {
	asciiMap := make(map[rune][]string)

	char := ' '

	// Each character uses 8 lines + 1 separator line
	for i := 0; i+8 <= len(lines) && char <= '~'; i += 9 {
		asciiMap[char] = append([]string{}, lines[i:i+8]...)
		char++
	}

	return asciiMap
}
