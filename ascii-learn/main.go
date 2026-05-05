package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run . <text>")
	}
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("error", err)
	}
	filename := os.Args[1]
	result := GenerateArt(filename, banner)
	fmt.Print(result)
}
