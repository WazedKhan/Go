package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Random quote generator struct to manage quote
type RandomQuoteGenerator struct {
	quotes []string
}

// NewRandomQuoteGenerator initializes the quote generator with default quotes
func NewRandomQuoteGenerator() *RandomQuoteGenerator {
	return &RandomQuoteGenerator{
		quotes: []string{
			"Life is what happens when you're busy making other plans.",
			"The purpose of our lives is to be happy.",
			"Get busy living or get busy dying.",
		},
	}
}

// AddQuote adds a new quote to the list
func (r *RandomQuoteGenerator) AddQuote(quote string) {
	r.quotes = append(r.quotes, quote)
	fmt.Println("Quote added successfully!")
}

// DisplayQuotes displays all available quotes
func (r *RandomQuoteGenerator) DisplayQuotes() {
	if len(r.quotes) == 0 {
		fmt.Println("No quotes available!")
		return
	}
	fmt.Println("Available Quotes:")
	for idx, quote := range r.quotes {
		fmt.Printf("[%d]: %s\n", idx+1, quote)
	}
}

// GenerateRandomQuote picks a random quote
func (r *RandomQuoteGenerator) GenerateRandomQuote() {
	if len(r.quotes) == 0 {
		fmt.Println("No quotes available to display!")
		return
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(r.quotes))
	fmt.Printf("Random Quote: %s\n", r.quotes[randomIndex])
}

func main() {
	r := NewRandomQuoteGenerator()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Random Quote Generator!")

	for {
		fmt.Print("Enter a command (add, list, random, exit): ")
		scanner.Scan()
		command := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch command {
		case "add":
			fmt.Print("Enter your quote: ")
			scanner.Scan()
			quote := strings.TrimSpace(scanner.Text())
			if quote == "" {
				fmt.Println("Quote cannot be empty.")
			} else {
				r.AddQuote(quote)
			}
		case "list":
			r.DisplayQuotes()

		case "random":
			r.GenerateRandomQuote()

		case "exit":
			fmt.Println("Exiting the program. Have a great day!")
			return

		default:
			fmt.Println("Invalid command. Please use 'add', 'list', 'random', or 'exit'.")
		}
	}
}
