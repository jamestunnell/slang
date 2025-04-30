package virtualmachine

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
	"sync/atomic"

	"golang.org/x/exp/maps"

	"github.com/google/uuid"
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/archives"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
	"github.com/jamestunnell/slang/rpc/server"
)

type VM struct {
	id        uuid.UUID
	name      string
	rpcServer *rpc.Server
	rpcAddr   atomic.Value
	running   atomic.Bool
	stop      chan struct{}

	packageMux      sync.RWMutex
	packageArchives map[string]slang.PackageArchive
	packageASTs     map[string]slang.Package
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
		id:              id,
		name:            name,
		rpcServer:       rpcServer,
		rpcAddr:         rpcAddr,
		stop:            make(chan struct{}),
		packageArchives: map[string]slang.PackageArchive{},
		packageASTs:     map[string]slang.Package{},
	}

	gob.Register(fmt.Errorf("%w", errors.New("")))
	gob.Register(&expressions.Identifier{})
	gob.Register(&expressions.Float{})
	gob.Register(&expressions.Int{})

	gob.Register(&archives.TarGz{})

	rpcServer.Register(&server.Packages{VM: vm})
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

	log.Printf("VM: starting RPC server on %s\n", listener.Addr())

	go vm.rpcServer.Accept(listener)

	vm.rpcAddr.Store(listener.Addr().String())

	go vm.runUntilStopped(listener)

	return nil
}

func (vm *VM) Stop() {
	vm.stop <- struct{}{}
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

func (vm *VM) EvaluateExpr(expr slang.Expression) (slang.Object, error) {
	return nil, errNotImplemented
}

func (vm *VM) ListPackages() []slang.PackageMeta {
	vm.packageMux.RLock()

	defer vm.packageMux.RUnlock()

	metas := make([]slang.PackageMeta, len(vm.packageArchives))

	for i, archive := range maps.Values(vm.packageArchives) {
		metas[i] = archive.GetMeta()
	}

	return metas
}

func (vm *VM) GetPackageArchive(meta slang.PackageMeta) (slang.PackageArchive, bool) {
	vm.packageMux.RLock()

	defer vm.packageMux.RUnlock()

	a, found := vm.packageArchives[meta.String()]

	return a, found
}

func (vm *VM) AddPackage(archive slang.PackageArchive) error {
	vm.packageMux.Lock()

	defer vm.packageMux.Unlock()

	key := packageKey(archive.GetMeta())

	archiveFS, err := archive.Unpack()
	if err != nil {
		return fmt.Errorf("failed to unpack archive: %w", err)
	}

	modules, err := parsing.ParsePackage(archiveFS, parsers.NewFileParser())
	if err != nil {
		return fmt.Errorf("failed to parse package: %w", err)
	}

	vm.packageArchives[key] = archive
	vm.packageASTs[key] = ast.NewPackage(archive.GetMeta(), modules...)

	return nil
}

func (vm *VM) RemovePackage(meta slang.PackageMeta) bool {
	vm.packageMux.Lock()

	defer vm.packageMux.Unlock()

	key := packageKey(meta)
	if _, found := vm.packageArchives[key]; !found {
		return false
	}

	delete(vm.packageArchives, key)
	delete(vm.packageASTs, key)

	return true
}

func packageKey(meta slang.PackageMeta) string {
	return meta.Address.String()
}
