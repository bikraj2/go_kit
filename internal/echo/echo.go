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
}

var echoValidOptions = []string{"n", "e"}

func (echo *Echo) ProcessCommands(args []string) error {

	err := echo.processFlags(args)
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
	if echo.EscapeCharacter {
		_, err = fmt.Fprintln(out, echo.Text)
	} else {
		_, err = fmt.Fprintf(out, "%s\n", echo.Text)
	}
	if err != nil {
		return err
	}
	return nil
}
func (echo *Echo) processFlags(args []string) error {
	start_flag_parse := false
	for i, arg := range args {
		if i == 0 {
			continue
		}
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
	default:
		return fmt.Errorf("%v is not a valid flag", opt)
	}
	return nil
}
