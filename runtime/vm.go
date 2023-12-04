package runtime

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/runtime/objects"
)

type VM struct {
	frames          []*Frame
	frameCount      int
	lastPoppedFrame *Frame

	stack      []slang.Object
	stackCount int
	lastPopped slang.Object

	constants []slang.Object
	numConsts int

	globals []slang.Object
}

const (
	GlobalsSize = 65536
	MaxFrames   = 256
	StackSize   = 2048
)

var (
	ErrEndOfProgram       = errors.New("end of program reached")
	ErrPopEmptyStack      = errors.New("cannot pop an empty stack")
	ErrStackOverflow      = errors.New("stack overflow")
	ErrPopEmptyFrameStack = errors.New("cannot pop empty frame stack")

	False = objects.NewBool(false)
)

func NewVM(code *Bytecode) *VM {
	compiledFn := objects.NewCompiledFunc(code.Instructions, 0)
	mainClosure := objects.NewClosure(compiledFn)
	mainFrame := NewFrame(mainClosure, 0)
	return &VM{
		frames:     []*Frame{mainFrame},
		frameCount: 1,
		stack:      make([]slang.Object, StackSize),
		stackCount: 0,
		constants:  code.Constants,
		numConsts:  len(code.Constants),
		globals:    make([]slang.Object, GlobalsSize),
		lastPopped: nil,
	}
}

func (vm *VM) LastPopped() slang.Object {
	return vm.lastPopped
}

