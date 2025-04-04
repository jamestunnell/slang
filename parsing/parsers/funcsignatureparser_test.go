package parsers_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/types"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
)

type funcSigParserSuccessTest struct {
	TestName  string
	Input     string
	InParams  []*types.NameType
	OutParams []*types.NameType
}

func TestFuncSignatureParserFailure(t *testing.T) {
	testFuncSignatureParserFail(t, "no input", "")
}

func TestFuncSignatureParserSuccess(t *testing.T) {
	tests := []*funcSigParserSuccessTest{
		{
			TestName:  "empty sig",
			Input:     `()`,
			InParams:  []*types.NameType{},
			OutParams: []*types.NameType{},
		},
		{
			TestName:  "one param",
			Input:     `(x flt)`,
			InParams:  []*types.NameType{types.NewNameType("x", types.NewFlt())},
			OutParams: []*types.NameType{},
		},
		{
			TestName: "two params",
			Input:    `(a int, b mymodule.MyType)`,
			InParams: []*types.NameType{
				types.NewNameType("a", types.NewInt()),
				types.NewNameType("b", types.NewStruct("mymodule", "MyType")),
			},
			OutParams: []*types.NameType{},
		},
		{
			TestName: "two multi-name",
			Input:    `(a, b int, c, d str)`,
			InParams: []*types.NameType{
				types.NewNameType("a", types.NewInt()),
				types.NewNameType("b", types.NewInt()),
				types.NewNameType("c", types.NewStr()),
				types.NewNameType("d", types.NewStr()),
			},
			OutParams: []*types.NameType{},
		},
		{
			TestName:  "empty in and out params both with parens",
			Input:     `() ()`,
			InParams:  []*types.NameType{},
			OutParams: []*types.NameType{},
		},
		{
			TestName: "one in param, one out param",
			Input:    `(a int) (b int)`,
			InParams: []*types.NameType{
				types.NewNameType("a", types.NewInt()),
			},
			OutParams: []*types.NameType{
				types.NewNameType("b", types.NewInt()),
			},
		},
		{
			TestName: "two out params",
			Input:    `() (t my.Type, e err)`,
			InParams: []*types.NameType{},
			OutParams: []*types.NameType{
				types.NewNameType("t", types.NewStruct("my", "Type")),
				types.NewNameType("e", types.NewErr()),
			},
		},
	}

	for _, test := range tests {
		testFuncSignatureParserSuccess(t, test)
	}
}

func testFuncSignatureParserSuccess(t *testing.T, test *funcSigParserSuccessTest) {
	t.Run(test.TestName, func(t *testing.T) {
		p := parsers.NewFuncSignatureParser()
		l := lexing.NewLexer(strings.NewReader(test.Input))
		seq := parsing.NewTokenSeq(l)

		p.Run(seq)

		if !assert.Empty(t, p.GetErrors()) {
			logParseErrs(t, p.GetErrors())

			t.FailNow()
		}

		actual := expressions.NewFunc(p.InParams, p.OutParams)
		expected := expressions.NewFunc(test.InParams, test.OutParams)

		assert.True(t, actual.IsEqual(expected))
	})
}

func testFuncSignatureParserFail(t *testing.T, testName, input string) {
	t.Run(testName, func(t *testing.T) {
		p := parsers.NewFuncSignatureParser()
		l := lexing.NewLexer(strings.NewReader(input))
		seq := parsing.NewTokenSeq(l)

		p.Run(seq)

		assert.NotEmpty(t, p.GetErrors)
	})
}
