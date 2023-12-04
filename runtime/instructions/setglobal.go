package instructions

import "github.com/jamestunnell/slang/runtime"

func NewSetGlobal(idx uint16) *runtime.Instruction {
	return runtime.NewInstruction(
		runtime.OpSETGLOBAL, runtime.NewUint16Operand(idx))
}
