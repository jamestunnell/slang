package instructions

import "github.com/jamestunnell/slang/runtime"

func NewGetFree(idx uint16) *runtime.Instruction {
	return runtime.NewInstruction(
		runtime.OpGETFREE, runtime.NewUint16Operand(idx))
}
