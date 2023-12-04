package instructions

import "github.com/jamestunnell/slang/runtime"

func NewSetLocal(idx uint16) *runtime.Instruction {
	return runtime.NewInstruction(
		runtime.OpSETLOCAL, runtime.NewUint16Operand(idx))
}
