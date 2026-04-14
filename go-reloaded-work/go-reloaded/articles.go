package main

import "strings"

func HandlesArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" || words[i] == "A" {
			next := words[i+1]

			if len(next) == 0 {
				continue
			}
			first := strings.ToLower(string(next[0]))

			if isVowel(first) || first == "h" {

				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}

	}
	return words

}
func isVowel(ch string) bool {
	return ch == "a" || ch == "e" || ch == "i" || ch == "o" || ch == "u"
}
