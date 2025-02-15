package filedir

import (
	"fmt"
	"strings"
)

type MkDir struct {
	CurrDir  string
	FileMode string
	MkdirOptions
}
type MkdirOptions struct {
	ParentDir bool
	Mode      bool
	Verbose   bool
}

// -p	Create parent directories as needed (e.g., mkdir -p parent/child/grandchild)
// -m MODE	Set permissions (e.g., mkdir -m 755 new_dir)
// -v	Print a message for each created directory (verbose mode)
var validMkDirOptions = []string{"m", "p", "v"}

func (m *MkDir) ProcessCommand(args []string) error {
	defer m.resetFlags()
	err := m.processFlags(args)
	if err != nil {
		return err
	}
	args = args[1:]
	for i, arg := range args {
		if strings.HasPrefix(arg, "-") {
			continue
		}
		// fmt.Println(i, args[i])
		if i > 0 && args[i-1] == "-m" {
			// fmt.Println(args[i])
			m.FileMode = args[i]
			continue
		}
		err := create_dir(arg, m.ParentDir, m.FileMode)

		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MkdirOptions) processFlags(args []string) error {

	for i, arg := range args {
		if i == 0 {
			continue
		}
		if strings.HasPrefix(arg, "-") {
			flag := strings.TrimPrefix(arg, "-")
			valid := isValidOptions(flag, validMkDirOptions)
			if !valid {
				return fmt.Errorf("%v is not a valid flag", flag)
			}
			err := m.setOption(flag)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (mOpt *MkdirOptions) setOption(opt string) error {
	if mOpt.flagSet() {
		return ErrFlagCollision
	}
	switch strings.ToLower(opt) {
	case "p":
		mOpt.ParentDir = true
	case "m":
		mOpt.Mode = true
	case "v":
		mOpt.Verbose = true

	default:
		return fmt.Errorf("%v is not a valid flag", opt)
	}
	return nil
}

func (mOpt *MkdirOptions) flagSet() bool {
	// if lOpt.MoreInfo {
	// 	return true
	// }
	// return false
	return false
}

func (m *MkDir) resetFlags() {
	m.Verbose = false
	m.Mode = false
	m.ParentDir = false
}
