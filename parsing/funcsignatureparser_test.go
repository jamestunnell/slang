package parsing_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/lexer"
	"github.com/jamestunnell/slang/parsing"
)

type funcSigParserSuccessTest struct {
	TestName    string
	Input       string
	Params      []slang.Param
	ReturnTypes []slang.Type
}

func TestFuncSignatureParserFailure(t *testing.T) {
	testFuncSignatureParserFail(t, "no input", "")
}

func TestFuncSignatureParserSuccess(t *testing.T) {
	tests := []*funcSigParserSuccessTest{
		{
			TestName:    "empty sig",
			Input:       `()`,
			Params:      []slang.Param{},
			ReturnTypes: []slang.Type{},
		},
		{
			TestName:    "one param",
			Input:       `(x float)`,
			Params:      []slang.Param{ast.NewParam("x", ast.NewBasicType("float"))},
			ReturnTypes: []slang.Type{},
		},
		{
			TestName: "two params",
			Input:    `(a int, b mymodule.MyType)`,
			Params: []slang.Param{
				ast.NewParam("a", ast.NewBasicType("int")),
				ast.NewParam("b", ast.NewBasicType("mymodule", "MyType")),
			},
			ReturnTypes: []slang.Type{},
		},
		{
			TestName: "one param, one return type",
			Input:    `(a int) int`,
			Params: []slang.Param{
				ast.NewParam("a", ast.NewBasicType("int")),
			},
			ReturnTypes: []slang.Type{ast.NewBasicType("int")},
		},
		{
			TestName:    "two return types",
			Input:       `() (my.Type, error)`,
			Params:      []slang.Param{},
			ReturnTypes: []slang.Type{ast.NewBasicType("my", "Type"), ast.NewBasicType("error")},
		},
	}

	for _, test := range tests {
		testFuncSignatureParserSuccess(t, test)
	}
}

func testFuncSignatureParserSuccess(t *testing.T, test *funcSigParserSuccessTest) {
	t.Run(test.TestName, func(t *testing.T) {
		p := parsing.NewFuncSignatureParser()
		l := lexer.New(strings.NewReader(test.Input))
		seq := parsing.NewTokenSeq(l)

		p.Run(seq)

		if !assert.Empty(t, p.GetErrors()) {
			logParseErrs(t, p.GetErrors())

			t.FailNow()
		}

		fnActual := ast.NewFunction(p.Params, p.ReturnTypes)
		fnExpected := ast.NewFunction(test.Params, test.ReturnTypes)

		verifyFunc(t, fnExpected, fnActual)
	})
}

func testFuncSignatureParserFail(t *testing.T, testName, input string) {
	t.Run(testName, func(t *testing.T) {
		p := parsing.NewFuncSignatureParser()
		l := lexer.New(strings.NewReader(input))
		seq := parsing.NewTokenSeq(l)

		p.Run(seq)

		assert.NotEmpty(t, p.GetErrors)
	})
}

func verifyFunc(t *testing.T, expected, actual slang.Function) {
	paramNames := actual.GetParamNames()
	if assert.ElementsMatch(t, expected.GetParamNames(), paramNames) {
		for _, name := range paramNames {
			expectedType, ok1 := expected.GetParamType(name)
			actualType, ok2 := actual.GetParamType(name)

			if !assert.True(t, ok1 && ok2) {
				assert.Equal(t, actualType, expectedType)
			}
		}
	}

	assert.ElementsMatch(t, expected.GetReturnTypes(), actual.GetReturnTypes())

	// TODO verify body
}
