package main

import (
	"fmt"
	"strings"
)

// Funciton that is responsbile for figuring out if the string
// is commmand or not. If it is a valid command call appropriate
// funtion to handle that command

func (term *Terminal) processCommand(command string) {
	cmd := strings.Split(command, " ")
	term.Cmds = cmd
	switch term.Cmds[0] {
	case "color":
		err := term.processColor(term.Cmds)
		if err != nil {
			fmt.Println(err)
		}
	}
}
