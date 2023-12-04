package instructions

import "github.com/jamestunnell/slang/runtime"

func NewGetFree(idx uint8) *runtime.Instruction {
	return runtime.NewInstruction(
		runtime.OpGETFREE, runtime.NewUint8Operand(idx))
}
