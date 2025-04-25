package virtualmachine

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/rpc/server"
)

type VM struct {
	id                     uuid.UUID
	name                   string
	rpcServer              *rpc.Server
	rpcAddr                atomic.Value
	running, stopRequested atomic.Bool

	archives map[string]slang.PackageArchive
}

var (
	errAlreadyRunning = errors.New("VM is already running")
	errNotImplemented = errors.New("not implemented")
)

func New(name string) *VM {
	id := uuid.New()
	rpcServer := rpc.NewServer()

	var rpcAddr atomic.Value

	rpcAddr.Store("")

	vm := &VM{
		id:        id,
		name:      name,
		rpcServer: rpcServer,
		rpcAddr:   rpcAddr,
	}

	gob.Register(fmt.Errorf("%w", errors.New("")))
	gob.Register(&expressions.Identifier{})
	gob.Register(&expressions.Float{})
	gob.Register(&expressions.Int{})

	rpcServer.Register(&server.Archives{VM: vm})
	rpcServer.Register(&server.Expressions{VM: vm})
	rpcServer.Register(&server.VMInfo{VM: vm})

	return vm
}

func (vm *VM) GetInfo() slang.VMInfo {
	return slang.VMInfo{
		Name: vm.name,
		ID:   vm.id,
	}
}

func (vm *VM) IsRunning() bool {
	return vm.running.Load()
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

	vm.rpcAddr.Store(listener.Addr().String())

	go vm.run(listener)

	return nil
}

func (vm *VM) Stop() {
	vm.stopRequested.Store(true)
}

func (vm *VM) run(listener net.Listener) {
	vm.running.Store(true)

	log.Println("VM: running")

	for !vm.stopRequested.Load() {
		time.Sleep(100 * time.Millisecond)
	}

	log.Println("VM: closing listener")

	if err := listener.Close(); err != nil {
		log.Printf("failed to close RPC net listener: %v", err)
	}

	vm.running.Store(false)
	vm.rpcAddr.Store("")
	vm.stopRequested.Store(false)

	log.Println("VM: stopped")
}

func (vm *VM) EvaluateExpr(expr slang.Expression) (slang.Object, error) {
	return nil, errNotImplemented
}

func (vm *VM) ListPackages() []slang.PackageMeta {
	metas := make([]slang.PackageMeta, len(vm.archives))
	i := 0

	for _, archive := range vm.archives {
		metas[i] = archive.GetMeta()

		i++
	}

	return metas
}

func (vm *VM) GetPackage(meta slang.PackageMeta) (slang.PackageArchive, bool) {
	a, found := vm.archives[meta.String()]

	return a, found
}

func (vm *VM) AddPackage(a slang.PackageArchive) error {
	vm.archives[a.GetMeta().String()] = a

	return nil
}

func (vm *VM) RemovePackage(meta slang.PackageMeta) bool {
	key := meta.String()

	if _, found := vm.archives[key]; !found {
		return false
	}

	delete(vm.archives, key)

	return true
}
