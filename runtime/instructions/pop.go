package instructions

import "github.com/jamestunnell/slang/runtime"

func NewPop() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpPOP)
}
