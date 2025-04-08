package parsers_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/field"
	"github.com/jamestunnell/slang/ast/types"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
)

type funcSigParserSuccessTest struct {
	TestName string
	Input    string
	Inputs   []*field.Field
	Outputs  []*field.Field
}

func TestFuncSignatureParser_NoInput(t *testing.T) {
	testFuncSignatureParserFail(t, "")
}

func TestFuncSignatureParser_TwoFieldsOneLine(t *testing.T) {
	testFuncSignatureParserFail(t, `(a int b mymodule.MyType)`)
}

func TestFuncSignatureParserSuccess(t *testing.T) {
	tests := []*funcSigParserSuccessTest{
		{
			TestName: "empty sig",
			Input:    `()`,
			Inputs:   []*field.Field{},
			Outputs:  []*field.Field{},
		},
		{
			TestName: "one param",
			Input:    `(x flt)`,
			Inputs:   []*field.Field{field.New("x", types.NewFlt())},
			Outputs:  []*field.Field{},
		},
		{
			TestName: "multi-name",
			Input: `(
				a, b int
				c, d str
			)`,
			Inputs: []*field.Field{
				field.New("a", types.NewInt()),
				field.New("b", types.NewInt()),
				field.New("c", types.NewStr()),
				field.New("d", types.NewStr()),
			},
			Outputs: []*field.Field{},
		},
		{
			TestName: "empty in and out params both with parens",
			Input:    `() ()`,
			Inputs:   []*field.Field{},
			Outputs:  []*field.Field{},
		},
		{
			TestName: "one in param, one out param",
			Input:    `(a int) (b int)`,
			Inputs: []*field.Field{
				field.New("a", types.NewInt()),
			},
			Outputs: []*field.Field{
				field.New("b", types.NewInt()),
			},
		},
		{
			TestName: "two out params",
			Input: `() (
				t my.Type
				e err
			)`,
			Inputs: []*field.Field{},
			Outputs: []*field.Field{
				field.New("t", types.NewStruct("my", "Type")),
				field.New("e", types.NewErr()),
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

		actual := expressions.NewFunc(p.Inputs, p.Outputs)
		expected := expressions.NewFunc(test.Inputs, test.Outputs)

		assert.True(t, actual.IsEqual(expected))
	})
}

func testFuncSignatureParserFail(t *testing.T, input string) {
	p := parsers.NewFuncSignatureParser()
	l := lexing.NewLexer(strings.NewReader(input))
	seq := parsing.NewTokenSeq(l)

	p.Run(seq)

	assert.NotEmpty(t, p.GetErrors)
}
