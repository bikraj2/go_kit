package main

import (
	"fmt"
	"strings"

	"go_kit.com/internal/color"
	"go_kit.com/internal/echo"
	filedir "go_kit.com/internal/file_dir"
)

// Funciton that is responsbile for figuring out if the string
// is commmand or not. If it is a valid command call appropriate
// funtion to handle that command

type Commands struct {
	echo.Echo
	color.Color
	filedir.FileDir
}

func (term *Terminal) processCommand(command string) {

	cmd := strings.Split(command, " ")
	term.Cmds = cmd

	switch term.Cmds[0] {
	case "color":
		err := term.Color.ProcessCommand(term.Cmds)
		if err != nil {
			fmt.Println(err)
		}
		// case "echo":
		// err := term.
	case "ls":
		err := term.Ls.ProcessCommand(term.Cmds)
		term.Ls.CurrDir = term.CurrDir
		if err != nil {
			fmt.Println(err)
		}

	case "cd":
		term.Cd.CurrDir = term.CurrDir
		new_dir, err := term.Cd.ProcessCommand(term.Cmds)
		if err != nil {
			fmt.Println(err)
			return
		}
		term.CurrDir = new_dir
	default:
		fmt.Printf("%v is not a valid command\n", term.Cmds[0])
	}
}
