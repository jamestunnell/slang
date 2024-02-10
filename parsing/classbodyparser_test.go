package parsing_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexer"
	"github.com/jamestunnell/slang/parsing"
)

type bodyParserSuccessTest struct {
	TestName   string
	Input      string
	Statements []slang.Statement
	ErrorCount int
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
				statements.NewField("X", ast.NewBasicType("myMod", "myType")),
				statements.NewField("Y", ast.NewBasicType("float")),
			},
		},
		{
			TestName: "with empty methods",
			Input: `{
				method X(){}
				method Y(a int){}
			}`,
			Statements: []slang.Statement{
				statements.NewMethod("X", ast.NewFunction([]slang.Param{}, []slang.Type{})),
				statements.NewMethod("Y", ast.NewFunction([]slang.Param{ast.NewParam("a", ast.NewBasicType("int"))}, []slang.Type{})),
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
					[]slang.Param{
						ast.NewParam("x", ast.NewBasicType("int")),
						ast.NewParam("y", ast.NewBasicType("int")),
					},
					[]slang.Type{ast.NewBasicType("int")},
					statements.NewReturnVal(sub(id("x"), id("y"))),
				)),
			},
		},
		{
			TestName: "vars&consts",
			Input: `{
				var a int
				const b = "hello"
				var c float
				const d = 12 
			}`,
			Statements: []slang.Statement{
				statements.NewVar("a", ast.NewBasicType("int")),
				statements.NewConst("b", expressions.NewString("hello")),
				statements.NewVar("c", ast.NewBasicType("float")),
				statements.NewConst("d", expressions.NewInteger(12)),
			},
		},
		// {
		// 	TestName: "bad one-line statements are ignored",
		// 	Input: `{
		// 		field x int
		// 		field y 3
		// 		field z float
		// 	}`,
		// 	Statements: []slang.Statement{
		// 		statements.NewField("x", "int"),
		// 		statements.NewField("z", "float"),
		// 	},
		// 	ErrorCount: 1,
		// },
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

		assert.True(t, p.Run(seq))

		if !assert.Len(t, p.GetErrors(), test.ErrorCount) {
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

		assert.False(t, p.Run(seq))

		assert.NotEmpty(t, p.GetErrors)
	})
}
