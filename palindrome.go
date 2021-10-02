package palindrome

import "strings"

const (
	CODE_SPACE byte = 32
)

func IsPalindrome(input string) bool {
	length := len(input)

	if length == 0 || length == 1 {
		return true
	}

	input = strings.ToLower(input)
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

		if isSpace(left) {
			leftIndex++
			continue
		}

		if isSpace(right) {
			rightIndex--
			continue
		}

		if left != right {
			result = false
			break
		}

		leftIndex++
		rightIndex--
	}

	return result
}

func isSpace(b byte) bool {
	return b == CODE_SPACE
}
