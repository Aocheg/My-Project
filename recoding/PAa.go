package main

// import (
// 	"regexp"
// 	// "fmt"
// )

// func AP(words string) string {
// 	// handles punctuations
// 	re1 := regexp.MustCompile(`\s+([,.?!;:])`)
// 	words = re1.ReplaceAllString(words, `$1`)

// 	re2 := regexp.MustCompile(`([,.?!;:])([^\s])`)
// 	words = re2.ReplaceAllString(words, `$1 $2`)

// 	re33 := regexp.MustCompile(`\s+([,.?!;:]+)`)
// 	words = re33.ReplaceAllString(words, `$1`)

// 	// remove space after opening quote
// 	re0 := regexp.MustCompile(`(['"])\s+`)
// 	words = re0.ReplaceAllString(words, `$1`)

// 	// remove space before closing quote
// 	re01 := regexp.MustCompile(`\s+(['"])`)
// 	words = re01.ReplaceAllString(words, `$1`)

// 	// fix spacing outside quotes
// 	reFix1 := regexp.MustCompile(`([a-zA-Z0-9])(['"])`)
// 	words = reFix1.ReplaceAllString(words, `$1 $2`)

// 	reFix2 := regexp.MustCompile(`(['"])([a-zA-Z0-9])`)
// 	words = reFix2.ReplaceAllString(words, `$1 $2`)

// 	re := regexp.MustCompile(`"\s*'\s*(.*?)\s*'\s*"`)
// 	words = re.ReplaceAllString(words, `"'$1'"`)

// 	re4 := regexp.MustCompile(`\b([aA])\s+([aeiouhAEIOUH])`)
// 	words = re4.ReplaceAllString(words, "${1}n $2")

// 	return words
// }

// // func main (){
// // 	fmt.Println(AP("hello , world ! a apple ' hello '"))
// // }
