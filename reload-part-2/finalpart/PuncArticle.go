package main

import (
	"regexp"
	"strings"
)

func PuncArticle(words string) string {

	// re00 := regexp.MustCompile(`\s*([,.?!;:]+)\s*`)
	// words = re00.ReplaceAllString(words, `$1 `)

	// (`\s*([,.?!;:]+)\s*`). (`$1 `)

	// fixQuote := regexp.MustCompile(`\b(\w+)\s*'\s*(\w+)\s*`)
	// words = fixQuote.ReplaceAllString(words, `$1'$2 `)

	// (`\b(\w+)\s*'\s*(\w+)\s*`).(`$1'$2`)

	// fixSingle := regexp.MustCompile(`'\s*([^']*?)\s*'`)
	// words = fixSingle.ReplaceAllString(words, `'${1}'`)

	// (`'\s*([^']*?)\s*'`).(`'${1}'`)

	// fixdouble := regexp.MustCompile(`"\s*([^"]*?)\s*"`)
	// words = fixdouble.ReplaceAllString(words, `"${1}"`)

	// (`"\s*([^"]*?)\s*"`).(`"${1}"`)

	// fixNested := regexp.MustCompile(`"\s*'\s*([^']*?)\s*'\s*"`)
	// words = fixNested.ReplaceAllString(words, `"'$1'"`)

	// (`"\s*'\s*([^']*?)\s*'\s*"`)

	// fix := regexp.MustCompile(`([a-zA-Z])'([a-zA-Z]{3,})'`)
	// words = fix.ReplaceAllString(words, `$1 '$2'`)

	// (`([a-zA-Z])'([a-zA-Z]{3,})'`).(`$1 '$2'`)

	// fixSpace := regexp.MustCompile(`[ ]{2,}`)
	// words = fixSpace.ReplaceAllString(words, " ")

	// fixNewline := regexp.MustCompile(` \n`)
	// words = fixNewline.ReplaceAllString(words, `\n`)

	// fixArticle := regexp.MustCompile(`\b([aA])\s+([aeiouhAEIOUH])`)
	// words = fixArticle.ReplaceAllString(words, `${1}n $2`)

	// removeN := regexp.MustCompile(`\b([aA])n\s+([^aeiouhAEIOUH\W])`)
	// words = removeN.ReplaceAllString(words, `${1} $2`)

	// words = strings.TrimSpace(words)

	// return words

	re00 := regexp.MustCompile(`\s*([,.?!;:]+)\s*`)
	words = re00.ReplaceAllString(words, `$1 `)

	re01 := regexp.MustCompile(`\b(\w+)\s*'\s*(\w+)\s*`)
	words = re01.ReplaceAllString(words, `$1'$2 `)

	re02 := regexp.MustCompile(`'\s*([^']*?)\s*'`)
	words = re02.ReplaceAllString(words, `'${1}'`)

	re03 := regexp.MustCompile(`"\s*([^"]*?)\s*"`)
	words = re03.ReplaceAllString(words, `"${1}"`)

	re04 := regexp.MustCompile(`"\s*'\s*([^']*?)\s*'\s*"`)
	words = re04.ReplaceAllString(words, `"'$1'"`)

	re05 := regexp.MustCompile(`([a-zA-Z])'([a-zA-Z]{3,})'`)
	words = re05.ReplaceAllString(words, `$1 '$2'`)

	re06 := regexp.MustCompile(`[ ]{2,}`)
	words = re06.ReplaceAllString(words, " ")

	re07 := regexp.MustCompile(` \n`)
	words = re07.ReplaceAllString(words, `\n`)

	re08 := regexp.MustCompile(`\b([aA])\s+([aeiouhAEIOUH])`)
	words = re08.ReplaceAllString(words, `${1}n $2`)

	re09 := regexp.MustCompile(`\b([aA])n\s+([^aeiouhAEIOUH\W])`)
	words = re09.ReplaceAllString(words, `${1} $2`)

	words = strings.TrimSpace(words)

	return words
}
