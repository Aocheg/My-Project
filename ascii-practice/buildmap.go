package main

func BuildAsciiMap(lines []string) map[rune][]string {
	asciiMap := make(map[rune][]string)

	char := ' '

	for i := 0; i+8 <= len(lines); i += 9 {
		asciiMap[char] = lines[i : i+8]
		char++
	}
	return asciiMap
}
