package instructions

import "github.com/jamestunnell/slang/runtime"

func NewGetLocal(idx uint16) *runtime.Instruction {
	return runtime.NewInstruction(
		runtime.OpGETLOCAL, runtime.NewUint16Operand(idx))
}
