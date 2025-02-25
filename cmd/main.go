package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to GoKit! A terminal-like interface.")
	fmt.Println("You can type regular commands. Type `exit` to quit.")
	printASCII()
	fmt.Println()
	// Central Terminal Struct to Handle Everything
	var term Terminal
	term.initTerminal()
	for {
		term.printDir()
		if !term.InputScanner.Scan() {
			break
		}

		input := term.InputScanner.Text()
		input = strings.TrimSpace(input)

		term.processCommand(input)
	}
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
