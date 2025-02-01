package main

import (
	"fmt"
	"strings"

	"go_kit.com/internal/Color"
	"go_kit.com/internal/echo"
)

// Funciton that is responsbile for figuring out if the string
// is commmand or not. If it is a valid command call appropriate
// funtion to handle that command

type Commands struct {
	color.Color
	echo.Echo
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
	default:
		fmt.Printf("%v is not a valid command", term.Cmds[0])
	}
}
