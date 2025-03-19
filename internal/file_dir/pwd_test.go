package filedir_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	helper "go_kit.com/internal"
	filedir "go_kit.com/internal/file_dir"
)

func Test_pwd(t *testing.T) {

	currDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	real_dir, sym_dir := setup(t)
	defer teardown(real_dir)
	tests := []struct {
		name           string
		args           []string
		expected       string
		expected_error error
		setup_dir      string
	}{
		{
			name:      "normal use",
			args:      []string{},
			expected:  currDir + "\n",
			setup_dir: currDir,
		},
		{
			name:           "resolve symlinks",
			args:           []string{"-p"},
			expected_error: helper.ErrFlagCollision,
			expected:       real_dir + "\n",
			setup_dir:      sym_dir,
		},

		{
			name:           "dont resolve symlinks",
			args:           []string{"-l"},
			expected_error: helper.ErrFlagCollision,
			expected:       sym_dir + "\n",
			setup_dir:      sym_dir,
		},
		{
			name:           "invalid no of flags",
			args:           []string{"-l", "-p", "help"},
			expected_error: helper.ErrFlagCollision,
			expected:       "",
			setup_dir:      currDir,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			output := helper.CaptureStdout(func() {
				p := filedir.Pwd{}
				p.CurrDir = tt.setup_dir
				err := p.ProcessCommand(tt.args)
				if err != nil {
					if !errors.Is(err, tt.expected_error) {
						t.Error(err)
					}
				}
			})
			// expectedResolved, _ := filepath.EvalSymlinks(tt.expected)
			// outputResolved, _ := filepath.EvalSymlinks(output)
			if output != tt.expected {
				t.Errorf("expected:%s got:%s", tt.expected, output)
			}
		})
	}
}
func setup(t *testing.T) (string, string) {
	t.Helper()

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "pwd_test")
	if err != nil {
		t.Fatal(err)
	}

	// Create a real directory inside tempDir
	realDir := filepath.Join(tempDir, "real_dir")
	if err := os.Mkdir(realDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create a symlink pointing to realDir
	symlinkPath := filepath.Join(tempDir, "symlink_dir")
	if err := os.Symlink(realDir, symlinkPath); err != nil {
		t.Fatal(err)
	}
	realDir, err = filepath.EvalSymlinks(realDir)
	if err != nil {
		t.Fatal(err)
	}
	return realDir, symlinkPath
}

func teardown(tempDir string) {
	_ = os.RemoveAll(tempDir) // Remove the entire tempDir safely
}
