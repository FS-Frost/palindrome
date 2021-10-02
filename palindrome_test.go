package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	assert(t, "palíndromo", false)
	assert(t, "casa", false)
	assert(t, "oso", true)
	assert(t, "Oso", true)
	assert(t, "dabalearrozalazorraelabad", true)
	assert(t, "ataralarata", true)
	assert(t, "amor a roma", true)
	assert(t, "vacío", false)
	assert(t, "a", true)
	assert(t, "12321", true)
	assert(t, "1221", true)
	assert(t, "reconocer", true)
	assert(t, "12322", false)
	assertLongWord(t)
}

func assert(t *testing.T, input string, expected bool) {
	actual := IsPalindrome(input)

	if actual != expected {
		t.Errorf("[%s is palindrome] Should be %t", input, expected)
	}
}

func assertLongWord(t *testing.T) {
	input := ""
	n := 1000

	for i := 0; i < n; i++ {
		input += "a"
	}
	assert(t, input, true)
}
