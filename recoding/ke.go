package main

import (
	//"fmt"
	"regexp"
	// "strings"
)

func Fix(s string) string {
	// s := strings.Join(strings.Fields(s)," ")

	res1 := regexp.MustCompile(`"\s*'\s*(.*?)\s*'\s*"`)
	s = res1.ReplaceAllString(s,`"'$1'"`)

	// res2 := regexp.MustCompile(`"\s*(.*?)\s*"`)
	// s = res2.ReplaceAllString(s,`"$1"`)


	// res3 := regexp.MustCompile(`\s+([.,?:;!...]+)`)
	// s = res3.ReplaceAllString(s,`$1`)

	res33 := regexp.MustCompile(`\s+([.,?:;!...])`)
	s = res33.ReplaceAllString(s,`$1`)

	res3 := regexp.MustCompile(`([.,?:;!...])([^\s])`)
	s = res3.ReplaceAllString(s,`$1 $2`)

	

	// res4 := regexp.MustCompile(`([.,?:;!...]+) ([a-zA-Z0-9])`)
	// s = res4.ReplaceAllString(s,`$1 $2`)

	res5 := regexp.MustCompile(`a\s+([aeiouhAEIOUH])`)
	s = res5.ReplaceAllString(s,`an $1`)

	res6 := regexp.MustCompile(`A\s+([aeiouhAEIOUH])`)
	s = res6.ReplaceAllString(s,`An $1`)

	res7 := regexp.MustCompile(`an\s+([^aeiouhAEIOUH])`)
	s = res7.ReplaceAllString(s,`a $1`)

	res8 := regexp.MustCompile(`An\s+([^aeiouhAEIOUH])`)
	s = res8.ReplaceAllString(s,`A $1`)

	return s
}