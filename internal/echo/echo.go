package echo

import (
	"fmt"
	"strings"

	helper "go_kit.com/internal"
)

type Echo struct {
	Args []string
	EchoOptions
}
type EchoOptions struct {
	NewLine         bool
	EscapeCharacter bool
}

var echoValidOptions = []string{"n", "e"}

func (echo *Echo) ProcessCommands(args []string) error {
	fmt.Println(args)
	err := echo.processFlags(args)
	if err != nil {
		return err
	}
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			continue
		}
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
