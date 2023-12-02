package runtime_test

import (
	"bufio"
	"errors"
	"strings"
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/compiler"
	"github.com/jamestunnell/slang/lexer"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/runtime"
	"github.com/jamestunnell/slang/runtime/objects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVMEmptyCode(t *testing.T) {
	code := runtime.NewBytecode()
	vm := runtime.NewVM(code)

	err := vm.Step()

	assert.Error(t, err)
}

func TestVMPushConstant(t *testing.T) {
	code := runtime.NewBytecode()

	_, _ = code.AddConstant(objects.NewBool(true))
	_, _ = code.AddConstant(objects.NewBool(false))
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
	testVMWithExprStmt(t, "1 + 7", i(8))
	testVMWithExprStmt(t, "12 - 5", i(7))
	testVMWithExprStmt(t, "3 * 11", i(33))
	testVMWithExprStmt(t, "60 / 5", i(12))

	// comparison
	testVMWithExprStmt(t, "6 == 6", b(true))
	testVMWithExprStmt(t, "6 == 3", b(false))
	testVMWithExprStmt(t, "6 != 6", b(false))
	testVMWithExprStmt(t, "6 != 3", b(true))
	testVMWithExprStmt(t, "6 < 6", b(false))
	testVMWithExprStmt(t, "6 < 13", b(true))
	testVMWithExprStmt(t, "6 < 3", b(false))
	testVMWithExprStmt(t, "6 <= 6", b(true))
	testVMWithExprStmt(t, "6 <= 33", b(true))
	testVMWithExprStmt(t, "6 <= 3", b(false))
	testVMWithExprStmt(t, "6 > 6", b(false))
	testVMWithExprStmt(t, "6 > 3", b(true))
	testVMWithExprStmt(t, "6 > 13", b(false))
	testVMWithExprStmt(t, "6 >= 6", b(true))
	testVMWithExprStmt(t, "6 >= 3", b(true))
	testVMWithExprStmt(t, "6 >= 13", b(false))
}

func testVMWithExprStmt(
	t *testing.T,
	input string,
	expected slang.Object) {
	t.Run(input, func(t *testing.T) {
		l := lexer.New(bufio.NewReader(strings.NewReader(input)))
		toks := parsing.NewTokenSeq(l)
		p := parsing.NewExprOrAssignStatementParser()

		if !assert.True(t, p.Run(toks)) {
			logParseErrs(t, p.GetErrors())

			return
		}

		c := compiler.New()

		require.NoError(t, c.ProcessStmt(p.Stmt))

		code := c.GetCode()

		vm := runtime.NewVM(code)

		err := vm.Step()

		for err == nil {
			err = vm.Step()
		}

		if !errors.Is(err, runtime.ErrEndOfProgram) {
			t.Fatalf("VM failed unexpectedly: %v", err)
		}

		last := vm.LastPopped()

		require.NotNil(t, last)

		if !assert.True(t, expected.Equal(vm.LastPopped())) {
			t.Logf("%s (actual) != %s (expected)", last.Inspect(), expected.Inspect())
		}
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

func logParseErrs(t *testing.T, parseErrs []*parsing.ParseErr) {
	for _, parseErr := range parseErrs {
		t.Logf("unxpected parse err at %s: %v", parseErr.Token.Location, parseErr.Error)
	}
}
