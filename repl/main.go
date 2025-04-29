package main

import (
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

	os.Remove("debug.log")

	logFile, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Printf("REPL: failed to set up debug file logging: %v\n", err)

		os.Exit(1)
	}
	defer logFile.Close()

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
		log.Printf("REPL: failed to start VM: %v\n", err)

		os.Exit(1)
	}

	for !vm.IsRunning() {
		log.Println("REPL: waiting for VM to run")

		time.Sleep(25 * time.Millisecond)
	}

	defer func() {
		vm.Stop()
	}()

	c := makeClient(vm.GetRPCAddr())

	runREPL(c, vm.GetInfo())

	vm.Stop()

	for vm.IsRunning() {
		log.Println("REPL: waiting for VM to stop")

		time.Sleep(25 * time.Millisecond)
	}

	os.Exit(0)
}

func runExisting(rpcAddr string) {
	c := makeClient(rpcAddr)

	// use existing VM check for if connecting to RPC server for an existing VM
	vmInfo, err := c.GetInfo()
	if err != nil {
		log.Printf("REPL: failed to get VM info: %v\n", err)

		os.Exit(1)
	}

	runREPL(c, vmInfo)

	os.Exit(0)
}

func makeClient(tcpAddr string) virtualmachine.Client {
	log.Printf("REPL: making client for VM RPC at %s\n", tcpAddr)

	c, err := virtualmachine.MakeClient(tcpAddr)
	if err != nil {
		log.Printf("REPL: failed to make VM client: %v\n", err)

		os.Exit(1)
	}

	return c
}

func runREPL(
	c virtualmachine.Client,
	info slang.VMInfo,
) {
	log.Printf("REPL: starting app (VM name=%s)\n", info.Name)

	app := app.New(c, info)
	p := tea.NewProgram(app, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Printf("REPL: failed to run: %v\n", err)

		os.Exit(1)
	}

	log.Println("REPL: app stopped")
}
