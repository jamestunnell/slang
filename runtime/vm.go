package runtime

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/jamestunnell/slang"
)

type VM struct {
	code    *Bytecode
	stack   []slang.Object
	iOffset int
	iLength int
	cLength int
}

var (
	ErrEndOfProgram = errors.New("end of program reached")
)

func NewVM(code *Bytecode) *VM {
	return &VM{
		code:    code,
		stack:   []slang.Object{},
		iOffset: 0,
		iLength: len(code.Instructions),
		cLength: len(code.Constants),
	}
}

func (vm *VM) Step() error {
	if vm.iOffset >= vm.iLength {
		return ErrEndOfProgram
	}

	opcode := Opcode(vm.code.Instructions[vm.iOffset])

	switch opcode {
	case OpConstant:
		idx := binary.BigEndian.Uint16(vm.code.Instructions[vm.iOffset+1:])

		if int(idx) >= vm.cLength {
			return fmt.Errorf("constant index %d is out of bounds", idx)
		}

		vm.push(vm.code.Constants[idx])

		vm.iOffset += 3
	}

	return nil
}

func (vm *VM) Top() (slang.Object, bool) {
	if len(vm.stack) == 0 {
		return nil, false
	}

	return vm.stack[len(vm.stack)-1], true
}

func (vm *VM) push(obj slang.Object) {
	vm.stack = append(vm.stack, obj)
}
