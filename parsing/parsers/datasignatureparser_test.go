package parsers_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/jamestunnell/slang/ast/types"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
	"github.com/stretchr/testify/assert"
)

type dataSigParserSuccessTest struct {
	TestName  string
	Input     string
	NameTypes []*types.NameType
}

func TestDataSignatureParserFailure(t *testing.T) {
	testDataSignatureParserFail(t, "no input", "")
}

func TestDataSignatureParser_SuccessOneline(t *testing.T) {
	tests := []*dataSigParserSuccessTest{
		{
			TestName:  "empty sig",
			Input:     `()`,
			NameTypes: []*types.NameType{},
		},
		{
			TestName:  "one param",
			Input:     `(x flt)`,
			NameTypes: []*types.NameType{types.NewNameType("x", types.NewFlt())},
		},
		{
			TestName: "two params",
			Input:    `(a int, b mymodule.MyType)`,
			NameTypes: []*types.NameType{
				types.NewNameType("a", types.NewInt()),
				types.NewNameType("b", types.NewStruct("mymodule", "MyType")),
			},
		},
		{
			TestName: "two multi-name",
			Input:    `(a, b int, c, d str)`,
			NameTypes: []*types.NameType{
				types.NewNameType("a", types.NewInt()),
				types.NewNameType("b", types.NewInt()),
				types.NewNameType("c", types.NewStr()),
				types.NewNameType("d", types.NewStr()),
			},
		},
	}

	for _, test := range tests {
		testDataSignatureParserSuccess(t, test)
	}
}

func TestDataSignatureParser_SuccessMultiline(t *testing.T) {
	tests := []*dataSigParserSuccessTest{
		{
			TestName: "empty sig",
			Input: `(
			
			
			)`,
			NameTypes: []*types.NameType{},
		},
		{
			TestName: "one param",
			Input: `(
				x flt
			)`,
			NameTypes: []*types.NameType{types.NewNameType("x", types.NewFlt())},
		},
		{
			TestName: "two params",
			Input: `(
				a int
				b mymodule.MyType
			)`,
			NameTypes: []*types.NameType{
				types.NewNameType("a", types.NewInt()),
				types.NewNameType("b", types.NewStruct("mymodule", "MyType")),
			},
		},
		{
			TestName: "two multi-name",
			Input: `(
				a, b int
				c, d str
			)`,
			NameTypes: []*types.NameType{
				types.NewNameType("a", types.NewInt()),
				types.NewNameType("b", types.NewInt()),
				types.NewNameType("c", types.NewStr()),
				types.NewNameType("d", types.NewStr()),
			},
		},
		{
			TestName: "commented field",
			Input: `(
			    // a comment
				a int
			)`,
			NameTypes: []*types.NameType{
				types.NewNameType("a", types.NewInt()),
			},
		},
	}

	for _, test := range tests {
		testDataSignatureParserSuccess(t, test)
	}
}

func testDataSignatureParserSuccess(t *testing.T, test *dataSigParserSuccessTest) {
	t.Run(test.TestName, func(t *testing.T) {
		p := parsers.NewDataSignatureParser()
		l := lexing.NewLexer(strings.NewReader(test.Input))
		seq := parsing.NewTokenSeq(l)

		p.Run(seq)

		if !assert.Empty(t, p.GetErrors()) {
			logParseErrs(t, p.GetErrors())

			t.FailNow()
		}

		assert.True(t, slices.EqualFunc(p.NameTypes, test.NameTypes, nameTypesEqual))
	})
}

func testDataSignatureParserFail(t *testing.T, testName, input string) {
	t.Run(testName, func(t *testing.T) {
		p := parsers.NewDataSignatureParser()
		l := lexing.NewLexer(strings.NewReader(input))
		seq := parsing.NewTokenSeq(l)

		p.Run(seq)

		assert.NotEmpty(t, p.GetErrors)
	})
}

func nameTypesEqual(a, b *types.NameType) bool {
	return a.IsEqual(b)
}
