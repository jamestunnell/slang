package instructions

import "github.com/jamestunnell/slang/runtime"

func NewReturn() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpRETURN)
}

func NewReturnVal() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpRETURNVAL)
}
