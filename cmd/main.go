package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"ASSIGN_4/internal/palindrome"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a word: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Print("Case sensitive? (y/n): ")
	caseInput, _ := reader.ReadString('\n')
	caseInput = strings.TrimSpace(caseInput)

	caseSensitive := strings.ToLower(caseInput) == "y"

	if palindrome.IsPalindrome(input, caseSensitive) {
		fmt.Printf("'%s' is a palindrome.\n", input)
	} else {
		fmt.Printf("'%s' is not a palindrome.\n", input)
	}
}
