package main

import (
	"fmt"

	"github.com/FS-Frost/palindrome/palindrome"
)

func main() {
	input := ""
	n := 100000

	for i := 0; i < n; i++ {
		input += "a"
	}

	length := len(input)
	result := palindrome.IsPalindrome(input)
	maxLengthToShow := 10

	if length > maxLengthToShow {
		input = fmt.Sprintf("%s... (length: %d)", input[:maxLengthToShow], length)
	}

	fmt.Printf("<%s> is palindrome: %t\n", input, result)
}
