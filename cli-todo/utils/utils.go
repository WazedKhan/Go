package utils

import (
	"bufio"
	"os"
)

// NewScanner returns a bufio.Scanner instance for standard input
func NewScanner() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}
