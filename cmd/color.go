package main

import (
	"fmt"
)

// Stroring some regular color for easy import.

type Color struct {
	CurrentColor string
}

var colors = map[string]string{
	"Reset":   "\033[0m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"gray":    "\033[37m",
	"white":   "\033[97m",
}

func (clr *Color) initColor() error {
	clr.CurrentColor = colors["red"]
	return nil
}
func (clr *Color) processColor(args []string) error {
	if len(args) > 2 {
		return fmt.Errorf("%v is too many args to this function\nusage: color <name_color>", len(args)-1)
	}
	c := args[1]
	if c == "" {
		return fmt.Errorf("usage: color <name_color>")
	}
	color, exist := colors[c]
	if !exist {
		return fmt.Errorf("%v is not a valid color", c)
	}
	clr.CurrentColor = color
	return nil
}
