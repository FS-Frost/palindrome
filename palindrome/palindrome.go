package palindrome

import "fmt"

func IsPalindrome(input string) bool {
	length := len(input)

	if length == 0 || length == 1 {
		return true
	}

	leftIndex := 0
	rightIndex := length - 1
	result := true
	var left byte
	var right byte

	for {
		if leftIndex > rightIndex {
			break
		}

		left = input[leftIndex]
		right = input[rightIndex]

		checkForSpace(left, leftIndex)
		checkForSpace(right, rightIndex)

		if left != right {
			result = false
			break
		}

		leftIndex++
		rightIndex--
	}

	return result
}

func checkForSpace(b byte, pos int) {
	var spaceCode byte = 32
	if b == spaceCode {
		msg := fmt.Sprintf("only words allowed, space detected at column %d", pos)
		panic(msg)
	}
}
