package filedir

import (
	"fmt"

	"go_kit.com/internal/color"
)

type Ls struct {
}

func (l *Ls) processCommand(args []string) {
	if len(args) == 1 {
		dirs, err := list_file(".")
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, dir := range dirs {
			if dir.Type().IsDir() {
				fmt.Printf("%v%v\n", color.Colors["blue"], dir.Name())
			} else {
				fmt.Printf("%v\n", dir.Name())
			}
		}
	}
}
