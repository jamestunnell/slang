package virtualmachine

import (
	"fmt"
	"io/fs"
	"log"
	"sync"
	"time"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
)

type Worker struct {
	mut  sync.RWMutex
	stop chan struct{}

	meta         slang.PackageMeta
	state        slang.PackageState
	archive      slang.PackageArchive
	files        fs.FS
	ast          slang.PackageAST
	dependencies []slang.PackageAddress
	analysis     slang.PackageAnalysis
	bytecode     slang.PackageBytecode
	failure      slang.PackageFailure
}

const notImplemented = "action not implemented"

func NewWorker(
	meta slang.PackageMeta,
	archive slang.PackageArchive,
) *Worker {
	return &Worker{
		stop:         make(chan struct{}),
		archive:      archive,
		meta:         meta,
		state:        slang.PkgAdded,
		dependencies: []slang.PackageAddress{},
	}
}

func (w *Worker) Start() {
	go w.runUntilStopped()
}

func (w *Worker) Stop() {
	w.stop <- struct{}{}
}

func (w *Worker) GetState() slang.PackageState {
	w.mut.RLock()

	defer w.mut.RUnlock()

	return w.state
}

func (w *Worker) GetArchive() slang.PackageArchive {
	w.mut.RLock()

	defer w.mut.RUnlock()

	return w.archive
}

func (w *Worker) GetFiles() fs.FS {
	w.mut.RLock()

	defer w.mut.RUnlock()

	return w.files
}

func (w *Worker) GetAST() slang.PackageAST {
	w.mut.RLock()

	defer w.mut.RUnlock()

	return w.ast
}

func (w *Worker) GetDependencies() []slang.PackageAddress {
	w.mut.RLock()

	defer w.mut.RUnlock()

	return w.dependencies
}

func (w *Worker) GetAnalysis() slang.PackageAnalysis {
	w.mut.RLock()

	defer w.mut.RUnlock()

	return w.analysis
}

func (w *Worker) GetBytecode() slang.PackageBytecode {
	w.mut.RLock()

	defer w.mut.RUnlock()

	return w.bytecode
}

func (w *Worker) GetFailure() slang.PackageFailure {
	w.mut.RLock()

	defer w.mut.RUnlock()

	return w.failure
}

func (w *Worker) runUntilStopped() {
	for {
		select {
		case <-w.stop:
			return
		case <-time.After(10 * time.Millisecond):
			w.run()
		}
	}
}

func (w *Worker) run() {
	switch w.GetState() {
	case slang.PkgAdded:
		w.unpack()
	case slang.PkgUnpacked:
		w.parse()
	case slang.PkgParsed:
		w.resolve()
	case slang.PkgResolved:
		w.analyze()
	case slang.PkgAnalyzed:
		w.compile()
	case slang.PkgFailed:
	}
}

func (w *Worker) actionFailed(action slang.PackageAction, errMsg string) {
	w.mut.Lock()

	defer w.mut.Unlock()

	w.failure = slang.PackageFailure{
		FailedAction: action,
		ErrorMsg:     errMsg,
	}
	w.state = slang.PkgFailed

	log.Printf("%s: %s failed (%s)\n", w.meta.Address, action, errMsg)
}

func (w *Worker) actionSucceeded(
	action slang.PackageAction,
	nextState slang.PackageState,
	effect func(),
) {
	w.mut.Lock()

	defer w.mut.Unlock()

	w.state = nextState

	effect()

	log.Printf("%s: %s succeeded\n", w.meta.Address, action)
}

func (w *Worker) unpack() {
	files, err := w.archive.Unpack()
	if err != nil {
		w.actionFailed(slang.PkgUnpack, fmt.Sprintf("failed to unpack archive: %v", err))

		return
	}

	w.actionSucceeded(slang.PkgUnpack, slang.PkgUnpacked, func() { w.files = files })
}

func (w *Worker) parse() {
	modules, err := parsing.ParsePackage(w.GetFiles(), parsers.NewFileParser())
	if err != nil {
		w.actionFailed(slang.PkgParse, fmt.Sprintf("failed to unpack archive: %v", err))

		return
	}

	w.actionSucceeded(slang.PkgParse, slang.PkgParsed, func() { w.ast = ast.NewPackage(w.meta, modules...) })
}

func (w *Worker) resolve() {
	w.actionFailed(slang.PkgResolve, notImplemented)
}

func (w *Worker) analyze() {
	w.actionFailed(slang.PkgAnalyze, notImplemented)
}

func (w *Worker) compile() {
	w.actionFailed(slang.PkgCompile, notImplemented)
}
