package color

import (
	"fmt"
	"strings"
)

// Stroring some regular color for easy import.

type Color struct {
	CurrentColor string
	ResetColor   string
	ColorOptions
}
type ColorOptions struct {
	help bool
	set  bool
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

func (clr *Color) Help() {
	fmt.Println("Color CLI Tool Usage:")
	fmt.Println("Usage: color [options] [color_name]")
	fmt.Println("Options:")
	fmt.Println("  -help            Show this help message")
	fmt.Println("  -set <name> <hex_value>  Set a custom color with a hex value.")
	fmt.Println("                      <name> can be any name you choose, and <hex_value> should be a valid hex color.")
	fmt.Println("  <color_name>     Apply a predefined color to text. Available colors:")
	for name := range colors {
		fmt.Printf("    %v\n", name)
	}
	fmt.Println("\nExamples:")
	fmt.Println("  color -set myCustomColor #FF5733  Set a new custom color 'myCustomColor' with the hex value #FF5733.")
	fmt.Println("  color -help             Show this help message.")
	fmt.Println("  color red               Apply the 'red' color to text.")
	fmt.Println("  color yellow            Apply the 'yellow' color to text.")
	fmt.Println("Note: Colors applied using this tool are for text output in terminal.")
}
func (clr *Color) InitColor() error {
	clr.CurrentColor = colors["red"]
	clr.ResetColor = colors["Reset"]
	return nil
}
func (clr *Color) ProcessCommand(args []string) error {
	err := clr.processFlags(args)
	if err != nil {
		return err
	}
	if clr.set {
		clr.processSet(args[1:])
	} else if clr.help {
		clr.Help()
	} else {
		clr.processColor(args[1:])
	}
	return nil
}
func (clr *Color) processSet(args []string) error {
	if len(args) > 2 {
		return fmt.Errorf("%v is too many args to this function\nusage: color -set <name> <value>", len(args))
	}
	name := args[0]
	value := args[1]
	if len(name) > 32 {
		return fmt.Errorf("the length of name cannot be more than 32. it is now: %v", len(name))
	}

	// Check if valid HEX color.
	if !isValidHexColor(value) {
		return fmt.Errorf("%v is not a valid hex color ", value)
	}

	r, g, b, err := hexToRGB(value)
	if err != nil {
		return err

	}

	ansiCode := rgbToAnsi256(r, g, b)
	colors[name] = fmt.Sprintf("\033[%vm", ansiCode)
	return nil
}

func (clr *Color) processColor(args []string) error {

	if len(args) > 1 {
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

func (clr *Color) setOption(opt string) error {
	if clr.flagSet() {
		return FlagCollisionError
	}
	switch strings.ToLower(opt) {
	case "help":
		clr.help = true
	case "set":
		clr.set = true
	default:
		return fmt.Errorf("%v is not a valid flag", opt)
	}
	return nil
}
func (clr *Color) processFlags(args []string) error {
	argC := 0
	for i, arg := range args {
		if i == 0 {
			continue
		}
		if strings.HasPrefix(arg, "-") {
			flag := strings.TrimPrefix(arg, "-")
			valid := isValidOptions(flag)
			if !valid {
				return fmt.Errorf("%v is not a valid flag", flag)
			}
			argC += 1
			err := clr.setOption(flag)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (clr *Color) flagSet() bool {
	if clr.set {
		return true
	}
	if clr.help {
		return true
	}
	return false
}

// func ()
