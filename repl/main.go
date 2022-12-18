package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("go-synth wiring REPL")

	Start(os.Stdin, os.Stdout)
}
