package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to GoKit! A terminal-like interface.")
	fmt.Println("You can type regular commands. Type `exit` to quit.")

	// Central Terminal Struct to Handle Everything
	var term Terminal
	term.initTerminal()
	for {
		var dir string
		fmt.Println(term.CurrDir)
		if term.CurrDir != term.HomeDir {
			dir = strings.TrimLeft(term.CurrDir, term.HomeDir)
		} else {
			dir = term.CurrDir
		}
		fmt.Printf("%v~%v%v> ", term.Color.CurrentColor, dir, term.Color.ResetColor)

		if !term.InputScanner.Scan() {
			break
		}

		input := term.InputScanner.Text()
		input = strings.TrimSpace(input)

		term.processCommand(input)
	}
}
