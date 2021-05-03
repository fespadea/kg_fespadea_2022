package main

import (
	"fmt"
	"os"
)

func main() {
	inputNumbers := os.Args[1:]

	fmt.Println(numbersAsWords(inputNumbers))
}

func numbersAsWords(inputNumbers []string) string {
	return ""
}
