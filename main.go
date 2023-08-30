package main

import (
	"Goslang/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}

func add(x int, y int) int {
	w := x + y
	return w
}
