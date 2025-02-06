package filedir

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Cd struct {
	CurrDir string
	HomeDir string
}

func (c *Cd) ProcessCommand(args []string) (string, error) {
	if len(args) > 2 {
		return c.CurrDir, fmt.Errorf("%v is too many arguments\nUsage: cd <dir>", len(args)-1)
	}
	new_dir, err := c.processDir(args[1])
	return new_dir, err
}

func (c *Cd) processDir(dir string) (string, error) {
	if dir == "~" {
		return c.HomeDir, nil
	}
	if strings.HasPrefix(dir, "/") {
		return c.parseAbsPath(dir)
	}
	return c.parserRelativePath(dir)
}

func (c *Cd) parseAbsPath(dir string) (string, error) {
	stack := strings.Split(dir, "/")
	for _, part := range stack {
		switch part {
		case ".":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			// stack = append(stack, part)
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
		return "", ErrDirDoesnotExist
	}
	return new_dir, nil

}
func (c *Cd) parserRelativePath(dir string) (string, error) {
	parts := strings.Split(dir, "/")
	stack := strings.Split(c.CurrDir, "/")

	for _, part := range parts {
		switch part {
		case ".":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
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
		return "", ErrDirDoesnotExist
	}
	return new_dir, nil
}
func (c *Cd) changeToHome() (string, error) {
	return c.HomeDir, nil
}
