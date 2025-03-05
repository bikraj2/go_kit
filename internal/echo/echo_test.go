package echo_test

import (
	"io"
	"os"
	"testing"

	"go_kit.com/internal/echo"
)

func captureStdout(f func()) string {
	old := os.Stdout     // Save current stdout
	r, w, _ := os.Pipe() // Create pipe
	os.Stdout = w

	f() // Run the function that prints to stdout

	// Restore stdout and read from the pipe
	w.Close()
	os.Stdout = old
	out, _ := io.ReadAll(r)
	return string(out)
}

func Test_echo(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name: "basic ", args: []string{"helloworld"}, expected: "helloworld\n",
		},
		{
			name: "parse escape sequences", args: []string{"-e", "helloworld\n"}, expected: "helloworld\n\n",
		},
		{
			name: "dosnot parse escape sequences", args: []string{"helloworld\n"}, expected: "helloworld\\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			output := captureStdout(func() {
				e := echo.Echo{}
				e.ProcessCommands(tt.args)
			})
			if output != tt.expected {
				t.Errorf("expected: %s\n got:%s \n", tt.expected, output)
			}
		})
	}
}
