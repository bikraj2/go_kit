package filedir

import (
	"errors"
	"fmt"
	"os"
	"strings"

	helper "go_kit.com/internal"
)

type Cd struct {
	CurrDir string
	HomeDir string
}

func HelpCd() {
	fmt.Println("\nCD COMMAND - CUSTOM IMPLEMENTATION")
	fmt.Println("Change the current directory to a specified path.")
	fmt.Println("USAGE:")
	fmt.Println("  cd <directory>")
	fmt.Println("OPTIONS:")
	fmt.Println("  cd <dir>   Change to the specified directory.")
	fmt.Println("  cd ~       Change to the home directory.")
	fmt.Println("  cd ..      Move one level up in the directory structure.")
	fmt.Println("  cd .       Stay in the current directory.")

	fmt.Println("PATH BEHAVIOR:")
	fmt.Println("  - Absolute paths (starting with '/') are directly processed.")
	fmt.Println("  - Relative paths are resolved based on the current directory.")
	fmt.Println("  - '..' moves one level up, and '.' represents the current directory.")

	fmt.Println("ERROR HANDLING:")
	fmt.Println("  - If the specified directory does not exist, an error is returned.")
	fmt.Println("  - Too many arguments will result in an error (e.g., `cd dir1 dir2`).")

	fmt.Println("EXAMPLES:")
	fmt.Println("  cd Documents      # Navigate to 'Documents' in the current directory.")
	fmt.Println("  cd /usr/local     # Navigate to an absolute path.")
	fmt.Println("  cd ~/Projects     # Navigate to 'Projects' inside the home directory.")
	fmt.Println("  cd ..             # Move one directory up.")
	fmt.Println("  cd .              # Stay in the current directory.")
}
func (c *Cd) ProcessCommand(args []string) (string, error) {
	if len(args) > 1 {
		return c.CurrDir, fmt.Errorf("%w %v is too many arguments\nUsage: cd <dir>", helper.ErrInvalidNoOfFlags, len(args)-1)
	}
	if strings.HasPrefix(args[0], "-") {
		if args[1] != "-help" {
			return "", fmt.Errorf("cd only supports -help flag")
		}
		HelpCd()
		return c.CurrDir, nil
	}
	new_dir, err := c.processDir(args[0])
	return new_dir, err
}

func (c *Cd) processDir(dir string) (string, error) {
	if dir == "~" {
		return c.HomeDir, nil
	}
	if strings.HasPrefix(dir, "/") {
		return c.parseAbsPath(dir)
	} else if strings.HasPrefix(dir, "~") {
		c.CurrDir = c.HomeDir
		return c.parserRelativePath(strings.TrimPrefix(dir, "~"))
	}
	return c.parserRelativePath(dir)
}

func (c *Cd) parseAbsPath(dir string) (string, error) {
	stack := strings.Split(dir, "/")
	for _, part := range stack {
		switch part {
		case ".":
			// Do nothing
		case "..":
			// Check if parent directory exists before moving up
			if len(stack) > 0 {
				parentPath := strings.Join(stack[:len(stack)-1], "/")
				if _, err := os.Stat(parentPath); os.IsNotExist(err) {
					return "", helper.ErrDirDoesnotExist
				}
				stack = stack[:len(stack)-1]
			}
		default:
			// Check if the directory exists before appending
			tempPath := strings.Join(append(stack, part), "/")
			if _, err := os.Stat(tempPath); os.IsNotExist(err) {
				return "", helper.ErrDirDoesnotExist
			}
			stack = append(stack, part)
		}

	}

	var new_dir string
	for _, part := range stack {
		if part == "" {
			continue
		}
		new_dir = new_dir + "/" + part
	}
	// Checking if the new directory is a valid directory.
	_, err := os.Stat(new_dir)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return "", helper.ErrDirDoesnotExist
	}
	return new_dir, nil

}
func (c *Cd) parserRelativePath(dir string) (string, error) {
	parts := strings.Split(dir, "/")
	stack := strings.Split(c.CurrDir, "/")
	for _, part := range parts {
		switch part {
		case ".":
			// Do nothing
		case "..":
			// Check if parent directory exists before moving up
			if len(stack) > 0 {
				parentPath := strings.Join(stack[:len(stack)-1], "/")
				if _, err := os.Stat(parentPath); os.IsNotExist(err) {
					return "", helper.ErrDirDoesnotExist
				}
				stack = stack[:len(stack)-1]
			}
		default:
			// Check if the directory exists before appending

			tempPath := strings.Join(append(stack, part), "/")
			if _, err := os.Stat(tempPath); os.IsNotExist(err) {
				return "", helper.ErrDirDoesnotExist
			}
			stack = append(stack, part)
		}
	}
	var new_dir string

	for _, part := range stack {
		if part == "" {
			continue
		}
		new_dir = new_dir + "/" + part
	}
	// Checking if the new directory is a valid directory.
	_, err := os.Stat(new_dir)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return "", helper.ErrDirDoesnotExist
	}
	return new_dir, nil
}

func (c *Cd) changeToHome() (string, error) {
	return c.HomeDir, nil
}
