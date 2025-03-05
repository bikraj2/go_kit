package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrUnclosedQuote = errors.New("syntax error: unclosed quote")
)

func clear() {
	fmt.Print("\033[H\033[2J")
}

// hasUnclosedQuote: Check if there is a quotation mark that is unclosed
func hasUnclosedQuote(input string) bool {
	var singleQuotes, doubleQuotes int
	for _, char := range input {
		if char == '\'' {
			singleQuotes++
		} else if char == '"' {
			doubleQuotes++
		}
	}
	return (singleQuotes%2 != 0) || (doubleQuotes%2 != 0)
}

func (t *Terminal) readFullCommand() string {
	var input string
	for {
		if !t.InputScanner.Scan() {
			break
		}

		line := strings.TrimSpace(t.InputScanner.Text())
		input += " " + line
		if !hasUnclosedQuote(input) {
			break
		}
		fmt.Println("quote>")
	}
	return input

}

// parseArgs: Parses raw input into properly grouped arguments
func parseArgs(input string) ([]string, error) {
	var args []string
	var token strings.Builder
	inQuote := false
	quoteChar := byte(0) // Stores which quote is open (' or ")

	for i := 0; i < len(input); i++ {
		char := input[i]

		if char == '"' || char == '\'' {
			if inQuote {
				if char == quoteChar { // Closing the same quote type
					inQuote = false
					quoteChar = 0
				} else {
					token.WriteByte(char) // Keep nested different quotes
				}
			} else {
				inQuote = true
				quoteChar = char
			}
		} else if char == ' ' && !inQuote {
			// Space outside quotes means end of token
			if token.Len() > 0 {
				args = append(args, token.String())
				token.Reset()
			}
		} else {
			token.WriteByte(char)
		}
	}

	// If we exit loop with an open quote, return error
	if inQuote {
		return nil, ErrUnclosedQuote
	}

	// Append last token if exists
	if token.Len() > 0 {
		args = append(args, token.String())
	}

	return args, nil
}
func printASCII() {

	// Define ANSI color codes for different combinations
	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	blue := "\033[34m"
	magenta := "\033[35m"
	cyan := "\033[36m"
	white := "\033[37m"
	reset := "\033[0m"

	// Different color combinations
	colorCombos := [][]string{
		{red, yellow, white},   // Red, Yellow, White
		{blue, magenta, cyan},  // Blue, Magenta, Cyan
		{green, cyan, magenta}, // Green, Cyan, Magenta
		{yellow, white, blue},  // Yellow, White, Blue
		{cyan, red, green},     // Cyan, Red, Green
	}
	asciiArt := []string{
		"**********************************************",
		"*                                            *",
		"*     ____    ___      _  __  ___   _____    *",
		"*    / ___|  / _ \\    | |/ / |_ _| |_   _|   *",
		"*   | |  _  | | | |   | ' /   | |    | |     *",
		"*   | |_| | | |_| |   | . \\   | |    | |     *",
		"*    \\____|  \\___/    |_|\\_\\ |___|   |_|     *",
		"*                                            *",
		"**********************************************",
	}

	// Choose a color combination
	selectedCombo := colorCombos[2] // Change index to get different combos

	// Print the ASCII text with selected color combination
	for i, line := range asciiArt {
		fmt.Println(selectedCombo[i%len(selectedCombo)] + line + reset)
	}

}
