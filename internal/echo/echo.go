package echo

import (
	"fmt"
	"io"
	"os"
	"strings"

	helper "go_kit.com/internal"
)

type Echo struct {
	Args     []string
	Text     string
	Redirect string
	FileName string
	EchoOptions
}

type EchoOptions struct {
	NewLine         bool
	EscapeCharacter bool
	ShowHelp        bool
}

var echoValidOptions = []string{"n", "e", "help"}

func (echo *Echo) Help() {
	helpText := `
Usage: echo [options] <string> [redirection] [file_name]

Options:
  -n    Do not output the trailing newline
  -e    Enable interpretation of backslash escapes

Redirection:
  >     Redirect output to a file (overwrite)
  >>    Append output to a file
  <     Read input from a file (not implemented)
  <<    Read input from a here-document (not implemented)

Examples:
  echo "Hello, World!"
  echo -n "No newline at the end"
  echo -e "Line1\nLine2"
  echo "Save to file" > output.txt
  echo "Append to file" >> output.txt
`

	fmt.Println(helpText)
}
func (echo *Echo) ProcessCommands(args []string) error {

	err := echo.processFlags(args)
	if echo.ShowHelp {
		echo.Help()
		return nil
	}
	if err != nil {
		return err
	}

	// Extract redirection and file name
	textParts := []string{}
	parsingFlags := true
	for i, arg := range args {

		if strings.HasPrefix(arg, "-") {
			if !parsingFlags {
				return fmt.Errorf("usage: echo [flag] <string> [redirection, file_name]")
			}
			continue
		}
		parsingFlags = false

		if arg == ">" || arg == ">>" || arg == "<" || arg == "<<" {
			if i+1 < len(args) {
				echo.Redirect = arg
				echo.FileName = args[i+1]
			} else {
				return fmt.Errorf("missing file name after redirection operator")
			}
			break
		}

		textParts = append(textParts, arg)
	}
	echo.Text = strings.Join(textParts, " ")
	// if echo.FileName {
	// 	if echo.Redirect != "" {
	// 		if echo.Redirect == ">" {
	// 		}
	// 	}
	// }

	var out io.Writer
	switch echo.Redirect {
	case "":
		out = os.Stdout
	case ">":
		file, err := os.OpenFile(echo.FileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
		out = file
	case ">>":
		file, err := os.OpenFile(echo.FileName, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
		out = file
	}
	// -e -> write to Buffer else Fprintf
	// -n -> write with line else dont
	if echo.NewLine {
		echo.Text = echo.Text + "\n"
	}
	if echo.EscapeCharacter {
		_, err = fmt.Fprintf(out, "%s", echo.Text)
	} else {
		// Convert the string to a slice of runes to handle escape sequences as literal characters
		_, err = fmt.Fprintf(out, "%s", echo.Text)
	}
	if err != nil {
		return err
	}
	return nil
}
func (echo *Echo) processFlags(args []string) error {
	start_flag_parse := false
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			start_flag_parse = true
			flag := strings.TrimPrefix(arg, "-")
			valid := helper.IsValidOptions(flag, echoValidOptions)
			if !valid {
				return fmt.Errorf("%v is not a valid flag", flag)
			}
			err := echo.setOption(flag)
			if err != nil {
				return err
			}
		} else if start_flag_parse {
			return nil
		}
	}
	return nil
}

func (echoOpt *EchoOptions) setOption(opt string) error {
	switch strings.ToLower(opt) {
	case "n":
		echoOpt.NewLine = true
	case "e":
		echoOpt.EscapeCharacter = true
	case "help":
		echoOpt.ShowHelp = true

	default:
		return fmt.Errorf("%v is not a valid flag", opt)
	}
	return nil
}
func (echoOpt *EchoOptions) resetFlag() {
	echoOpt.NewLine = true
	echoOpt.ShowHelp = false
	echoOpt.EscapeCharacter = false
	// echoOpt
}