func (vm *VM) Step() error {
	f := vm.currentFrame()
	if f.InstrOffset >= f.InstrLength {
		return ErrEndOfProgram
	}

	instr := f.Instructions()

	opcode := Opcode(instr[f.InstrOffset])

	var err error

	switch opcode {
	case OpGETCONST:
		err = vm.exeGetConst(f)
	case OpGETGLOBAL:
		err = vm.exeGetGlobal(f)
	case OpGETLOCAL:
		err = vm.exeGetLocal(f)
	case OpGETFREE:
		err = vm.exeGetFree(f)
	case OpSETGLOBAL:
		err = vm.exeSetGlobal(f)
	case OpSETLOCAL:
		err = vm.exeSetLocal(f)
	case OpCLOSURE:
		err = vm.exeClosure(f)
	case OpJUMP:
		f.InstrOffset = binary.BigEndian.Uint64(f.Instructions()[f.InstrOffset+1:])
	case OpJUMPIFFALSE:
		err = vm.exeJumpIfFalse(f)
	case OpPOP:
		vm.pop()

		f.InstrOffset++
	case OpCALL:
		err = vm.exeCall(f)
	case OpRETURN:
		err = vm.exeReturn(f)
	case OpRETURNVAL:
		err = vm.exeReturnValue(f)
	case OpNEG:
		err = vm.exeUnaryOp(f, slang.MethodNEG)
	case OpNOT:
		err = vm.exeUnaryOp(f, slang.MethodNOT)
	case OpADD:
		err = vm.exeBinaryOp(f, slang.MethodADD)
	case OpSUB:
		err = vm.exeBinaryOp(f, slang.MethodSUB)
	case OpMUL:
		err = vm.exeBinaryOp(f, slang.MethodMUL)
	case OpDIV:
		err = vm.exeBinaryOp(f, slang.MethodDIV)
	case OpEQ:
		err = vm.exeBinaryOp(f, slang.MethodEQ)
	case OpNEQ:
		err = vm.exeBinaryOp(f, slang.MethodNEQ)
	case OpLT:
		err = vm.exeBinaryOp(f, slang.MethodLT)
	case OpLEQ:
		err = vm.exeBinaryOp(f, slang.MethodLEQ)
	case OpGT:
		err = vm.exeBinaryOp(f, slang.MethodGT)
	case OpGEQ:
		err = vm.exeBinaryOp(f, slang.MethodGEQ)
	case OpAND:
		err = vm.exeBinaryOp(f, slang.MethodAND)
	case OpOR:
		err = vm.exeBinaryOp(f, slang.MethodOR)
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

func (vm *VM) currentFrame() *Frame {
	return vm.frames[vm.frameCount-1]
}

func (vm *VM) exeGetConst(f *Frame) error {
	idx := binary.BigEndian.Uint16(f.Instructions()[f.InstrOffset+1:])

	if int(idx) >= vm.numConsts {
		return fmt.Errorf("constant index %d is out of bounds", idx)
	}

	f.InstrOffset += 3

	return vm.push(vm.constants[idx])
}

func (vm *VM) exeGetGlobal(f *Frame) error {
	idx := binary.BigEndian.Uint16(f.Instructions()[f.InstrOffset+1:])

	f.InstrOffset += 3

	return vm.push(vm.globals[idx])
}

func (vm *VM) exeGetLocal(f *Frame) error {
	localIdx := f.Instructions()[f.InstrOffset+1]

	f.InstrOffset += 2

	return vm.push(vm.stack[f.BaseStackCount+int(localIdx)])
}

func (vm *VM) exeGetFree(f *Frame) error {
	freeIdx := f.Instructions()[f.InstrOffset+1]

	f.InstrOffset += 2

	return vm.push(f.Closure.FreeVars[freeIdx])
}

func (vm *VM) exeSetGlobal(f *Frame) error {
	idx := binary.BigEndian.Uint16(f.Instructions()[f.InstrOffset+1:])

	vm.globals[idx] = vm.pop()

	f.InstrOffset += 3

	return nil
}

func (vm *VM) exeSetLocal(f *Frame) error {
	localIdx := f.Instructions()[f.InstrOffset+1]

	vm.stack[f.BaseStackCount+int(localIdx)] = vm.pop()

	f.InstrOffset += 2

	return nil
}

func (vm *VM) exeClosure(f *Frame) error {
	constIdx := binary.BigEndian.Uint16(f.Instructions()[f.InstrOffset+1:])
	numFree := f.Instructions()[f.InstrOffset+3]

	if err := vm.pushClosure(int(constIdx), int(numFree)); err != nil {
		return err
	}

	f.InstrOffset += 4

	return nil
}

func (vm *VM) exeJumpIfFalse(f *Frame) error {
	if vm.pop().Equal(False) {
		f.InstrOffset = binary.BigEndian.Uint64(f.Instructions()[f.InstrOffset+1:])
	} else {
		f.InstrOffset += 9
	}

	return nil
}

func (vm *VM) exeCall(f *Frame) error {
	numArgs := f.Instructions()[f.InstrOffset+1]

	f.InstrOffset += 2

	return vm.callFunc(int(numArgs))
}

func (vm *VM) callFunc(numArgs int) error {
	if vm.stackCount < (numArgs + 1) {
		return fmt.Errorf("stack count is too low for func+args")
	}

	closure, ok := vm.stack[vm.stackCount-1-numArgs].(*objects.Closure)
	if !ok {
		return fmt.Errorf("obj is not a compiled func")
	}

	// vm.lastPopped = ???

	frame := NewFrame(closure, vm.stackCount-numArgs)
	if err := vm.pushFrame(frame); err != nil {
		return err
	}

	vm.stackCount = frame.BaseStackCount + closure.Func.NumLocals

	return nil
}

func (vm *VM) exeReturn(f *Frame) error {
	vm.popFrame()

	vm.stackCount -= f.BaseStackCount

	f.InstrOffset += 1

	return nil
}

func (vm *VM) exeReturnValue(f *Frame) error {
	returnVal := vm.pop()

	vm.popFrame()

	vm.stackCount -= f.BaseStackCount

	f.InstrOffset++

	return vm.push(returnVal)
}

func (vm *VM) exeUnaryOp(f *Frame, method string) error {
	subject := vm.pop()
	result, err := subject.Send(method)
	if err != nil {
		return fmt.Errorf("method %s failed: %w", method, err)
	}

	f.InstrOffset++

	return vm.push(result)
}

func (vm *VM) exeBinaryOp(f *Frame, method string) error {
	right := vm.pop()
	left := vm.pop()

	result, err := left.Send(method, right)
	if err != nil {
		return fmt.Errorf("method %s failed: %w", method, err)
	}

	f.InstrOffset++

	return vm.push(result)
}

func (vm *VM) push(obj slang.Object) error {
	if vm.stackCount >= StackSize {
		return ErrStackOverflow
	}

	vm.stack[vm.stackCount] = obj

	vm.stackCount++

	return nil
}

func (vm *VM) pop() slang.Object {
	obj := vm.stack[vm.stackCount-1]

	vm.stackCount--

	vm.lastPopped = obj

	return obj
}

func (vm *VM) pushFrame(f *Frame) error {
	if vm.frameCount == MaxFrames {
		return fmt.Errorf("max frame count %d reached", MaxFrames)
	}

	vm.frames = append(vm.frames, f)
	vm.frameCount++

	return nil
}

func (vm *VM) popFrame() *Frame {
	f := vm.frames[vm.frameCount-1]

	vm.frames = vm.frames[:vm.frameCount-1]

	vm.frameCount--

	vm.lastPoppedFrame = f

	return f
}

func (vm *VM) pushClosure(constIndex, numFree int) error {
	constant := vm.constants[constIndex]
	function, ok := constant.(*objects.CompiledFunc)
	if !ok {
		return fmt.Errorf("not a function: %+v", constant)
	}

	free := make([]slang.Object, numFree)
	for i := 0; i < numFree; i++ {
		free[i] = vm.stack[vm.stackCount-numFree+i]
	}
	vm.stackCount -= numFree

	closure := objects.NewClosure(function, free...)

	return vm.push(closure)
}
