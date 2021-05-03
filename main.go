package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputNumbers := os.Args[1:]

	fmt.Println(numbersAsWords(inputNumbers))
}

func numbersAsWords(inputNumbers []string) string {
	var numberWords strings.Builder // string builder is more efficient for concatenation
	for i := 0; i < len(inputNumbers); i++ {
		for j := 0; j < len(inputNumbers[i]); j++ {
			convertedNumber, err := strconv.Atoi(inputNumbers[i][j : j+1]) // convert digit string to int
			if err != nil {
				return err.Error() // handle error if string wasn't a digit
			}
			fmt.Fprintf(&numberWords, convertDigitToWord(convertedNumber)) // add digit word to string
		}
		fmt.Fprintf(&numberWords, ",") // separate separate numbers with commas
	}
	return numberWords.String()[:numberWords.Len()-1] // Don't want the last comma
}

func convertDigitToWord(inputDigit int) string { // converts a single digit to its word equivalent
	switch inputDigit {
	case 0:
		return "Zero"
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	default:
		return "INPUT_NOT_DIGIT"
	}
}
