package filedir

import (
	"fmt"
	"go_kit.com/internal/color"
	"strings"
)

type Ls struct {
	CurrDir string
	LsOptions
}
type LsOptions struct {
	MoreInfo        bool
	ShowHiddenFiles bool
}

var lsOptions = []string{"l", "a"}

func (l *Ls) ProcessCommand(args []string) error {
	defer l.resetFlags()
	err := l.processFlags(args)
	if err != nil {
		return err
	}

	dirs, err := list_file(l.CurrDir)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		file_info, err := dir.Info()
		if err != nil {
			panic(err)
		}
		if strings.HasPrefix(file_info.Name(), ".") && !l.ShowHiddenFiles {
			continue
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
		fmt.Println()
	}
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
	case "a":
		lOpt.ShowHiddenFiles = true
	default:
		return fmt.Errorf("%v is not a valid flag", opt)
	}
	return nil
}

func (lOpt *LsOptions) flagSet() bool {
	// if lOpt.MoreInfo {
	// 	return true
	// }
	// return false
	return false
}

func (l *Ls) resetFlags() {
	l.MoreInfo = false
	l.ShowHiddenFiles = false

}

// func ()
