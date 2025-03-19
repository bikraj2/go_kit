package filedir_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	helper "go_kit.com/internal"
	filedir "go_kit.com/internal/file_dir"
)

func Test_cd(t *testing.T) {
	// Test the cd .. ../../ ./ dir_1/dir_2/../dir_3

	// Creating second level of dirs
	// Creating third level of dirs
	setupDirs()
	defer deleteDirs()

	currDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name           string
		args           []string
		expected_error error
		expected       string
	}{
		{
			name:     "no change",
			args:     []string{"."},
			expected: currDir,
		},
		{
			name:     "into one nested dir",
			args:     []string{"./test_env"},
			expected: filepath.Join(currDir, "test_env"),
		},
		{
			name:     "into multiple nested dir",
			args:     []string{"./test_env/first_1/second_1/third_1/"},
			expected: filepath.Join(currDir, "test_env", "first_1", "second_1", "third_1"),
		},
		{
			name:     "multiple '..' and '.' and dirs",
			args:     []string{"./test_env/first_1/../../test_env/first_1/second_2/third_1/"},
			expected: filepath.Join(currDir, "test_env", "first_1", "second_2", "third_1"),
		},
		{
			name:           "invalid Dir",
			args:           []string{"./asdfasdfastest_env/random/../../test_env/first_1/second_1/third_9/"},
			expected_error: helper.ErrDirDoesnotExist,
			expected:       "",
		},
		{
			name:           "invalid no of args",
			args:           []string{"./asdfasdfastest_env/random/../../test_env/first_1/second_1/third_9/", "asdfasdf"},
			expected_error: helper.ErrInvalidNoOfFlags,
			expected:       "",
		},

		{
			name:     "root Dir",
			args:     []string{"~/Projects"},
			expected: filepath.Join(homeDir, "Projects"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := filedir.Cd{}
			c.CurrDir = currDir
			c.HomeDir = homeDir
			dir, err := c.ProcessCommand(tt.args)
			if err != nil {
				if !errors.Is(err, tt.expected_error) {
					t.Error(err)
				}
			} else if dir != tt.expected {
				t.Errorf("expected: %s\n got:%s \n", tt.expected, dir)
			}
		})
	}
}

func setupDirs() error {
	first_level := []string{"first_1", "first_2", "first_3", "first_4"}
	second_level := []string{"second_1", "second_2", "second_3", "second_4"}
	third_level := []string{"third_1", "third_2", "third_3", "third_4"}
	root_dir := "test_env"
	// Creating first level of dirs
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	err = createDir(homeDir, []string{"test"})
	if err != nil {
		return err
	}
	err = createDir(root_dir, first_level)
	if err != nil {
		return err
	}
	for i := range first_level {
		err = createDir(filepath.Join(root_dir, first_level[i]), second_level)
		if err != nil {
			return err
		}
	}
	for i := range second_level {
		err = createDir(filepath.Join(root_dir, first_level[0], second_level[i]), third_level)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteDirs() error {
	return os.RemoveAll("test_env")
}

func createDir(root string, dirs []string) error {
	for _, dir := range dirs {
		path := filepath.Join(root, dir)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

//
// func removeDir(root string, dirs []string) error {
//
// 	for _, dir := range dirs {
// 		path := filepath.Join(root, dir)
// 		err := os.Remove(path)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
