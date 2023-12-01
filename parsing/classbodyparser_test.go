package parsing_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexer"
	"github.com/jamestunnell/slang/parsing"
)

type bodyParserSuccessTest struct {
	TestName   string
	Input      string
	Statements []slang.Statement
}

func TestClassBodyParserFailure(t *testing.T) {
	testClassBodyParserFail(t, "no input", "")
}

func TestClassBodyParserSuccess(t *testing.T) {
	tests := []*bodyParserSuccessTest{
		{
			TestName:   "empty",
			Input:      `{}`,
			Statements: []slang.Statement{},
		},
		{
			TestName: "with fields",
			Input: `{
				field X myMod.myType
				field Y float
			}`,
			Statements: []slang.Statement{
				statements.NewField("X", "myMod.myType"),
				statements.NewField("Y", "float"),
			},
		},
		{
			TestName: "with empty methods",
			Input: `{
				method X(){}
				method Y(a int){}
			}`,
			Statements: []slang.Statement{
				statements.NewMethod("X", ast.NewFunction([]*ast.Param{}, []string{})),
				statements.NewMethod("Y", ast.NewFunction([]*ast.Param{ast.NewParam("a", "int")}, []string{})),
			},
		},
		{
			TestName: "non-empty method",
			Input: `{
				method sub(x int, y int) int {
					return x - y
				}
			}`,
			Statements: []slang.Statement{
				statements.NewMethod("sub", ast.NewFunction(
					[]*ast.Param{{Name: "x", Type: "int"}, {Name: "y", Type: "int"}},
					[]string{"int"},
					statements.NewReturn(sub(id("x"), id("y"))),
				)),
			},
		},
	}

	for _, test := range tests {
		testClassBodyParserSuccess(t, test)
	}
}

func testClassBodyParserSuccess(t *testing.T, test *bodyParserSuccessTest) {
	newParser := func() parsing.BodyParser { return parsing.NewClassBodyParser() }

	testBodyParserSuccess(t, test, newParser)
}

func testBodyParserSuccess(
	t *testing.T,
	test *bodyParserSuccessTest,
	newParser func() parsing.BodyParser) {
	t.Run(test.TestName, func(t *testing.T) {
		p := newParser()
		l := lexer.New(strings.NewReader(test.Input))
		seq := parsing.NewTokenSeq(l)

		p.Run(seq)

		if !assert.Empty(t, p.GetErrors()) {
			logParseErrs(t, p.GetErrors())

			return
		}

		verifyStatemnts(t, test.Statements, p.GetStatements())
	})
}

func testClassBodyParserFail(t *testing.T, testName, input string) {
	t.Run(testName, func(t *testing.T) {
		p := parsing.NewClassBodyParser()
		l := lexer.New(strings.NewReader(input))
		seq := parsing.NewTokenSeq(l)

		p.Run(seq)

		assert.NotEmpty(t, p.GetErrors)
	})
}
