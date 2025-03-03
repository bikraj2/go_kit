package filedir

import (
	"fmt"
	"strconv"
	"strings"

	helper "go_kit.com/internal"
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
	Help      bool
}

// -p	Create parent directories as needed (e.g., mkdir -p parent/child/grandchild)
// -m MODE	Set permissions (e.g., mkdir -m 755 new_dir)
// -v	Print a message for each created directory (verbose mode)

var validMkDirOptions = []string{"m", "p", "v", "help"}

// HelpMkDir provides detailed usage information for the mkdir command
func HelpMkDir() {
	fmt.Println("\nMKDIR COMMAND - CUSTOM IMPLEMENTATION")
	fmt.Println("Create directories with various options.")

	fmt.Println("USAGE:")
	fmt.Println("  mkdir [OPTIONS] <directory>")

	fmt.Println("OPTIONS:")
	fmt.Println("  -p       Create parent directories as needed (e.g., mkdir -p parent/child/grandchild).")
	fmt.Println("  -m MODE  Set permissions (e.g., mkdir -m 755 new_dir).")
	fmt.Println("  -v       Print a message for each created directory (verbose mode).")

	fmt.Println("BEHAVIOR:")
	fmt.Println("  - If the directory already exists, an error is returned (unless -p is used).")
	fmt.Println("  - If -p is used, all missing parent directories will be created automatically.")
	fmt.Println("  - If -m is used, the directory will be created with the specified permissions.")
	fmt.Println("  - If -v is used, a message is printed for each directory created.")

	fmt.Println("ERROR HANDLING:")
	fmt.Println("  - Invalid options return an error.")
	fmt.Println("  - If the provided directory path is invalid, an error is returned.")
	fmt.Println("  - If -m is provided without a valid mode, an error is returned.")

	fmt.Println("EXAMPLES:")
	fmt.Println("  mkdir new_folder         # Create 'new_folder' in the current directory.")
	fmt.Println("  mkdir -p parent/child    # Create 'parent' and 'child' directories if they don't exist.")
	fmt.Println("  mkdir -m 755 mydir       # Create 'mydir' with 755 permissions.")
	fmt.Println("  mkdir -v project         # Create 'project' and print a confirmation message.")
	fmt.Println("  mkdir -p -m 700 data/logs # Create 'data/logs' with 700 permissions if missing.")
}

func (m *MkDir) ProcessCommand(args []string) error {
	defer m.resetFlags()
	err := m.processFlags(args)
	if err != nil {
		return err
	}
	if m.Help {
		HelpMkDir()
		return nil
	}
	args = args[1:]
	count := 0
	for i, arg := range args {
		if strings.HasPrefix(arg, "-") {
			continue
		}
		// fmt.Println(i, args[i])
		if i > 0 && args[i-1] == "-m" {
			if _, err := strconv.ParseInt(args[i], 10, 32); err != nil {
				return fmt.Errorf("enter Valid file mode")
			}
			m.FileMode = args[i]
			continue
		}
		count += 1
		err := helper.Create_dir(arg, m.ParentDir, m.FileMode)
		if err != nil {
			return err
		}
		if m.Verbose {
			fmt.Printf("%s created.\n", arg)
		}

	}
	if count == 0 {
		return fmt.Errorf("usage: mkdir [options] <dirs>")
	}
	return nil
}
func (m *MkdirOptions) processFlags(args []string) error {

	start_flag_parse := false
	for i, arg := range args {
		if i == 0 {
			continue
		}
		if strings.HasPrefix(arg, "-") {
			start_flag_parse = true
			flag := strings.TrimPrefix(arg, "-")
			valid := helper.IsValidOptions(flag, validMkDirOptions)
			if !valid {
				return fmt.Errorf("%v is not a valid flag", flag)
			}
			err := m.setOption(flag)
			if err != nil {
				return err

			}
		} else if start_flag_parse {
			return nil
		}
	}
	return nil
}

func (mOpt *MkdirOptions) setOption(opt string) error {
	if mOpt.flagSet() && opt == "help" || mOpt.Help {
		return helper.ErrFlagCollision
	}
	switch strings.ToLower(opt) {
	case "help":
		mOpt.Help = true
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
	// If Help is set, no other flag should be set
	return mOpt.ParentDir || mOpt.Mode || mOpt.Verbose
}

func (m *MkDir) resetFlags() {
	m.Verbose = false
	m.Mode = false
	m.ParentDir = false
	m.Help = false
}
