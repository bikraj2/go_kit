package helper

import (
	"errors"
	"os"
	"strconv"
	"syscall"
)

var (
	ErrFlagCollision    = errors.New("flags cannot be set at the sametime")
	ErrInvalidNoOfFlags = errors.New("")
	ErrDirDoesnotExist  = errors.New("dir doesnot exist")
)

func List_file(dir string) ([]os.DirEntry, error) {

	dirs_local, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return dirs_local, nil
}
func IsValidOptions(option string, validOptions []string) bool {
	for _, opt := range validOptions {
		if option == opt {
			return true
		}
	}
	return false
}

func Create_dir(dir string, parent bool, fileMode string) error {
	decMode, err := strconv.ParseUint(fileMode, 8, 32)
	if err != nil {
		return err
	}
	old_mask := syscall.Umask(0)
	mode := os.FileMode(decMode)
	err = nil
	if parent {
		err = os.MkdirAll(dir, mode)
	} else {
		err = os.Mkdir(dir, mode)
	}
	syscall.Umask(old_mask)
	return err
}
