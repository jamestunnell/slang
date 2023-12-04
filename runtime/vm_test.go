package runtime_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/compiler"
	"github.com/jamestunnell/slang/lexer"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/runtime"
	"github.com/jamestunnell/slang/runtime/instructions"
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
	is := runtime.Instructions{
		instructions.NewGetConst(1),
		instructions.NewGetConst(0),
		instructions.NewGetConst(1),
	}
	constants := []slang.Object{
		objects.NewBool(true),
		objects.NewBool(false),
	}
	code := &runtime.Bytecode{
		Instructions: is.Assemble(),
		Constants:    constants,
		MaxGlobals:   0,
	}

	vm := runtime.NewVM(code)

	stepOKAndVerifyTop(t, vm, code.Constants[1])
	stepOKAndVerifyTop(t, vm, code.Constants[0])
	stepOKAndVerifyTop(t, vm, code.Constants[1])

	err := vm.Step()

	assert.Error(t, err)
	assert.ErrorIs(t, err, runtime.ErrEndOfProgram)
}

func TestVMIntegerExprStmts(t *testing.T) {
	// basic integer arithmetic
	testVMWithFuncBodyStmts(t, "1 + 7", i(8))
	testVMWithFuncBodyStmts(t, "12 - 5", i(7))
	testVMWithFuncBodyStmts(t, "3 * 11", i(33))
	testVMWithFuncBodyStmts(t, "60 / 5", i(12))

	// more elaborate integer arithmetic
	testVMWithFuncBodyStmts(t, "(65 - 11) + (78 * 13)", i(1068))
	testVMWithFuncBodyStmts(t, "40 + 77 / (13 - 47)", i(38))
	testVMWithFuncBodyStmts(t, "(-12 * 5) / 3", i(-20))

	// basic integer comparison
	testVMWithFuncBodyStmts(t, "6 == 6", b(true))
	testVMWithFuncBodyStmts(t, "6 == 3", b(false))
	testVMWithFuncBodyStmts(t, "6 != 6", b(false))
	testVMWithFuncBodyStmts(t, "6 != 3", b(true))
	testVMWithFuncBodyStmts(t, "6 < 6", b(false))
	testVMWithFuncBodyStmts(t, "6 < 13", b(true))
	testVMWithFuncBodyStmts(t, "6 < 3", b(false))
	testVMWithFuncBodyStmts(t, "6 <= 6", b(true))
	testVMWithFuncBodyStmts(t, "6 <= 33", b(true))
	testVMWithFuncBodyStmts(t, "6 <= 3", b(false))
	testVMWithFuncBodyStmts(t, "6 > 6", b(false))
	testVMWithFuncBodyStmts(t, "6 > 3", b(true))
	testVMWithFuncBodyStmts(t, "6 > 13", b(false))
	testVMWithFuncBodyStmts(t, "6 >= 6", b(true))
	testVMWithFuncBodyStmts(t, "6 >= 3", b(true))
	testVMWithFuncBodyStmts(t, "6 >= 13", b(false))

	// and/or/not expression
	testVMWithFuncBodyStmts(t, "(7 == 12 or true) and 13 < 50", b(true))
	testVMWithFuncBodyStmts(t, "3 == 4 and true", b(false))
	testVMWithFuncBodyStmts(t, "!(3 == 4) and true", b(true))
}

func TestVM_IfExpr(t *testing.T) {
	testVMWithFuncBodyStmts(t, "if true {75}", i(75))
	testVMWithFuncBodyStmts(t, `if true {
			75
		}
		100`, i(100))
	testVMWithFuncBodyStmts(t, "if true {if false {12}}", b(false))
	testVMWithFuncBodyStmts(t, "if true {if true {12}}", i(12))
}

