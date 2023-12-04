package runtime

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/runtime/objects"
)

type VM struct {
	code       *Bytecode
	stack      []slang.Object
	globals    []slang.Object
	iOffset    uint64
	iLength    uint64
	cLength    int
	lastPopped slang.Object
}

const MaxGlobals = 65536

var (
	ErrEndOfProgram  = errors.New("end of program reached")
	ErrPopEmptyStack = errors.New("cannot pop an empty stack")

	False = objects.NewBool(false)
)

func NewVM(code *Bytecode) *VM {
	return &VM{
		code:       code,
		stack:      []slang.Object{},
		globals:    make([]slang.Object, MaxGlobals),
		iOffset:    0,
		iLength:    uint64(len(code.Instructions)),
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
		err = vm.exeConst()
	case OpGETGLOBAL:
		err = vm.exeGetGlobal()
	case OpSETGLOBAL:
		err = vm.exeSetGlobal()
	case OpJUMP:
		vm.iOffset = binary.BigEndian.Uint64(vm.code.Instructions[vm.iOffset+1:])
	case OpJUMPIFFALSE:
		err = vm.exeJumpIfFalse()
	case OpPOP:
		vm.pop()

		vm.iOffset++
	case OpNEG:
		err = vm.exeUnaryOp(slang.MethodNEG)
	case OpNOT:
		err = vm.exeUnaryOp(slang.MethodNOT)
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

func (vm *VM) exeConst() error {
	idx := binary.BigEndian.Uint16(vm.code.Instructions[vm.iOffset+1:])

	if int(idx) >= vm.cLength {
		return fmt.Errorf("constant index %d is out of bounds", idx)
	}

	vm.push(vm.code.Constants[idx])

	vm.iOffset += 3

	return nil
}

func (vm *VM) exeGetGlobal() error {
	idx := binary.BigEndian.Uint16(vm.code.Instructions[vm.iOffset+1:])

	vm.push(vm.globals[idx])

	vm.iOffset += 3

	return nil
}

func (vm *VM) exeSetGlobal() error {
	idx := binary.BigEndian.Uint16(vm.code.Instructions[vm.iOffset+1:])

	obj, ok := vm.pop()
	if !ok {
		return fmt.Errorf("failed to get set global value: %w", ErrPopEmptyStack)
	}

	vm.globals[idx] = obj

	vm.iOffset += 3

	return nil
}

func (vm *VM) exeJumpIfFalse() error {
	val, ok := vm.pop()
	if !ok {
		return fmt.Errorf("failed to get jump-if-false value: %w", ErrPopEmptyStack)
	}

	if val.Equal(False) {
		vm.iOffset = binary.BigEndian.Uint64(vm.code.Instructions[vm.iOffset+1:])
	} else {
		vm.iOffset += 9
	}

	return nil
}

func (vm *VM) exeUnaryOp(method string) error {
	val, ok := vm.pop()
	if !ok {
		return fmt.Errorf("failed to get unary operand: %w", ErrPopEmptyStack)
	}

	result, err := val.Send(method)
	if err != nil {
		return fmt.Errorf("method %s failed: %w", method, err)
	}

	vm.push(result)

	vm.iOffset++

	return nil
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
