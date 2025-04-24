package virtualmachine

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"path"
	"sync/atomic"

	"github.com/jamestunnell/slang/rpc/server"
)

type VM struct {
	id       string
	archives *server.Archives
	vmInfo   *server.VMInfo

	rpcServer              *rpc.Server
	rpcAddr                atomic.Value
	running, stopRequested atomic.Bool
}

var errAlreadyRunning = errors.New("VM is already running")

func New(id string) *VM {
	archives := server.NewArchives()
	vmInfo := server.NewVMInfo(id)

	rpcPath := path.Join("/slang", id, "rpc")
	rpcServer := rpc.NewServer()

	rpcServer.Register(archives)
	rpcServer.HandleHTTP(rpcPath, path.Join(rpcPath, "debug"))

	var rpcAddr atomic.Value

	rpcAddr.Store("")

	return &VM{
		id:        id,
		archives:  archives,
		vmInfo:    vmInfo,
		rpcServer: rpcServer,
		rpcAddr:   rpcAddr,
	}
}

func (vm *VM) GetID() string {
	return vm.id
}

func (vm *VM) GetRPCAddr() string {
	return vm.rpcAddr.Load().(string)
}

func (vm *VM) Start() error {
	if vm.running.Load() {
		return errAlreadyRunning
	}

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return fmt.Errorf("failed to start RPC net listener: %w", err)
	}

	log.Printf("VM: starting RPC server on %s", listener.Addr())

	go vm.rpcServer.Accept(listener)

	go vm.run(listener)

	vm.rpcAddr.Store(listener.Addr().String())

	return nil
}

func (vm *VM) Stop() {
	vm.stopRequested.Store(true)
}

func (vm *VM) run(listener net.Listener) {
	vm.running.Store(true)

	log.Printf("VM: running")

	for !vm.stopRequested.Load() {
	}

	vm.stopRequested.Store(false)

	if err := listener.Close(); err != nil {
		log.Printf("failed to close RPC net listener: %v", err)
	}

	vm.running.Store(false)
	vm.rpcAddr.Store("")

	log.Printf("VM: stopped")
}
