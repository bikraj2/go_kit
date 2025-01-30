package main

import (
	"bufio"
	"os"
)

type Info struct {
	CurrDir string
	HomeDir string
}
type Commands struct {
	Cmds []string
}

type Terminal struct {
	Color
	Info
	Commands
	InputScanner *bufio.Scanner
}

func (info *Info) initInfo() {
	// Starting Curr Directory
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//Home Directory of the user to remove unnecasary dir info.
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	info.CurrDir = dir
	info.HomeDir = homeDir
}
func (term *Terminal) initTerminal() {

	// Scanner to keep reading commands/input from the user.
	scanner := bufio.NewScanner(os.Stdin)

	/// Assining newly Initialized.
	term.InputScanner = scanner

	// Initaliing Dependent Struct
	term.initInfo()
	term.initColor()
}
