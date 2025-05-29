package virtualmachine

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/rpc"
	"slices"
	"sync"
	"sync/atomic"

	"golang.org/x/exp/maps"

	"github.com/google/uuid"
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/archives"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/rpc/server"
)

type VM struct {
	id        uuid.UUID
	name      string
	rpcServer *rpc.Server
	rpcAddr   atomic.Value
	running   atomic.Bool
	stop      chan struct{}

	workersMut sync.RWMutex
	workers    map[string]*Worker
}

var (
	errAlreadyRunning = errors.New("VM is already running")
	errNotImplemented = errors.New("not implemented")
)

const channelDepth = 10

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
		stop:      make(chan struct{}),
		workers:   map[string]*Worker{},
	}

	gob.Register(fmt.Errorf("%w", errors.New("")))
	gob.Register(&expressions.Identifier{})
	gob.Register(&expressions.Const[float64]{})
	gob.Register(&expressions.Const[int64]{})
	gob.Register(&expressions.Const[string]{})

	gob.Register(&archives.TarGz{})

	rpcServer.Register(&server.Packages{VM: vm})
	rpcServer.Register(&server.Info{VM: vm})

	return vm
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

	log.Printf("VM: starting RPC server on %s\n", listener.Addr())

	go vm.rpcServer.Accept(listener)

	vm.rpcAddr.Store(listener.Addr().String())

	go vm.runUntilStopped(listener)

	return nil
}

func (vm *VM) Stop() {
	vm.stop <- struct{}{}
}

func (vm *VM) GetName() string {
	return vm.name
}

func (vm *VM) GetID() uuid.UUID {
	return vm.id
}

func (vm *VM) IsRunning() bool {
	return vm.running.Load()
}

func (vm *VM) UpsertPackage(meta slang.PackageMeta, archive slang.PackageArchive) {
	vm.workersMut.Lock()

	defer vm.workersMut.Unlock()

	key := meta.Address.String()

	if worker, found := vm.workers[key]; found {
		worker.Stop()
	}

	worker := NewWorker(meta, archive)

	vm.workers[key] = worker

	worker.Start()
}

func (vm *VM) RemovePackage(addr slang.PackageAddress) bool {
	vm.workersMut.Lock()

	defer vm.workersMut.Unlock()

	key := addr.String()

	if _, found := vm.workers[key]; !found {
		return false
	}

	delete(vm.workers, key)

	return true
}

func (vm *VM) ListPackages() []slang.PackageAddress {
	vm.workersMut.RLock()

	defer vm.workersMut.RUnlock()

	keys := maps.Keys(vm.workers)

	slices.Sort(keys)

	addrs := make([]slang.PackageAddress, len(keys))

	for i, key := range keys {
		addrs[i] = slang.PackageAddress{}

		(&addrs[i]).Parse(key)
	}

	return addrs
}

func (vm *VM) GetPackageState(addr slang.PackageAddress) (slang.PackageState, bool) {
	vm.workersMut.RLock()

	defer vm.workersMut.RUnlock()

	worker, found := vm.workers[addr.String()]
	if !found {
		return 0, false
	}

	return worker.GetState(), true
}

func (vm *VM) GetPackageArchive(addr slang.PackageAddress) (slang.PackageArchive, bool) {
	vm.workersMut.RLock()

	defer vm.workersMut.RUnlock()

	worker, found := vm.workers[addr.String()]
	if !found {
		return nil, false
	}

	return worker.GetArchive(), true
}

func (vm *VM) GetPackageFiles(addr slang.PackageAddress) (fs.FS, bool) {
	vm.workersMut.RLock()

	defer vm.workersMut.RUnlock()

	worker, found := vm.workers[addr.String()]
	if !found {
		return nil, false
	}

	if worker.GetState() < slang.PkgUnpacked {
		return nil, false
	}

	return worker.GetFiles(), true
}

func (vm *VM) GetPackageAST(addr slang.PackageAddress) (slang.PackageAST, bool) {
	vm.workersMut.RLock()

	defer vm.workersMut.RUnlock()

	worker, found := vm.workers[addr.String()]
	if !found {
		return nil, false
	}

	if worker.GetState() < slang.PkgCompiled {
		return nil, false
	}

	return worker.GetAST(), true
}

func (vm *VM) GetPackageDependencies(addr slang.PackageAddress) ([]slang.PackageAddress, bool) {
	vm.workersMut.RLock()

	defer vm.workersMut.RUnlock()

	worker, found := vm.workers[addr.String()]
	if !found {
		return []slang.PackageAddress{}, false
	}

	if worker.GetState() < slang.PkgResolved {
		return []slang.PackageAddress{}, false
	}

	return worker.GetDependencies(), true
}

func (vm *VM) GetPackageAnalysis(addr slang.PackageAddress) (slang.PackageAnalysis, bool) {
	vm.workersMut.RLock()

	defer vm.workersMut.RUnlock()

	worker, found := vm.workers[addr.String()]
	if !found {
		return slang.PackageAnalysis{}, false
	}

	if worker.GetState() < slang.PkgAnalyzed {
		return slang.PackageAnalysis{}, false
	}

	return worker.GetAnalysis(), true
}

func (vm *VM) GetPackageBytecode(addr slang.PackageAddress) (slang.PackageBytecode, bool) {
	vm.workersMut.RLock()

	defer vm.workersMut.RUnlock()

	worker, found := vm.workers[addr.String()]
	if !found {
		return slang.PackageBytecode{}, false
	}

	if worker.GetState() < slang.PkgCompiled {
		return slang.PackageBytecode{}, false
	}

	return worker.GetBytecode(), true
}

func (vm *VM) GetPackageFailure(addr slang.PackageAddress) (slang.PackageFailure, bool) {
	vm.workersMut.RLock()

	defer vm.workersMut.RUnlock()

	worker, found := vm.workers[addr.String()]
	if !found {
		return slang.PackageFailure{}, false
	}

	if worker.GetState() != slang.PkgFailed {
		return slang.PackageFailure{}, false
	}

	return worker.GetFailure(), true
}

func (vm *VM) runUntilStopped(listener net.Listener) {
	vm.running.Store(true)

	log.Println("VM: running")

	<-vm.stop

	log.Println("VM: closing listener")

	if err := listener.Close(); err != nil {
		log.Printf("failed to close RPC net listener: %v\n", err)
	}

	vm.running.Store(false)
	vm.rpcAddr.Store("")

	log.Println("VM: stopped")
}
