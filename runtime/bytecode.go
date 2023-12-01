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

func (bc *Bytecode) AddConstant(obj slang.Object) bool {
	if len(bc.Constants) == MaxVMConstants {
		return false
	}

	bc.Constants = append(bc.Constants, obj)

	return true
}

func (bc *Bytecode) AddInstructionUint16(opcode Opcode, operand uint16) {
	data := make([]byte, 3)

	data[0] = byte(opcode)

	binary.BigEndian.PutUint16(data[1:], operand)

	bc.Instructions = append(bc.Instructions, data...)
}
