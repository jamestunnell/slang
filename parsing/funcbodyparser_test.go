package parsing_test

import (
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
)

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
		{
			TestName: "assign to object field",
			Input: `{
				this.X = 2

				person.Name = "Jill"
			}`,
			Statements: []slang.Statement{
				statements.NewAssign(
					expressions.NewAccessMember(expressions.NewIdentifier("this"), "X"),
					expressions.NewInteger(2),
				),
				statements.NewAssign(
					expressions.NewAccessMember(expressions.NewIdentifier("person"), "Name"),
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
						expressions.NewCall(
							expressions.NewAccessMember(
								expressions.NewIdentifier("a"),
								"b",
							),
							expressions.NewPositionalArg(expressions.NewIdentifier("x")),
							expressions.NewPositionalArg(expressions.NewIdentifier("y")),
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
						expressions.NewString(""),
						expressions.NewIdentifier("word"),
						expressions.NewString(" is a "),
						expressions.NewCall(
							expressions.NewAccessMember(expressions.NewIdentifier("fanciness"), "String")),
						expressions.NewString(" word"),
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
