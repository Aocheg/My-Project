package main

func RenderLine(input string, banner map[rune][]string) []string {
	var result []string
	for i := 0; i < 8; i++ {
		word := ""
		for _, ch := range input {
			word += banner[ch][i]

		}
		result = append(result, word)
	}
	return result
}
