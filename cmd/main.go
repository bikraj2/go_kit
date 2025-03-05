package main

import (
	"fmt"
	"log"
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
		fullCommand := term.readFullCommand()
		args, err := parseArgs(fullCommand)
		if err != nil {
			log.Println(err)
		}
		term.Cmds = args
		term.processCommand()
	}
}
