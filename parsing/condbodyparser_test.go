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
					expressions.NewFuncCall(
						expressions.NewMemberAccess(expressions.NewIdentifier("this"), "MyMethod")),
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