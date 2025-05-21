package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"ASSIGN_4/internal/palindrome"
)

// main is the entry point of the application.
// It reads a word from standard input and checks whether the word is a palindrome.
// The result is printed to the terminal.
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a word: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if palindrome.IsPalindrome(input) {
		fmt.Printf("'%s' is a palindrome.\n", input)
	} else {
		fmt.Printf("'%s' is not a palindrome.\n", input)
	}
}
