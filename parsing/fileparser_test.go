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
		statements.NewVar("x", ast.NewBasicType("int")),
		statements.NewVar("y", ast.NewBasicType("int")),
		statements.NewFunc("init", ast.NewFunction(
			[]slang.Param{},
			[]slang.Type{},
			statements.NewAssign(
				expressions.NewIdentifier("x"),
				expressions.NewCall(
					expressions.NewAccessMember(expressions.NewIdentifier("rand"), "Int"),
				),
			),
			statements.NewAssign(
				expressions.NewIdentifier("y"),
				expressions.NewCall(
					expressions.NewAccessMember(expressions.NewIdentifier("rand"), "Int"),
				),
			),
		)),
		statements.NewFunc("GetX", ast.NewFunction(
			[]slang.Param{},
			[]slang.Type{ast.NewBasicType("int")},
			statements.NewReturnVal(
				expressions.NewIdentifier("x"),
			),
		)),
		statements.NewFunc("GetY", ast.NewFunction(
			[]slang.Param{},
			[]slang.Type{ast.NewBasicType("int")},
			statements.NewReturnVal(
				expressions.NewIdentifier("y"),
			),
		)),
	}
	testFileParserSuccess(t, "just a class statement", file, expected, 0)
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
			statements.NewField("total", ast.NewBasicType("float")),
			statements.NewMethod(
				"Add",
				ast.NewFunction(
					[]slang.Param{ast.NewParam("x", ast.NewBasicType("float"))},
					[]slang.Type{},
					statements.NewAssign(
						expressions.NewAccessMember(
							expressions.NewIdentifier("this"), "total"),
						expressions.NewAdd(
							expressions.NewAccessMember(
								expressions.NewIdentifier("this"), "total"),
							expressions.NewIdentifier("x"),
						),
					),
				),
			),
			statements.NewMethod(
				"Mul",
				ast.NewFunction(
					[]slang.Param{ast.NewParam("x", ast.NewBasicType("float"))},
					[]slang.Type{},
					statements.NewAssign(
						expressions.NewAccessMember(
							expressions.NewIdentifier("this"), "total"),
						expressions.NewMultiply(
							expressions.NewAccessMember(
								expressions.NewIdentifier("this"), "total"),
							expressions.NewIdentifier("x"),
						),
					),
				),
			),
			statements.NewMethod(
				"Total",
				ast.NewFunction(
					[]slang.Param{},
					[]slang.Type{ast.NewBasicType("float")},
					statements.NewReturnVal(
						expressions.NewAccessMember(
							expressions.NewIdentifier("this"), "total"),
					),
				),
			),
		),
		statements.NewFunc("TestAccumulator", ast.NewFunction(
			[]slang.Param{ast.NewParam("t", ast.NewBasicType("test", "Test"))},
			[]slang.Type{},
			statements.NewAssign(
				expressions.NewIdentifier("accum"),
				expressions.NewCall(
					expressions.NewIdentifier("Accumulator"),
				),
			),
			statements.NewExpression(
				expressions.NewCall(
					expressions.NewAccessMember(expressions.NewIdentifier("accum"), "Add"),
					expressions.NewPositionalArg(expressions.NewFloat(2.0)),
				),
			),
			statements.NewExpression(
				expressions.NewCall(
					expressions.NewAccessMember(expressions.NewIdentifier("accum"), "Mul"),
					expressions.NewPositionalArg(expressions.NewFloat(2.0)),
				),
			),
			statements.NewExpression(
				expressions.NewCall(
					expressions.NewAccessMember(expressions.NewIdentifier("t"), "AssertAlmostEq"),
					expressions.NewPositionalArg(expressions.NewCall(
						expressions.NewAccessMember(expressions.NewIdentifier("accum"), "Total"),
					)),
					expressions.NewPositionalArg(expressions.NewFloat(4.0)),
				),
			),
			statements.NewExpression(
				expressions.NewCall(
					expressions.NewAccessMember(expressions.NewIdentifier("accum"), "Add"),
					expressions.NewPositionalArg(expressions.NewFloat(1.0)),
				),
			),
			statements.NewExpression(
				expressions.NewCall(
					expressions.NewAccessMember(expressions.NewIdentifier("accum"), "Mul"),
					expressions.NewPositionalArg(expressions.NewFloat(0.5)),
				),
			),
			statements.NewExpression(
				expressions.NewCall(
					expressions.NewAccessMember(expressions.NewIdentifier("t"), "AssertAlmostEq"),
					expressions.NewPositionalArg(expressions.NewCall(
						expressions.NewAccessMember(expressions.NewIdentifier("accum"), "Total"),
					)),
					expressions.NewPositionalArg(expressions.NewFloat(2.5)),
				),
			),
		)),
	}
	testFileParserSuccess(t, "just a class statement", file, expected, 0)
}

// func TestFileParserBadOnelineStatements(t *testing.T) {
// 	file := strings.NewReader(`
// 		use "xyz"
// 		use what

// 		var x int
// 		var y 3
// 	`)
// 	expected := []slang.Statement{
// 		statements.NewUse("xyz"),
// 		statements.NewVar("x", "int"),
// 	}
// 	testFileParserSuccess(t, "just a class statement", file, expected, 2)
// }

func testFileParserSuccess(
	t *testing.T,
	name string,
	file io.Reader,
	expectedStmts []slang.Statement,
	expectedErrCount int,
) {
	t.Run(name, func(t *testing.T) {
		stmts, parseErrs := parsing.ParseFile(file)

		if !assert.Len(t, parseErrs, expectedErrCount) {
			logParseErrs(t, parseErrs)

			return
		}

		verifyStatemnts(t, expectedStmts, stmts)
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
