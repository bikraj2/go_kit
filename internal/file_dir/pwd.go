package filedir

import (
	"fmt"
	"strings"
)

type Pwd struct {
	CurrDir string
	PwdOptions
}
type PwdOptions struct {
	L bool
	P bool
}

var pwdValidOptions = []string{"l", "p"}

func HelpPwd() {
}
func (p *Pwd) ProcessCommand(args []string) error {

	defer p.resetFlags()
	err := p.processFlags(args[1:])
	if err != nil {
		return err
	}
	return nil
}

func (p *PwdOptions) processFlags(args []string) error {

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
		return ErrFlagCollision
	}
	switch strings.ToLower(opt) {
	case "l":
		pOpt.L = true
	case "p":
		pOpt.P = true
	default:
		return fmt.Errorf("%v is not a valid flag", opt)
	}
	return nil
}

func (pOpt *PwdOptions) flagSet() bool {
	return pOpt.L || pOpt.P
}

func (p *Pwd) resetFlags() {
	p.L = true
	p.P = false
}
