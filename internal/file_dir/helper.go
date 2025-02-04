package filedir

import (
	"errors"
	"os"
)

var (
	ErrFlagCollision = errors.New("flags cannot be set at the sametime")
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
