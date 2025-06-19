package palindrome

import "strings"

// IsPalindrome checks whether the given word is a palindrome.
// It accepts a caseSensitive flag to control case comparison.
func IsPalindrome(word string, caseSensitive bool) bool {
	if !caseSensitive {
		word = strings.ToLower(word)
	}

	runes := []rune(word)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			return false
		}
	}
	return true
}
