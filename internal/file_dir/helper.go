package filedir

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"syscall"
)

var (
	ErrFlagCollision   = errors.New("flags cannot be set at the sametime")
	ErrDirDoesnotExist = errors.New("dir doesnot exist")
)

func list_file(dir string) ([]os.DirEntry, error) {

	dirs_local, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return dirs_local, nil
}
func isValidOptions(option string, validOptions []string) bool {
	for _, opt := range validOptions {
		if option == opt {
			return true
		}
	}
	return false
}

func create_dir(dir string, parent bool, fileMode string) error {
	decMode, err := strconv.ParseUint(fileMode, 8, 32)
	if err != nil {
		return err
	}
	fmt.Println()
	old_mask := syscall.Umask(0)
	fmt.Println(decMode)

	mode := os.FileMode(decMode)
	fmt.Println(mode)
	err = nil
	if parent {
		err = os.MkdirAll(dir, mode)
	} else {
		err = os.Mkdir(dir, mode)
	}

	syscall.Umask(old_mask)
	return err
}
