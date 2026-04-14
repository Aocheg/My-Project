package main

func Process(words []string) []string {
	words = HandleNumbers(words)
	words = HandleCase(words)
	words = HandlesArticles(words)
	words = HandlePunctuation(words)
	words = HandleQuotes(words)

	return words
}
