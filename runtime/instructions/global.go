package instructions

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/runtime"
)

func NewGetGlobal(idx uint16, sym *slang.Symbol) *runtime.Instruction {
	return runtime.NewInstruction(
		runtime.OpGETGLOBAL, runtime.NewUint16OperandWithSymbol(idx, sym))
}

func NewSetGlobal(idx uint16) *runtime.Instruction {
	return runtime.NewInstruction(
		runtime.OpSETGLOBAL, runtime.NewUint16Operand(idx))
}
