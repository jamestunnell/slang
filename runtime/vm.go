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

	opcode := Opcode(f.Instructions()[f.InstrOffset])

	var exec func(*Frame) error

	switch opcode {
	case OpGETCONST:
		exec = vm.execGetConst
	case OpGETGLOBAL:
		exec = vm.execGetGlobal
	case OpGETLOCAL:
		exec = vm.execGetLocal
	case OpGETFREE:
		exec = vm.execGetFree
	case OpSETGLOBAL:
		exec = vm.execSetGlobal
	case OpSETLOCAL:
		exec = vm.execSetLocal
	case OpCLOSURE:
		exec = vm.execClosure
	case OpCURRENTCLOSURE:
		exec = vm.execCurrentClosure
	case OpJUMP:
		exec = vm.execJump
	case OpJUMPIFFALSE:
		exec = vm.execJumpIfFalse
	case OpPOP:
		exec = vm.execPop
	case OpCALL:
		exec = vm.execCall
	case OpRETURN:
		exec = vm.execReturn
	case OpRETURNVAL:
		exec = vm.execReturnValue
	case OpNEG:
		exec = vm.execNeg
	case OpNOT:
		exec = vm.execNot
	case OpADD:
		exec = vm.execAdd
	case OpSUB:
		exec = vm.execSub
	case OpMUL:
		exec = vm.execMul
	case OpDIV:
		exec = vm.execDiv
	case OpEQ:
		exec = vm.execEq
	case OpNEQ:
		exec = vm.execNeq
	case OpLT:
		exec = vm.execLt
	case OpLEQ:
		exec = vm.execLeq
	case OpGT:
		exec = vm.execGt
	case OpGEQ:
		exec = vm.execGeq
	case OpAND:
		exec = vm.execAnd
	case OpOR:
		exec = vm.execOr
	}

	if exec == nil {
		return fmt.Errorf("unknown opcode %d", opcode)
	}

	return exec(f)
}

func (vm *VM) StackCount() int {
	return vm.stackCount
}

func (vm *VM) Top() (slang.Object, bool) {
	if vm.stackCount == 0 {
		return nil, false
	}

	return vm.stack[vm.stackCount-1], true
}

func (vm *VM) currentFrame() *Frame {
	return vm.frames[vm.frameCount-1]
}

func (vm *VM) execGetConst(f *Frame) error {
	idx := binary.BigEndian.Uint16(f.Instructions()[f.InstrOffset+1:])

	if int(idx) >= vm.numConsts {
		return fmt.Errorf("constant index %d is out of bounds", idx)
	}

	f.InstrOffset += 3

	return vm.push(vm.constants[idx])
}

func (vm *VM) execGetGlobal(f *Frame) error {
	idx := binary.BigEndian.Uint16(f.Instructions()[f.InstrOffset+1:])

	f.InstrOffset += 3

	return vm.push(vm.globals[idx])
}

func (vm *VM) execGetLocal(f *Frame) error {
	localIdx := f.Instructions()[f.InstrOffset+1]

	f.InstrOffset += 2

	return vm.push(vm.stack[f.BaseStackCount+int(localIdx)])
}

func (vm *VM) execGetFree(f *Frame) error {
	freeIdx := f.Instructions()[f.InstrOffset+1]

	f.InstrOffset += 2

	return vm.push(f.Closure.FreeVars[freeIdx])
}

func (vm *VM) execSetGlobal(f *Frame) error {
	idx := binary.BigEndian.Uint16(f.Instructions()[f.InstrOffset+1:])

	vm.globals[idx] = vm.pop()

	f.InstrOffset += 3

	return nil
}

func (vm *VM) execSetLocal(f *Frame) error {
	localIdx := f.Instructions()[f.InstrOffset+1]

	vm.stack[f.BaseStackCount+int(localIdx)] = vm.pop()

	f.InstrOffset += 2

	return nil
}

func (vm *VM) execClosure(f *Frame) error {
	constIdx := binary.BigEndian.Uint16(f.Instructions()[f.InstrOffset+1:])
	numFree := f.Instructions()[f.InstrOffset+3]

	if err := vm.pushClosure(int(constIdx), int(numFree)); err != nil {
		return err
	}

	f.InstrOffset += 4

	return nil
}

func (vm *VM) execCurrentClosure(f *Frame) error {
	f.InstrOffset++

	return vm.push(f.Closure)
}

func (vm *VM) execJump(f *Frame) error {
	f.InstrOffset = binary.BigEndian.Uint64(f.Instructions()[f.InstrOffset+1:])

	return nil
}

func (vm *VM) execJumpIfFalse(f *Frame) error {
	if vm.pop().Equal(False) {
		f.InstrOffset = binary.BigEndian.Uint64(f.Instructions()[f.InstrOffset+1:])
	} else {
		f.InstrOffset += 9
	}

	return nil
}

func (vm *VM) execPop(f *Frame) error {
	vm.pop()

	f.InstrOffset++

	return nil
}

func (vm *VM) execCall(f *Frame) error {
	numArgs := f.Instructions()[f.InstrOffset+1]

	f.InstrOffset += 2

	return vm.callFunc(int(numArgs))
}

func (vm *VM) callFunc(numArgs int) error {
	if vm.stackCount < (numArgs + 1) {
		return fmt.Errorf("stack count is too low for func+args")
	}

	closure, ok := vm.stack[vm.stackCount-(numArgs+1)].(*objects.Closure)
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

func (vm *VM) execReturn(f *Frame) error {
	vm.popFrame()

	vm.stackCount = f.BaseStackCount - 1

	f.InstrOffset++

	return nil
}

func (vm *VM) execReturnValue(f *Frame) error {
	returnVal := vm.pop()

	vm.popFrame()

	vm.stackCount = f.BaseStackCount - 1

	f.InstrOffset++

	return vm.push(returnVal)
}

func (vm *VM) execNeg(f *Frame) error {
	return vm.exeUnaryOp(f, slang.MethodNEG)
}

func (vm *VM) execNot(f *Frame) error {
	return vm.exeUnaryOp(f, slang.MethodNOT)
}

func (vm *VM) execAdd(f *Frame) error {
	return vm.exeBinaryOp(f, slang.MethodADD)
}

func (vm *VM) execSub(f *Frame) error {
	return vm.exeBinaryOp(f, slang.MethodSUB)
}

func (vm *VM) execMul(f *Frame) error {
	return vm.exeBinaryOp(f, slang.MethodMUL)
}

func (vm *VM) execDiv(f *Frame) error {
	return vm.exeBinaryOp(f, slang.MethodDIV)
}

func (vm *VM) execEq(f *Frame) error {
	return vm.exeBinaryOp(f, slang.MethodEQ)
}

func (vm *VM) execNeq(f *Frame) error {
	return vm.exeBinaryOp(f, slang.MethodNEQ)
}

func (vm *VM) execLt(f *Frame) error {
	return vm.exeBinaryOp(f, slang.MethodLT)
}

func (vm *VM) execLeq(f *Frame) error {
	return vm.exeBinaryOp(f, slang.MethodLEQ)
}

func (vm *VM) execGt(f *Frame) error {
	return vm.exeBinaryOp(f, slang.MethodGT)
}

func (vm *VM) execGeq(f *Frame) error {
	return vm.exeBinaryOp(f, slang.MethodGEQ)
}

func (vm *VM) execAnd(f *Frame) error {
	return vm.exeBinaryOp(f, slang.MethodAND)
}

func (vm *VM) execOr(f *Frame) error {
	return vm.exeBinaryOp(f, slang.MethodOR)
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
