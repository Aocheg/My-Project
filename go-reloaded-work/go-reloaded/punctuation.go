package main

func HandlePunctuation(words []string) []string {
	var result []string
	for i := 0; i < len(words); i++ {
		word := words[i]

		if isPunctuationWord(word) {
			if len(result) > 0 {
				result[len(result)-1] = result[len(result)-1] + word
			} else {
				result = append(result, word)
			}
			continue
		}
		result = append(result, word)
	}
	return result

}

func isPunctuationWord(word string) bool {
	if len(word) > 1 {
		for i := 0; i < len(word); i++ {
			if !isPunctuation(rune(word[i])) {
				return false
			}
		}
		return len(word) > 0
	}
	if len(word) == 1 && isPunctuation(rune(word[0])) {
		return true
	}
	return false
}
