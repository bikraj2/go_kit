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
		err := term.Commands.Color.ProcessCommand(term.Cmds)
		if err != nil {
			fmt.Println(err)
		}
		// case "echo":
		// err := term.
	case "ls":
		term.Commands.CurrDir = term.CurrDir
		err := term.FileDir.ProcessCommand(term.Cmds)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Printf("%v is not a valid command\n", term.Cmds[0])
	}
}
