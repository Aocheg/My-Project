package main

import (
	"regexp"
	"strings"
)

func PuncArticle(word string) string {

	re00 := regexp.MustCompile(`\s*([,.?!;:]+)\s*`)
	word = re00.ReplaceAllString(word, `$1 `)

	re01 := regexp.MustCompile(`\b(\w+)\s*'\s*(\w+)\b`)
	word = re01.ReplaceAllString(word, `$1'$2`)

	re02 := regexp.MustCompile(`[ ]{2,}`)
	word = re02.ReplaceAllString(word, " ")

	re03 := regexp.MustCompile(` \n`)
	word = re03.ReplaceAllString(word, "\n")

	word = strings.TrimSpace(word)

	re04 := regexp.MustCompile(`\b([aA])\s+([aeiouhAEIOUH(])`)
	word = re04.ReplaceAllString(word, `${1}n $2`)

	re05 := regexp.MustCompile(`\b([aA])n\s+([^aeiouhAEIOUH(\W])`)
	word = re05.ReplaceAllString(word, `${1} $2`)

	return word
}
