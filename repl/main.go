package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexflint/go-arg"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/repl/app"
	"github.com/jamestunnell/slang/virtualmachine"
)

type Args struct {
	RPCAddr string `help:"RPC server address for a running slang VM" default:""`
}

func main() {
	var args Args

	arg.MustParse(&args)

	var vm slang.VirtualMachine

	// use existing VM check for if connecting to RPC server for an existing VM
	if args.RPCAddr != "" {
		vm = makeClient(args.RPCAddr)

		log.Printf("REPL: connected to existing VM %s\n", vm.GetName())
	} else {
		vm = virtualmachine.New(virtualmachine.RandomID(2))

		log.Printf("REPL: created new VM %s\n", vm.GetName())
	}

	logFile, err := tea.LogToFile(vm.GetName()+".log", "debug")
	if err != nil {
		fmt.Printf("REPL: failed to set up debug file logging: %v\n", err)

		os.Exit(1)
	}

	defer logFile.Close()

	log.Printf("REPL: starting app (VM name=%s)\n", vm.GetName())

	app := app.New(vm)
	p := tea.NewProgram(app, tea.WithAltScreen(), tea.WithInputTTY())

	if _, err := p.Run(); err != nil {
		fmt.Printf("REPL: failed to run: %v\n", err)

		os.Exit(1)
	}

	log.Println("REPL: app stopped")
}

func makeClient(tcpAddr string) slang.VirtualMachine {
	log.Printf("REPL: making client for VM RPC at %s\n", tcpAddr)

	c, err := virtualmachine.MakeClient(tcpAddr)
	if err != nil {
		fmt.Printf("REPL: failed to make VM client: %v\n", err)

		os.Exit(1)
	}

	return c
}
