package parsing_test

import (
	"encoding/json"
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

func TestFileParserGlobalVars(t *testing.T) {
	file := strings.NewReader(`
		use "rand"

	  var x int
		var y int

		func init() {
			x = rand.Int()
			y = rand.Int()
		}

		func GetX() int {
			return x
		}

		func GetY() int {
			return y
		}
	`)
	expected := []slang.Statement{
		statements.NewUse("rand"),
		statements.NewVar("x", "int"),
		statements.NewVar("y", "int"),
		statements.NewFunc("init", ast.NewFunction(
			[]*ast.Param{},
			[]string{},
			statements.NewAssign(
				expressions.NewIdentifier("x"),
				expressions.NewCall(
					expressions.NewMemberAccess(expressions.NewIdentifier("rand"), "Int"),
				),
			),
			statements.NewAssign(
				expressions.NewIdentifier("y"),
				expressions.NewCall(
					expressions.NewMemberAccess(expressions.NewIdentifier("rand"), "Int"),
				),
			),
		)),
		statements.NewFunc("GetX", ast.NewFunction(
			[]*ast.Param{},
			[]string{"int"},
			statements.NewReturnVal(
				expressions.NewIdentifier("x"),
			),
		)),
		statements.NewFunc("GetY", ast.NewFunction(
			[]*ast.Param{},
			[]string{"int"},
			statements.NewReturnVal(
				expressions.NewIdentifier("y"),
			),
		)),
	}
	testFileParserSuccess(t, "just a class statement", file, expected)
}

func TestFileParserClassWithTest(t *testing.T) {
	file := strings.NewReader(`
		use "test"

		class Accumulator {
			field total float

			method Add(x float) {
				this.total = this.total + x
			}

			method Mul(x float) {
				this.total = this.total * x
			}

			method Total() float {
				return this.total
			}
		}

		func TestAccumulator(t test.Test) {
			accum = Accumulator()
			
			accum.Add(2.0)
			accum.Mul(2.0)

			t.AssertAlmostEq(accum.Total(), 4.0)

			accum.Add(1.0)
			accum.Mul(0.5)

			t.AssertAlmostEq(accum.Total(), 2.5)
		}
	`)
	expected := []slang.Statement{
		statements.NewUse("test"),
		statements.NewClass("Accumulator", "",
			statements.NewField("total", "float"),
			statements.NewMethod(
				"Add",
				ast.NewFunction(
					[]*ast.Param{ast.NewParam("x", "float")},
					[]string{},
					statements.NewAssign(
						expressions.NewMemberAccess(
							expressions.NewIdentifier("this"), "total"),
						expressions.NewAdd(
							expressions.NewMemberAccess(
								expressions.NewIdentifier("this"), "total"),
							expressions.NewIdentifier("x"),
						),
					),
				),
			),
			statements.NewMethod(
				"Mul",
				ast.NewFunction(
					[]*ast.Param{ast.NewParam("x", "float")},
					[]string{},
					statements.NewAssign(
						expressions.NewMemberAccess(
							expressions.NewIdentifier("this"), "total"),
						expressions.NewMultiply(
							expressions.NewMemberAccess(
								expressions.NewIdentifier("this"), "total"),
							expressions.NewIdentifier("x"),
						),
					),
				),
			),
			statements.NewMethod(
				"Total",
				ast.NewFunction(
					[]*ast.Param{},
					[]string{"float"},
					statements.NewReturnVal(
						expressions.NewMemberAccess(
							expressions.NewIdentifier("this"), "total"),
					),
				),
			),
		),
		statements.NewFunc("TestAccumulator", ast.NewFunction(
			[]*ast.Param{ast.NewParam("t", "test.Test")},
			[]string{},
			statements.NewAssign(
				expressions.NewIdentifier("accum"),
				expressions.NewCall(
					expressions.NewIdentifier("Accumulator"),
				),
			),
			statements.NewExpression(
				expressions.NewCall(
					expressions.NewMemberAccess(expressions.NewIdentifier("accum"), "Add"),
					expressions.NewPositionalArg(expressions.NewFloat(2.0)),
				),
			),
			statements.NewExpression(
				expressions.NewCall(
					expressions.NewMemberAccess(expressions.NewIdentifier("accum"), "Mul"),
					expressions.NewPositionalArg(expressions.NewFloat(2.0)),
				),
			),
			statements.NewExpression(
				expressions.NewCall(
					expressions.NewMemberAccess(expressions.NewIdentifier("t"), "AssertAlmostEq"),
					expressions.NewPositionalArg(expressions.NewCall(
						expressions.NewMemberAccess(expressions.NewIdentifier("accum"), "Total"),
					)),
					expressions.NewPositionalArg(expressions.NewFloat(4.0)),
				),
			),
			statements.NewExpression(
				expressions.NewCall(
					expressions.NewMemberAccess(expressions.NewIdentifier("accum"), "Add"),
					expressions.NewPositionalArg(expressions.NewFloat(1.0)),
				),
			),
			statements.NewExpression(
				expressions.NewCall(
					expressions.NewMemberAccess(expressions.NewIdentifier("accum"), "Mul"),
					expressions.NewPositionalArg(expressions.NewFloat(0.5)),
				),
			),
			statements.NewExpression(
				expressions.NewCall(
					expressions.NewMemberAccess(expressions.NewIdentifier("t"), "AssertAlmostEq"),
					expressions.NewPositionalArg(expressions.NewCall(
						expressions.NewMemberAccess(expressions.NewIdentifier("accum"), "Total"),
					)),
					expressions.NewPositionalArg(expressions.NewFloat(2.5)),
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
		if !assert.True(t, stmt.Equal(actual[i])) {
			actualD, _ := json.Marshal(actual[i])
			expectedD, _ := json.Marshal(stmt)

			t.Logf("statment %d not equal: \nactual: %s\nexpected: %s", i, string(actualD), string(expectedD))
		}
	}
}

func logParseErrs(t *testing.T, parseErrs []*parsing.ParseErr) {
	for _, parseErr := range parseErrs {
		t.Logf("unxpected parse err at %s: %v", parseErr.Token.Location, parseErr.Error)
	}
}
