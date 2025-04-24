package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexflint/go-arg"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/jamestunnell/slang/repl/app"
	"github.com/jamestunnell/slang/rpc/client"
	"github.com/jamestunnell/slang/virtualmachine"
)

type Args struct {
	RPCAddr string `help:"RPC server address for a running slang VM" default:""`
}

func main() {
	var args Args

	arg.MustParse(&args)

	logFile, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("Failed to set up debug file logging:", err)

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
		log.Printf("Failed to start VM: %v", err)

		os.Exit(1)
	}

	defer func() {
		vm.Stop()
	}()

	runREPL(&app.Args{
		VMID:    vm.GetID(),
		RPCAddr: vm.GetRPCAddr(),
	})
}

func runExisting(rpcAddr string) {
	// use existing VM check for if connecting to RPC server for an existing VM
	vmInfo, err := client.NewVMInfo(rpcAddr)
	if err != nil {
		fmt.Println("Failed to connect with VM RPC:", err)

		os.Exit(1)
	}

	vmID, err := vmInfo.GetID()
	if err != nil {
		fmt.Println("Failed to get VM ID:", err)

		os.Exit(1)
	}

	runREPL(&app.Args{
		VMID:    vmID,
		RPCAddr: rpcAddr,
	})
}

func runREPL(args *app.Args) {
	p := tea.NewProgram(app.New(args), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Failed to run:", err)
		os.Exit(1)
	}

	os.Exit(0)
}
