package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input         string
		caseSensitive bool
		expected      bool
	}{
		{"Radar", false, true},
		{"Radar", true, false},
		{"Level", false, true},
		{"Level", true, false},
		{"World", false, false},
		{"Madam", false, true},
		{"GoLang", false, false},
		{"", false, true},       // Edge case: empty string
		{"A", false, true},      // Edge case: single character
		{"12321", false, true},  // Numeric palindrome
		{"12345", false, false},
	}

	for _, test := range tests {
		result := IsPalindrome(test.input, test.caseSensitive)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q, caseSensitive=%v) = %v; expected %v", test.input, test.caseSensitive, result, test.expected)
		}
	}
}
