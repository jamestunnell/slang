package parsers_test

import (
	"testing"

	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/ast/types"
	"github.com/jamestunnell/slang/parsing/parsers"
)

func TestFuncBodyParser_EmptyBody(t *testing.T) {
	testFuncBodyParserSuccess(t, `{}`)
}

func TestFuncBodyParser_VarsAndConsts(t *testing.T) {
	testFuncBodyParserSuccess(t, `{
			var a int
			const b "hello"
			var c flt
			const d 12
		}`,
		statements.NewVar("a", types.NewInt()),
		statements.NewConst("b", expressions.NewStr("hello")),
		statements.NewVar("c", types.NewFlt()),
		statements.NewConst("d", expressions.NewInt(12)),
	)
}

func TestFuncBodyParser_WithComments(t *testing.T) {
	testFuncBodyParserSuccess(t, `{
			// this is a leading
			// standalone comment

			// not empty
			const x "hello"

			// this is a
			// standalone comment

			// also not empty
			const y 10

			// this is a trailing
			// standalone comment
		}`,
		withComment(statements.NewComment(), "this is a leading standalone comment"),
		withComment(statements.NewConst("x", expressions.NewStr("hello")), "not empty"),
		withComment(statements.NewComment(), "this is a standalone comment"),
		withComment(statements.NewConst("y", expressions.NewInt(10)), "also not empty"),
		withComment(statements.NewComment(), "this is a trailing standalone comment"),
	)
}

func TestFuncBodyParser_IfBlock(t *testing.T) {
	testFuncBodyParserSuccess(t, `{
			if x < 10 {
				print(x)

				x = x + 1
			}
		}`,
		statements.NewIf(
			expressions.NewLess(expressions.NewIdentifier("x"), expressions.NewInt(10)),
			[]*statements.Statement{
				statements.NewExpression(expressions.NewInvoke(
					expressions.NewIdentifier("print"),
					expressions.NewInvokeArgPos(expressions.NewIdentifier("x")),
				)),
				statements.NewAssign(
					expressions.NewIdentifier("x"),
					expressions.NewAdd(expressions.NewIdentifier("x"), expressions.NewInt(1)),
				),
			},
		),
	)
}

func TestFuncBodyParser_CallMemberMethod(t *testing.T) {
	testFuncBodyParserSuccess(t, `{
			this.MyMethod()
		}`,
		statements.NewExpression(
			expressions.NewInvoke(
				expressions.NewAccessMember(expressions.NewIdentifier("this"), "MyMethod")),
		),
	)
}

func TestFuncBodyParser_CallMemberAccessMethodCall(t *testing.T) {
	testFuncBodyParserSuccess(t, `{
			a.b(x y).c
		}`,
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
	)
}

func TestFuncBodyParser_AssignStringInterp(t *testing.T) {
	testFuncBodyParserSuccess(t, `{
			myVar = "${word} is a ${fanciness.String()} word"
		}`,
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
	)
}

func testFuncBodyParserSuccess(
	t *testing.T,
	input string,
	expected ...*statements.Statement) {
	testBodyParserSuccess(t, parsers.NewFuncBodyParser(), input, expected...)
}
