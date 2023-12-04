package instructions

import "github.com/jamestunnell/slang/runtime"

func NewNot() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpNOT)
}

func NewNeg() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpNEG)
}
