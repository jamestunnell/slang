package runtime_test

import (
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/objects"
	"github.com/jamestunnell/slang/runtime"
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
	code.AddInstructionUint16(runtime.OpConstant, 1)
	code.AddInstructionUint16(runtime.OpConstant, 0)
	code.AddInstructionUint16(runtime.OpConstant, 1)

	vm := runtime.NewVM(code)

	stepOKAndVerifyTop(t, vm, code.Constants[1])
	stepOKAndVerifyTop(t, vm, code.Constants[0])
	stepOKAndVerifyTop(t, vm, code.Constants[1])

	err := vm.Step()

	assert.Error(t, err)
	assert.ErrorIs(t, err, runtime.ErrEndOfProgram)
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

	return assert.True(t, obj.Equal(expected))
}
