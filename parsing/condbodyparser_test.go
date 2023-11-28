package parsing_test

import (
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
)

func TestCondBodyParser(t *testing.T) {
	tests := []*bodyParserSuccessTest{
		{
			TestName:   "empty",
			Input:      `{}`,
			Statements: []slang.Statement{},
		},
		{
			TestName: "assign to object field",
			Input: `{
				this.X = 2

				person.Name = "Jill"
			}`,
			Statements: []slang.Statement{
				statements.NewAssign(
					expressions.NewMemberAccess(expressions.NewIdentifier("this"), "X"),
					expressions.NewInteger(2),
				),
				statements.NewAssign(
					expressions.NewMemberAccess(expressions.NewIdentifier("person"), "Name"),
					expressions.NewString("Jill"),
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
					expressions.NewCall(
						expressions.NewMemberAccess(expressions.NewIdentifier("this"), "MyMethod")),
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
					expressions.NewMemberAccess(
						expressions.NewCall(
							expressions.NewMemberAccess(
								expressions.NewIdentifier("a"),
								"b",
							),
							expressions.NewIdentifier("x"),
							expressions.NewIdentifier("y"),
						),
						"c",
					),
				),
			},
		},
		{
			TestName: "assign nested string interpolation",
			Input: `{
				myVar = "${word} is a ${if word.Size() < 10 {"short"} else {"not short"}} word"
			}`,
			Statements: []slang.Statement{
				statements.NewAssign(
					expressions.NewIdentifier("myVar"),
					expressions.NewConcat(
						expressions.NewString(""),
						expressions.NewIdentifier("word"),
						expressions.NewString(" is a "),
						expressions.NewIfElse(
							expressions.NewLess(
								expressions.NewCall(
									expressions.NewMemberAccess(expressions.NewIdentifier("word"), "Size"),
								),
								expressions.NewInteger(10),
							),
							[]slang.Statement{statements.NewExpression(expressions.NewString("short"))},
							[]slang.Statement{statements.NewExpression(expressions.NewString("not short"))},
						),
						expressions.NewString(" word"),
					),
				),
			},
		},
	}

	for _, test := range tests {
		testCondBodyParserSuccess(t, test)
	}
}

func testCondBodyParserSuccess(t *testing.T, test *bodyParserSuccessTest) {
	newParser := func() parsing.StatementParser { return parsing.NewCondBodyParser() }

	testBodyParserSuccess(t, test, newParser)
}
