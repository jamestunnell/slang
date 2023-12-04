package instructions

import "github.com/jamestunnell/slang/runtime"

func NewGetLocal(idx uint8) *runtime.Instruction {
	return runtime.NewInstruction(
		runtime.OpGETLOCAL, runtime.NewUint8Operand(idx))
}

func NewSetLocal(idx uint8) *runtime.Instruction {
	return runtime.NewInstruction(
		runtime.OpSETLOCAL, runtime.NewUint8Operand(idx))
}
