package filedir

import (
	"os"
)

func list_file(dir string) ([]os.DirEntry, error) {
	dirs_local, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return dirs_local, nil
}
