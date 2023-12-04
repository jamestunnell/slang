package runtime

import (
	"github.com/jamestunnell/slang"
)

type Bytecode struct {
	Constants    []slang.Object
	MaxGlobals   int
	Instructions []byte
}

const MaxVMConstants = 65535

func NewBytecode() *Bytecode {
	return &Bytecode{
		Constants:    []slang.Object{},
		MaxGlobals:   0,
		Instructions: []byte{},
	}
}

const DummyJumpTarget = uint64(0xFEEDFEEDFEEDFEED)
