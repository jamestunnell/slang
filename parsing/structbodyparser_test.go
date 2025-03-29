package parsing_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/parsing"
)

type bodyParserSuccessTest struct {
	TestName   string
	Input      string
	Statements []slang.Statement
	ErrorCount int
}

func TestStructBodyParserFailure(t *testing.T) {
	testStructBodyParserFail(t, "no input", "")
}

func TestStructBodyParserSuccess(t *testing.T) {
	tests := []*bodyParserSuccessTest{
		{
			TestName:   "empty",
			Input:      `{}`,
			Statements: []slang.Statement{},
		},
		{
			TestName: "basic field",
			Input: `{
				x int
			}`,
			Statements: []slang.Statement{
				statements.NewStructField([]string{"x"}, ast.NewBasicType("int")),
			},
		},
		{
			TestName: "multi-name field",
			Input: `{
				x, y, z myMod.myType
			}`,
			Statements: []slang.Statement{
				statements.NewStructField([]string{"x", "y", "z"}, ast.NewBasicType("myMod", "myType")),
			},
		},
		{
			TestName: "two basic fields",
			Input: `{
				a int
				b string
			}`,
			Statements: []slang.Statement{
				statements.NewStructField([]string{"a"}, ast.NewBasicType("int")),
				statements.NewStructField([]string{"b"}, ast.NewBasicType("string")),
			},
		},
	}

	for _, test := range tests {
		newParser := func() parsing.BodyParser { return parsing.NewStructBodyParser() }

		testBodyParserSuccess(t, test, newParser)
	}
}

func testBodyParserSuccess(
	t *testing.T,
	test *bodyParserSuccessTest,
	newParser func() parsing.BodyParser) {
	t.Run(test.TestName, func(t *testing.T) {
		p := newParser()
		l := lexing.NewLexer(strings.NewReader(test.Input))
		seq := parsing.NewTokenSeq(l)

		assert.True(t, p.Run(seq))

		if !assert.Len(t, p.GetErrors(), test.ErrorCount) {
			logParseErrs(t, p.GetErrors())

			return
		}

		verifyStatemnts(t, test.Statements, p.GetStatements())
	})
}

func testStructBodyParserFail(t *testing.T, testName, input string) {
	t.Run(testName, func(t *testing.T) {
		p := parsing.NewStructBodyParser()
		l := lexing.NewLexer(strings.NewReader(input))
		seq := parsing.NewTokenSeq(l)

		assert.False(t, p.Run(seq))

		assert.NotEmpty(t, p.GetErrors)
	})
}
