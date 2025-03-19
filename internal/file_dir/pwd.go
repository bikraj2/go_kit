package filedir

import (
	"fmt"
	"path/filepath"
	"strings"

	helper "go_kit.com/internal"
)

type Pwd struct {
	CurrDir string
	PwdOptions
}
type PwdOptions struct {
	L    bool
	P    bool
	Help bool
}

var pwdValidOptions = []string{"l", "p", "help"}

func HelpPwd() {

	fmt.Println(`Usage: pwd [OPTION]
Print the full filename of the current working directory.

Options:
  -L    Print the logical current working directory (default).
  -P    Print the physical current working directory (resolving symlinks).

Examples:
  pwd         # Prints the logical working directory.
  pwd -P      # Prints the physical working directory (resolves symlinks).
  pwd -L      # Explicitly prints the logical working directory.

Note:
  If both -L and -P are specified, an error will be thrown due to flag collision.`)
}
func (p *Pwd) ProcessCommand(args []string) error {
	// fmt.Println("here")
	defer p.resetFlags()

	err := p.processFlags(args)
	if err != nil {
		return err
	}
	if p.Help {
		HelpPwd()
		return nil
	}
	dir := p.CurrDir
	if p.P {

		dir, err = filepath.EvalSymlinks(p.CurrDir)
		// fmt.Println("Resolving physical link")
		if err != nil {
			return err
		}
	}

	fmt.Println(dir)
	return nil
}
func (p *PwdOptions) processFlags(args []string) error {

	for _, arg := range args {

		if strings.HasPrefix(arg, "-") {
			flag := strings.TrimPrefix(arg, "-")
			valid := helper.IsValidOptions(flag, pwdValidOptions)
			if !valid {
				return fmt.Errorf("%v is not a valid flag", flag)
			}
			err := p.setOption(flag)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("pwd doesnot accept any argument")
		}
	}
	return nil
}
func (pOpt *PwdOptions) setOption(opt string) error {
	if pOpt.flagSet() {
		return helper.ErrFlagCollision
	}
	switch strings.ToLower(opt) {
	case "l":
		pOpt.L = true
	case "p":
		pOpt.P = true
	case "help":
		pOpt.Help = true

	default:
		return fmt.Errorf("%v is not a valid flag", opt)
	}
	return nil
}

func (pOpt *PwdOptions) flagSet() bool {
	return pOpt.L || pOpt.P || pOpt.Help
}

func (p *Pwd) resetFlags() {
	p.L = false
	p.P = false
	p.Help = false
}
