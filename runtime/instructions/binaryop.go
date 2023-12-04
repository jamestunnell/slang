package instructions

import "github.com/jamestunnell/slang/runtime"

func NewAdd() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpADD)
}

func NewSub() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpSUB)
}

func NewMul() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpMUL)
}

func NewDiv() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpDIV)
}

func NewEqual() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpEQ)
}

func NewNotEqual() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpNEQ)
}

func NewLess() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpLT)
}

func NewLessEqual() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpLEQ)
}

func NewGreater() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpGT)
}

func NewGreaterEqual() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpGEQ)
}

func NewAnd() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpAND)
}

func NewOr() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpOR)
}
