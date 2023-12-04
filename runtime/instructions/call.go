package instructions

import "github.com/jamestunnell/slang/runtime"

func NewCall(numArgs uint8) *runtime.Instruction {
	return runtime.NewInstruction(
		runtime.OpCALL, runtime.NewUint8Operand(numArgs))
}
