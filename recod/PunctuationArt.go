package main

import (
	"regexp"
)
func Pa(words string)string{

	re00 := regexp.MustCompile(`\s+([,.?!;:])`)
	words = re00.ReplaceAllString(words, `$1`)

	re01 := regexp.MustCompile(`([,.?!;:])([^\s+])`)
	words = re01.ReplaceAllString(words, `$1 $2`)

	re02 := regexp.MustCompile(`\s+([,.?!;:]+)`)
	words = re02.ReplaceAllString(words, `$1`)

	re03 := regexp.MustCompile(`(['"])\s+`)
	words = re03.ReplaceAllString(words, `$1`)

	re04 := regexp.MustCompile(`\s+(['"])`)
	words = re04.ReplaceAllString(words, `$1`)

	re05 := regexp.MustCompile(`([a-zA-Z0-9])(['"])`)
	words = re05.ReplaceAllString(words, `$1 $2`)

	re06 := regexp.MustCompile(`(['"])([a-zA-Z0-9])`)
	words = re06.ReplaceAllString(words, `$1 $2`)

	re07 := regexp.MustCompile(`"\s*'\s*(.*?)\s*'\s*"`)
	words = re07.ReplaceAllString(words, `"'$1'"`)

	re08 := regexp.MustCompile(`"\s*(.*?)\s*"`)
	words = re08.ReplaceAllString(words, `"$1"`)

	re09 := regexp.MustCompile(`\b([aA])\s+([aeiouhAEIOUH])`)
	words = re09.ReplaceAllString(words, "${1}n $2")

	re10 := regexp.MustCompile(`\b([aA])n\s+([^aeiouhAEIOUH\W])`)
	words = re10.ReplaceAllString(words, "${1} $2")


	return words
}
