package color

import (
	"fmt"
	"strings"
)

// Stroring some regular color for easy import.

type Color struct {
	CurrentColor string
	ResetColor   string
}
type ColorOptions struct {
	help bool
	set  bool
}

var colors = map[string]string{
	"reset":   "\033[0m",
	"red":     "\033[38;2;255;0;0m",
	"green":   "\033[38;2;0;255;0m",
	"blue":    "\033[38;2;0;0;255m",
	"yellow":  "\033[38;2;255;255;0m",
	"magenta": "\033[38;2;255;0;255m",
	"cyan":    "\033[38;2;0;255;255m",
	"white":   "\033[38;2;255;255;255m",
	"gray":    "\033[38;2;128;128;128m",
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
	clr.ResetColor = colors["reset"]
	return nil
}
func (clr *Color) ProcessCommand(args []string) error {
	var option ColorOptions
	err := option.processFlags(args)
	if err != nil {
		return err
	}
	if option.set {
		err = clr.processSet(args[1:])
	} else if option.help {
		clr.Help()
	} else {
		err = clr.processColor(args)
	}

	if err != nil {
		return err
	}
	return nil
}
func (clr *Color) processSet(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("%v is too few args to this function\nusage: color -set <name> <value>", len(args)-1)
	}
	if len(args) > 3 {
		return fmt.Errorf("%v is too many args to this function\nusage: color -set <name> <value>", len(args)-1)
	}
	name := args[1]
	value := args[2]
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

	ansiCode := rgbToAnsiTrueColor(r, g, b)
	colors[name] = fmt.Sprintf("\033[%vm", ansiCode)
	return nil
}

func (clr *Color) processColor(args []string) error {

	if len(args) < 2 {
		return fmt.Errorf("%v is too few args to this function\nusage: color <name_color>", len(args)-1)
	}
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

func (option *ColorOptions) setOption(opt string) error {
	if option.flagSet() {
		return ErrFlagCollision
	}
	switch strings.ToLower(opt) {
	case "help":
		option.help = true
	case "set":
		option.set = true
	default:
		return fmt.Errorf("%v is not a valid flag", opt)
	}
	return nil
}
func (option *ColorOptions) processFlags(args []string) error {
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
			err := option.setOption(flag)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (opt *ColorOptions) flagSet() bool {
	if opt.set {
		return true
	}
	if opt.help {
		return true
	}
	return false
}

// func ()
