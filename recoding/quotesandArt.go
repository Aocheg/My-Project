package main 
// import(
// 	"regexp"
// )
// func ArtiQuotes(words string)string{

// 	re4 := regexp.MustCompile(`([a-zA-Z0-9])(['"])`)
// 	words = re4.ReplaceAllString(words, `$1 $2`)

// 	re5 := regexp.MustCompile(`(['"])([a-zA-Z0-9])`)
// 	words = re5.ReplaceAllString(words, `$1 $2`)

// 	re6 := regexp.MustCompile(`"\s*'\s*(.*?)\s*'\s*"`)
// 	words = re6.ReplaceAllString(words, `"'$1'"`)

// 	re7 := regexp.MustCompile(`\b([aA])\s+([aeiouhAEIOUH])`)
// 	words = re7.ReplaceAllString(words, "${1}n $2")

// 	words
// }