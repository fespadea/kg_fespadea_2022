package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputNumbers := os.Args[1:]

	fmt.Println(numbersAsWords(inputNumbers))
}

func numbersAsWords(inputNumbers []string) string {
	numberWords := ""
	for i := 0; i < len(inputNumbers); i++ {
		convertedNumber, err := strconv.Atoi(inputNumbers[i])
		if err != nil {
			return err.Error()
		}
		numberWords += convertNumberToWords(convertedNumber)
		if i < len(inputNumbers)-1 {
			numberWords += ","
		}
	}
	return numberWords
}

func convertNumberToWords(inputNumber int) string {
	return ""
}
