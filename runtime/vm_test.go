package runtime_test

import (
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/runtime"
	"github.com/jamestunnell/slang/runtime/objects"
	"github.com/stretchr/testify/assert"
)

func TestVMEmptyCode(t *testing.T) {
	code := runtime.NewBytecode()
	vm := runtime.NewVM(code)

	err := vm.Step()

	assert.Error(t, err)
}

func TestVMPushConstant(t *testing.T) {
	code := runtime.NewBytecode()

	_ = code.AddConstant(objects.NewBool(true))
	_ = code.AddConstant(objects.NewBool(false))
	code.AddInstructionUint16Operands(runtime.OpCONST, 1)
	code.AddInstructionUint16Operands(runtime.OpCONST, 0)
	code.AddInstructionUint16Operands(runtime.OpCONST, 1)

	vm := runtime.NewVM(code)

	stepOKAndVerifyTop(t, vm, code.Constants[1])
	stepOKAndVerifyTop(t, vm, code.Constants[0])
	stepOKAndVerifyTop(t, vm, code.Constants[1])

	err := vm.Step()

	assert.Error(t, err)
	assert.ErrorIs(t, err, runtime.ErrEndOfProgram)
}

func TestVMBinaryOps(t *testing.T) {
	// arithmetic
	testVMBinaryOp(t, "1 + -7", i(1), i(-7), runtime.OpADD, i(-6))
	testVMBinaryOp(t, "12 - 5", i(12), i(5), runtime.OpSUB, i(7))
	testVMBinaryOp(t, "3 * 11", i(3), i(11), runtime.OpMUL, i(33))
	testVMBinaryOp(t, "60 / 5", i(60), i(5), runtime.OpDIV, i(12))

	// comparison
	testVMBinaryOp(t, "6 == 6", i(6), i(6), runtime.OpEQ, b(true))
	testVMBinaryOp(t, "6 == 3", i(6), i(3), runtime.OpEQ, b(false))
	testVMBinaryOp(t, "6 != 6", i(6), i(6), runtime.OpNEQ, b(false))
	testVMBinaryOp(t, "6 != 3", i(6), i(3), runtime.OpNEQ, b(true))
	testVMBinaryOp(t, "6 < 6", i(6), i(6), runtime.OpLT, b(false))
	testVMBinaryOp(t, "6 < 13", i(6), i(13), runtime.OpLT, b(true))
	testVMBinaryOp(t, "6 < 3", i(6), i(3), runtime.OpLT, b(false))
	testVMBinaryOp(t, "6 <= 6", i(6), i(6), runtime.OpLEQ, b(true))
	testVMBinaryOp(t, "6 <= 33", i(6), i(33), runtime.OpLEQ, b(true))
	testVMBinaryOp(t, "6 <= 3", i(6), i(3), runtime.OpLEQ, b(false))
	testVMBinaryOp(t, "6 > 6", i(6), i(6), runtime.OpGT, b(false))
	testVMBinaryOp(t, "6 > 3", i(6), i(3), runtime.OpGT, b(true))
	testVMBinaryOp(t, "6 > 13", i(6), i(13), runtime.OpGT, b(false))
	testVMBinaryOp(t, "6 >= 6", i(6), i(6), runtime.OpGEQ, b(true))
	testVMBinaryOp(t, "6 >= 3", i(6), i(3), runtime.OpGEQ, b(true))
	testVMBinaryOp(t, "6 >= 13", i(6), i(13), runtime.OpGEQ, b(false))
}

func testVMBinaryOp(
	t *testing.T,
	name string,
	left, right slang.Object,
	opcode runtime.Opcode,
	expectedResult slang.Object) {
	t.Run(name, func(t *testing.T) {
		code := runtime.NewBytecode()

		_ = code.AddConstant(left)
		_ = code.AddConstant(right)
		code.AddInstructionUint16Operands(runtime.OpCONST, 0)
		code.AddInstructionUint16Operands(runtime.OpCONST, 1)
		code.AddInstructionNoOperands(opcode)

		vm := runtime.NewVM(code)

		stepOKAndVerifyTop(t, vm, code.Constants[0])

		assert.Equal(t, 1, vm.StackSize())

		stepOKAndVerifyTop(t, vm, code.Constants[1])

		assert.Equal(t, 2, vm.StackSize())

		stepOKAndVerifyTop(t, vm, expectedResult)

		assert.Equal(t, 1, vm.StackSize())
	})
}

func stepOKAndVerifyTop(
	t *testing.T, vm *runtime.VM, expected slang.Object) bool {
	err := vm.Step()

	if !assert.NoError(t, err) {
		return false
	}

	obj, ok := vm.Top()

	if !assert.True(t, ok) {
		return false
	}

	if !assert.True(t, obj.Equal(expected)) {
		t.Logf("%s (actual) != %s (expected)", obj.Inspect(), expected.Inspect())
		return false
	}

	return true
}

func i(val int64) slang.Object {
	return objects.NewInt(val)
}

func b(val bool) slang.Object {
	return objects.NewBool(val)
}
