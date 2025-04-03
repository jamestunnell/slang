package parsing_test

import (
	"strings"
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/parsing"
	"github.com/stretchr/testify/assert"
)

type bodyParserSuccessTest struct {
	TestName   string
	Input      string
	Statements []slang.Statement
	ErrorCount int
}

func TestFuncBodyParser(t *testing.T) {
	tests := []*bodyParserSuccessTest{
		{
			TestName:   "empty",
			Input:      `{}`,
			Statements: []slang.Statement{},
		},
		{
			TestName: "vars&consts",
			Input: `{
				var a int
				const b = "hello"
				var c flt
				const d = 12
			}`,
			Statements: []slang.Statement{
				statements.NewVar("a", &ast.IntType{}),
				statements.NewConst("b", expressions.NewStr("hello")),
				statements.NewVar("c", &ast.FltType{}),
				statements.NewConst("d", expressions.NewInt(12)),
			},
		},
		{
			TestName: "with comments",
			Input: `{
			    // this is a leading
				// standalone comment

			    // not empty
				const x = "hello"

				// this is a
				// standalone comment

				// also not empty
				const y = 10

				// this is a trailing
				// standalone comment
			}`,
			Statements: []slang.Statement{
				statements.NewComment("this is a leading", "standalone comment"),
				statements.WithComment(
					statements.NewConst("x", expressions.NewStr("hello")),
					"not empty",
				),
				statements.NewComment("this is a", "standalone comment"),
				statements.WithComment(
					statements.NewConst("y", expressions.NewInt(10)),
					"also not empty",
				),
				statements.NewComment("this is a trailing", "standalone comment"),
			},
		},
		{
			TestName: "assign to object field",
			Input: `{
				this.X = 2

				person.Name = "Jill"
			}`,
			Statements: []slang.Statement{
				statements.NewAssign(
					expressions.NewAccessMember(expressions.NewIdentifier("this"), "X"),
					expressions.NewInt(2),
				),
				statements.NewAssign(
					expressions.NewAccessMember(expressions.NewIdentifier("person"), "Name"),
					expressions.NewStr("Jill"),
				),
			},
		},
		{
			TestName: "call member method",
			Input: `{
				this.MyMethod()
			}`,
			Statements: []slang.Statement{
				statements.NewExpression(
					expressions.NewInvoke(
						expressions.NewAccessMember(expressions.NewIdentifier("this"), "MyMethod")),
				),
			},
		},
		{
			TestName: "member access/method call",
			Input: `{
				a.b(x, y).c
			}`,
			Statements: []slang.Statement{
				statements.NewExpression(
					expressions.NewAccessMember(
						expressions.NewInvoke(
							expressions.NewAccessMember(
								expressions.NewIdentifier("a"),
								"b",
							),
							expressions.NewInvokeArgPos(expressions.NewIdentifier("x")),
							expressions.NewInvokeArgPos(expressions.NewIdentifier("y")),
						),
						"c",
					),
				),
			},
		},
		{
			TestName: "assign string interpolation",
			Input: `{
				myVar = "${word} is a ${fanciness.String()} word"
			}`,
			Statements: []slang.Statement{
				statements.NewAssign(
					expressions.NewIdentifier("myVar"),
					expressions.NewConcat(
						expressions.NewStr(""),
						expressions.NewIdentifier("word"),
						expressions.NewStr(" is a "),
						expressions.NewInvoke(
							expressions.NewAccessMember(expressions.NewIdentifier("fanciness"), "String")),
						expressions.NewStr(" word"),
					),
				),
			},
		},
	}

	for _, test := range tests {
		testFuncBodyParserSuccess(t, test)
	}
}

func testFuncBodyParserSuccess(t *testing.T, test *bodyParserSuccessTest) {
	newParser := func() parsing.BodyParser { return parsing.NewFuncBodyParser() }

	testBodyParserSuccess(t, test, newParser)
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
