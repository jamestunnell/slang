package parsing_test

import (
	"io"
	"strings"
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
	"github.com/stretchr/testify/assert"
)

func TestFileParser(t *testing.T) {
	file := strings.NewReader(`
		use "first/path"
			
		use "path/number/2"
		use "path-3"

		class AmountAdder {
			field Amount float

			method AddAmount(x float) float {
				return this.Amount + x
			}

			method ChangeAmount(amt float) {
				this.Amount = amt
			}
		}

		func TestAmountAdder(t test.Test) {
			aa = AmountAdder.New()

			aa.ChangeAmount(2.5)

			result = aa.AddAmount(12.0)

			t.AssertAlmostEq(result, 14.5)
		}
	`)
	expected := []slang.Statement{
		statements.NewUse("first/path"),
		statements.NewUse("path/number/2"),
		statements.NewUse("path-3"),
		statements.NewClass("AmountAdder", "",
			statements.NewField("Amount", "float"),
			statements.NewMethod(
				"AddAmount",
				ast.NewFunction(
					[]*ast.Param{ast.NewParam("x", "float")},
					[]string{"float"},
					statements.NewReturn(
						expressions.NewAdd(
							expressions.NewMemberAccess(
								expressions.NewIdentifier("this"), "Amount"),
							expressions.NewIdentifier("x"),
						),
					),
				),
			),
			statements.NewMethod(
				"ChangeAmount",
				ast.NewFunction(
					[]*ast.Param{ast.NewParam("amt", "float")},
					[]string{},
					statements.NewAssign(
						expressions.NewMemberAccess(
							expressions.NewIdentifier("this"), "Amount"),
						expressions.NewIdentifier("amt"),
					),
				),
			),
		),
		statements.NewFunc("TestAmountAdder", ast.NewFunction(
			[]*ast.Param{ast.NewParam("t", "test.Test")},
			[]string{},
			statements.NewAssign(
				expressions.NewIdentifier("aa"),
				expressions.NewCall(
					expressions.NewMemberAccess(expressions.NewIdentifier("AmountAdder"), "New")),
			),
			statements.NewExpression(
				expressions.NewCall(
					expressions.NewMemberAccess(expressions.NewIdentifier("aa"), "ChangeAmount"),
					expressions.NewFloat(2.5),
				),
			),
			statements.NewAssign(
				expressions.NewIdentifier("result"),
				expressions.NewCall(
					expressions.NewMemberAccess(expressions.NewIdentifier("aa"), "AddAmount"),
					expressions.NewFloat(12.0),
				),
			),
			statements.NewExpression(
				expressions.NewCall(
					expressions.NewMemberAccess(expressions.NewIdentifier("t"), "AssertAlmostEq"),
					expressions.NewIdentifier("result"),
					expressions.NewFloat(14.5),
				),
			),
		)),
	}
	testFileParserSuccess(t, "just a class statement", file, expected)
}

func testFileParserSuccess(
	t *testing.T,
	name string,
	file io.Reader,
	expected []slang.Statement,
) {
	t.Run(name, func(t *testing.T) {
		stmts, parseErrs := parsing.ParseFile(file)

		if !assert.Empty(t, parseErrs) {
			logParseErrs(t, parseErrs)

			return
		}

		verifyStatemnts(t, expected, stmts)
	})
}

func verifyStatemnts(t *testing.T, expected, actual []slang.Statement) {
	if !assert.Equal(t, len(expected), len(actual)) {
		return
	}

	for i, stmt := range expected {
		assert.True(t, stmt.Equal(actual[i]), "statment %d not equal", i)
	}
}

func logParseErrs(t *testing.T, parseErrs []*parsing.ParseErr) {
	for _, parseErr := range parseErrs {
		t.Logf("unxpected parse err at %s: %v", parseErr.Token.Location, parseErr.Error)
	}
}
