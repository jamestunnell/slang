package instructions

import "github.com/jamestunnell/slang/runtime"

func NewReturnVal() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpRETURNVAL)
}
