package runtime

import (
	"encoding/binary"

	"github.com/jamestunnell/slang"
)

type Bytecode struct {
	Constants    []slang.Object
	Instructions []byte
}

const MaxVMConstants = 65535

func NewBytecode() *Bytecode {
	return &Bytecode{
		Constants:    []slang.Object{},
		Instructions: []byte{},
	}
}

func (bc *Bytecode) AddConstant(obj slang.Object) (uint16, bool) {
	if len(bc.Constants) == MaxVMConstants {
		return 0, false
	}

	idx := len(bc.Constants)

	bc.Constants = append(bc.Constants, obj)

	return uint16(idx), true
}

func (bc *Bytecode) AddInstructionNoOperands(opcode Opcode) {
	bc.Instructions = append(bc.Instructions, byte(opcode))
}

func (bc *Bytecode) AddInstructionUint16Operands(opcode Opcode, operands ...uint16) {
	data := make([]byte, 1+2*len(operands))

	data[0] = byte(opcode)

	for i, operand := range operands {
		binary.BigEndian.PutUint16(data[1+2*i:], operand)
	}

	bc.Instructions = append(bc.Instructions, data...)
}
