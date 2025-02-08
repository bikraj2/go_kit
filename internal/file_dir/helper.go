package filedir

import (
	"errors"
	"os"
)

var (
	ErrFlagCollision   = errors.New("flags cannot be set at the sametime")
	ErrDirDoesnotExist = errors.New("dir doesnot exist")
)

func list_file(dir string) ([]os.DirEntry, error) {

	// TODO: Fix the bug here the dir info is not passed correctly on the first call of this function.
	dirs_local, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return dirs_local, nil
}
func isValidOptions(option string, validOptions []string) bool {
	for _, opt := range validOptions {
		print(opt)
		if option == opt {
			return true
		}
	}
	return false
}
