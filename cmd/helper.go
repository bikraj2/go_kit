package main

import "fmt"

func clear() {
	fmt.Print("\033[H\033[2J")
}
