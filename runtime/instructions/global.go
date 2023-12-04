package instructions

import "github.com/jamestunnell/slang/runtime"

func NewGetGlobal(idx uint16) *runtime.Instruction {
	return runtime.NewInstruction(
		runtime.OpGETGLOBAL, runtime.NewUint16Operand(idx))
}

func NewSetGlobal(idx uint16) *runtime.Instruction {
	return runtime.NewInstruction(
		runtime.OpSETGLOBAL, runtime.NewUint16Operand(idx))
}
