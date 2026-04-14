package main

func SplitWords(text string) []string {
	var words []string
	var current []rune

	for i := 0; i < len(text); i++ {
		ch := rune(text[i])

		if text[i] == ' ' || text[i] == '\t' || text[i] == '\n' {
			if len(current) > 0 {
				words = append(words, string(current))
				current = nil
			}
			continue
		}
		if ch == '(' {
			if len(current) > 0 {
				words = append(words, string(current))
				current = nil
			}
			var cmd []rune
			cmd = append(cmd, ch)

			i++
			for i < len(text) {
				cmd = append(cmd, rune(text[i]))
				if text[i] == ')' {
					break
				}
				i++
			}
			words = append(words, string(cmd))
			continue
		}
		if ch == '\'' {
			if len(current) > 0 {
				words = append(words, string(current))
				current = nil
			}
			words = append(words, string(ch))
			continue
		}
		if isPunctuation(ch) {
			if len(current) > 0 {
				words = append(words, string(current))
				current = nil
			}
			var punct []rune
			punct = append(punct, ch)

			for i+1 < len(text) && isPunctuation(rune(text[i+1])) {
				i++
				punct = append(punct, rune(text[i]))
			}
			words = append(words, string(punct))
			continue
		}
		current = append(current, ch)
	}
	if len(current) > 0 {
		words = append(words, string(current))

	}
	return words

}
func isPunctuation(ch rune) bool {
	switch ch {
	case '.', ',', '!', '?', ':', ';':
		return true
	}
	return false

}
