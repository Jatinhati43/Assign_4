package palindrome

import "testing"

// TestIsPalindrome tests the IsPalindrome function with various input cases.
// It checks both typical and edge cases, including mixed case, empty strings,
// single characters, and numeric strings to verify correct palindrome detection.
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Radar", true},
		{"Level", true},
		{"World", false},
		{"Madam", true},
		{"GoLang", false},
		{"", true},       // Edge case: empty string
		{"A", true},      // Edge case: single character
		{"12321", true},  // Numeric palindrome
		{"12345", false},
	}

	for _, test := range tests {
		result := IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
