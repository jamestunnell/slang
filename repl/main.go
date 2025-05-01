package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	// use existing VM check for if connecting to RPC server for an existing VM
	if args.RPCAddr != "" {
		runExisting(args.RPCAddr)
	} else {
		runNew()
	}
}

func runNew() {
	vm := virtualmachine.New(virtualmachine.RandomID(2))

	if err := vm.Start(); err != nil {
		fmt.Printf("REPL: failed to start VM: %v\n", err)

		os.Exit(1)
	}

	for !vm.IsRunning() {
		time.Sleep(25 * time.Millisecond)
	}

	defer func() {
		vm.Stop()
	}()

	runREPL(makeClient(vm.GetRPCAddr()))

	vm.Stop()

	for vm.IsRunning() {
		time.Sleep(25 * time.Millisecond)
	}

	os.Exit(0)
}

func runExisting(rpcAddr string) {
	runREPL(makeClient(rpcAddr))

	os.Exit(0)
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

func runREPL(vm slang.VirtualMachine) {
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
