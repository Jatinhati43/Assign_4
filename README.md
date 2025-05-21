# Palindrome Checker in Go

A simple command-line Go application that checks whether a given word is a palindrome, ignoring case sensitivity. This project follows the recommended [golang-standards/project-layout](https://github.com/golang-standards/project-layout) for scalable Go applications.

---
## Features

- Case-insensitive palindrome check
- Handles empty strings and numeric strings
- Clean, idiomatic Go code with unit tests
- Uses standard Go project layout for maintainability

---

- `cmd/main.go`: The application entry point, reads a word through command line.
- `internal/palidrome/palin_check.go`: Implements palindrome checking logic.
- `internal/palindrome/palin_check_test.go`: Unit tests for the palindrome logic.

---
## Example

Enter a word: Radar


'Radar' is a palindrome.

---
## 🚀 Getting Started

### Prerequisites

- Go 1.21 or later installed

### Running the Application

Build and run with the filename as argument:

```bash
go run ./cmd 


🧪 Running Tests

go test ./internal/...

✅ Run Tests with Coverage

go test -cover ./internal/...
