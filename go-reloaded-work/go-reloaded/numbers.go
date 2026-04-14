package main

import "strconv"

func HandleNumbers(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			value := words[i-1]

			decimal, err := strconv.ParseInt(value, 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(decimal, 10)
			}
			words = removeIndex(words, i)
			i--
		}
		if words[i] == "(bin)" && i > 0 {
			value := words[i-1]

			decimal, err := strconv.ParseInt(value, 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(decimal, 10)
			}
			words = removeIndex(words, i)
			i--
		}
	}
	return words

}
func removeIndex(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
