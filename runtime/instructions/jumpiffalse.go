package instructions

import "github.com/jamestunnell/slang/runtime"

func NewJumpIfFalse(target uint64) (*runtime.Instruction, *runtime.Uint64Operand) {
	targetOperand := runtime.NewUint64Operand(target)

	return runtime.NewInstruction(runtime.OpJUMPIFFALSE, targetOperand), targetOperand
}
