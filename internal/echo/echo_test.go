package echo_test

import (
	"testing"

	helper "go_kit.com/internal"
	"go_kit.com/internal/echo"
)

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
			name: "show help", args: []string{"-help"}, expected: "helloworld\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			output := helper.CaptureStdout(func() {
				e := echo.Echo{}
				e.ProcessCommands(tt.args)
			})
			if output != tt.expected {
				t.Errorf("expected: %s\n got:%s \n", tt.expected, output)
			}
		})
	}
}
