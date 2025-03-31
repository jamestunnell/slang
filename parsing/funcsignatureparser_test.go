package parsing_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/lexing"
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
			TestName: "two multi-name",
			Input:    `(a, b int, c, d string)`,
			Params: []slang.Param{
				ast.NewParam("a", ast.NewBasicType("int")),
				ast.NewParam("b", ast.NewBasicType("int")),
				ast.NewParam("c", ast.NewBasicType("string")),
				ast.NewParam("d", ast.NewBasicType("string")),
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
		l := lexing.NewLexer(strings.NewReader(test.Input))
		seq := parsing.NewTokenSeq(l)

		p.Run(seq)

		if !assert.Empty(t, p.GetErrors()) {
			logParseErrs(t, p.GetErrors())

			t.FailNow()
		}

		actual := expressions.NewFunc(p.Params, p.ReturnTypes)
		expected := expressions.NewFunc(test.Params, test.ReturnTypes)

		assert.True(t, actual.Equal(expected))
	})
}

func testFuncSignatureParserFail(t *testing.T, testName, input string) {
	t.Run(testName, func(t *testing.T) {
		p := parsing.NewFuncSignatureParser()
		l := lexing.NewLexer(strings.NewReader(input))
		seq := parsing.NewTokenSeq(l)

		p.Run(seq)

		assert.NotEmpty(t, p.GetErrors)
	})
}