func TestVM_IfElseExpr(t *testing.T) {
	testVMWithFuncBodyStmts(t, "if true {75} else {58}", i(75))
	testVMWithFuncBodyStmts(t, "if false {75} else {58}", i(58))
	testVMWithFuncBodyStmts(t, `
		if true {
			75
		} else {
			20
		}
		88`, i(88))
	testVMWithFuncBodyStmts(t, `
		if false {
			75
		} else {
			20
		}
		88`, i(88))
}

func TestVM_AssignStmt(t *testing.T) {
	testVMWithFuncBodyStmts(t, "x = 12", i(12))
	testVMWithFuncBodyStmts(t, `
		x = 12
		5 * x`, i(60))
	testVMWithFuncBodyStmts(t, `
		x = 80
		y = -4
		x * y`, i(-320))
	testVMWithFuncBodyStmts(t, `
		x = 80
		y = -5
		x = x + y
		2 * x`, i(150))
}

func TestVM_FuncLiteral(t *testing.T) {
	testVMWithFuncBodyStmts(t, `
		func() int { return 7 }()
	`, i(7))

	testVMWithFuncBodyStmts(t, `
		add5 = func(x int) int { return x + 5 }
		add5(12)
	`, i(17))

	testVMWithFuncBodyStmts(t, `
	    y = 77
		addY = func(x int) int { return x + y }
		addY(10)
	`, i(87))
}

func TestVM_FuncStatement(t *testing.T) {
	// not a closure
	testVMWithFuncBodyStmts(t, `
		func Or(a bool, b bool) bool {
			return a or b
		}

		Or(true, false)
	`, b(true))

	// closure
	testVMWithFuncBodyStmts(t, `
	  amount = 2.5

		func AddAmount(x float) float {
			return x + amount
		}

		amount = 7.0

		AddAmount(13.0)
	`, f(20.0))

	// recursive
	testVMWithFuncBodyStmts(t, `
		func Fib(n int) int {
			if n == 0 {
				return 0
			}

			if n == 1 {
				return 1
			}

			return Fib(n-1) + Fib(n-2)
		}

		Fib(10)
	`, i(55))
}

func testVMWithFuncBodyStmts(
	t *testing.T,
	input string,
	expected slang.Object) {
	t.Run(input, func(t *testing.T) {
		vm, ok := setupVM(t, "{"+input+"}", parsing.NewFuncBodyParser())
		if !ok {
			return
		}

		if !runVM(t, vm) {
			return
		}

		verifyObject(t, vm.LastPopped(), expected)
	})
}

func verifyObject(t *testing.T, actual, expected slang.Object) bool {
	if !assert.True(t, actual.Equal(expected)) {
		t.Logf("%#v (actual) != %#v (expected)", actual, expected)

		return false
	}

	return true
}

func runVM(t *testing.T, vm *runtime.VM) bool {
	err := vm.Step()

	for err == nil {
		err = vm.Step()
	}

	if !assert.ErrorIs(t, err, runtime.ErrEndOfProgram) {
		t.Logf("VM failed unexpectedly: %v", err)

		return false
	}

	return true
}

func setupVM(t *testing.T, input string, parser parsing.BodyParser) (*runtime.VM, bool) {
	l := lexer.New(bufio.NewReader(strings.NewReader(input)))
	toks := parsing.NewTokenSeq(l)

	if !assert.True(t, parser.Run(toks)) {
		logParseErrs(t, parser.GetErrors())

		return nil, false
	}

	c := compiler.New()

	for _, stmt := range parser.GetStatements() {
		compilerErr := c.ProcessStmt(stmt)
		if !assert.NoError(t, compilerErr) {
			return nil, false
		}
	}

	code := c.MakeBytecode()

	return runtime.NewVM(code), true
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
		t.Logf("%#v (actual) != %#v (expected)", obj, expected)
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

func f(val float64) slang.Object {
	return objects.NewFloat(val)
}

func logParseErrs(t *testing.T, parseErrs []*parsing.ParseErr) {
	for _, parseErr := range parseErrs {
		t.Logf("unxpected parse err at %s: %v", parseErr.Token.Location, parseErr.Error)
	}
}
