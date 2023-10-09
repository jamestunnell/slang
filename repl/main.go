package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("slang REPL (CTRL-C to exit)")

	Start(os.Stdin, os.Stdout)
}
