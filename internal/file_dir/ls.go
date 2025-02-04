package filedir

import (
	"fmt"
	"strings"

	"go_kit.com/internal/color"
)

type Ls struct {
	CurrDir string
	LsOptions
}
type LsOptions struct {
	MoreInfo bool
}

var lsOptions = []string{"l"}

func (l *Ls) processCommand(args []string) error {
	err := l.processFlags(args)
	if err != nil {
		return err
	}

	dirs, err := list_file(l.CurrDir)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		fmt.Println()
		file_info, err := dir.Info()
		if err != nil {
			panic(err)
		}
		if l.MoreInfo {
			fmt.Printf("%-14s %-14v", file_info.ModTime().Format("02 Jan 2006"), file_info.Size())
		}
		if dir.Type().IsDir() {

			// Additional Information on directory.
			// Icon selection
			if file_info.Size() == 0 {
				fmt.Print("📂")
			} else {
				fmt.Print("📁")
			}
			fmt.Printf("%v%v%v", color.Colors["cyan2"], dir.Name(), color.Colors["reset"])
		} else {
			fmt.Printf("%v", dir.Name())
		}
	}
	fmt.Println()
	l.MoreInfo = false
	return nil
}

func (l *LsOptions) processFlags(args []string) error {

	for i, arg := range args {
		if i == 0 {
			continue
		}
		if strings.HasPrefix(arg, "-") {
			flag := strings.TrimPrefix(arg, "-")
			valid := isValidOptions(flag, lsOptions)
			if !valid {
				return fmt.Errorf("%v is not a valid flag", flag)
			}
			err := l.setOption(flag)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (lOpt *LsOptions) setOption(opt string) error {
	if lOpt.flagSet() {
		return ErrFlagCollision
	}
	switch strings.ToLower(opt) {
	case "l":
		lOpt.MoreInfo = true
	default:
		return fmt.Errorf("%v is not a valid flag", opt)
	}
	return nil
}

func (lOpt *LsOptions) flagSet() bool {
	if lOpt.MoreInfo {
		return true
	}
	return false
}

// func ()
