package runtime

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/jamestunnell/slang"
)

type VM struct {
	code       *Bytecode
	stack      []slang.Object
	iOffset    int
	iLength    int
	cLength    int
	lastPopped slang.Object
}

var (
	ErrEndOfProgram  = errors.New("end of program reached")
	ErrPopEmptyStack = errors.New("cannot pop an empty stack")
)

func NewVM(code *Bytecode) *VM {
	return &VM{
		code:       code,
		stack:      []slang.Object{},
		iOffset:    0,
		iLength:    len(code.Instructions),
		cLength:    len(code.Constants),
		lastPopped: nil,
	}
}

func (vm *VM) LastPopped() slang.Object {
	return vm.lastPopped
}

func (vm *VM) Step() error {
	if vm.iOffset >= vm.iLength {
		return ErrEndOfProgram
	}

	opcode := Opcode(vm.code.Instructions[vm.iOffset])

	var err error

	switch opcode {
	case OpCONST:
		idx := binary.BigEndian.Uint16(vm.code.Instructions[vm.iOffset+1:])

		if int(idx) >= vm.cLength {
			err = fmt.Errorf("constant index %d is out of bounds", idx)

			break
		}

		vm.push(vm.code.Constants[idx])

		vm.iOffset += 3
	case OpPOP:
		vm.pop()

		vm.iOffset++
	case OpADD:
		err = vm.exeBinaryOp(slang.MethodADD)
	case OpSUB:
		err = vm.exeBinaryOp(slang.MethodSUB)
	case OpMUL:
		err = vm.exeBinaryOp(slang.MethodMUL)
	case OpDIV:
		err = vm.exeBinaryOp(slang.MethodDIV)
	case OpEQ:
		err = vm.exeBinaryOp(slang.MethodEQ)
	case OpNEQ:
		err = vm.exeBinaryOp(slang.MethodNEQ)
	case OpLT:
		err = vm.exeBinaryOp(slang.MethodLT)
	case OpLEQ:
		err = vm.exeBinaryOp(slang.MethodLEQ)
	case OpGT:
		err = vm.exeBinaryOp(slang.MethodGT)
	case OpGEQ:
		err = vm.exeBinaryOp(slang.MethodGEQ)
	case OpAND:
		err = vm.exeBinaryOp(slang.MethodAND)
	case OpOR:
		err = vm.exeBinaryOp(slang.MethodOR)
	default:
		err = fmt.Errorf("unknown opcode %d", opcode)
	}

	return err
}

func (vm *VM) StackSize() int {
	return len(vm.stack)
}

func (vm *VM) Top() (slang.Object, bool) {
	if len(vm.stack) == 0 {
		return nil, false
	}

	return vm.stack[len(vm.stack)-1], true
}

func (vm *VM) exeBinaryOp(method string) error {
	right, ok := vm.pop()
	if !ok {
		return fmt.Errorf("failed to get right operand: %w", ErrPopEmptyStack)
	}

	left, ok := vm.pop()
	if !ok {
		return fmt.Errorf("failed to get left operand: %w", ErrPopEmptyStack)
	}

	result, err := left.Send(method, right)
	if err != nil {
		return fmt.Errorf("method %s failed: %w", method, err)
	}

	vm.push(result)

	vm.iOffset++

	return nil
}

func (vm *VM) push(obj slang.Object) {
	vm.stack = append(vm.stack, obj)
}

func (vm *VM) pop() (slang.Object, bool) {
	size := len(vm.stack)

	if size == 0 {
		return nil, false
	}

	obj := vm.stack[size-1]

	vm.stack = vm.stack[:size-1]

	vm.lastPopped = obj

	return obj, true
}
