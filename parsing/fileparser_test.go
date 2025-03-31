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
		statements.NewFunc("init",
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
		),
		statements.NewFunc("GetX",
			[]slang.Param{},
			[]slang.Type{ast.NewBasicType("int")},
			statements.NewReturnVal(
				expressions.NewIdentifier("x"),
			),
		),
		statements.NewFunc("GetY",
			[]slang.Param{},
			[]slang.Type{ast.NewBasicType("int")},
			statements.NewReturnVal(
				expressions.NewIdentifier("y"),
			),
		),
	}
	testFileParserSuccess(t, "global vars and funcs", file, expected, 0)
}

func TestFileParserGlobalConst(t *testing.T) {
	file := strings.NewReader(`
		const myConst = 25.7
		var x int
	`)
	expected := []slang.Statement{
		statements.NewConst("myConst", expressions.NewFloat(25.7)),
		statements.NewVar("x", ast.NewBasicType("int")),
	}
	testFileParserSuccess(t, "global const", file, expected, 0)
}

func TestFileParserStructWithTest(t *testing.T) {
	file := strings.NewReader(`
		use "test"

		struct Accumulator {
			total float
		}

		func Add(a Accumulator, x float) {
			a.total = a.total + x
		}

		func Mul(a Accumulator, x float) {
			a.total = a.total * x
		}

		func Total(a Accumulator) float {
			return a.total
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
		statements.NewStruct("Accumulator",
			statements.NewStructField([]string{"total"}, ast.NewBasicType("float"))),
		statements.NewFunc("Add",
			[]slang.Param{
				ast.NewParam("a", ast.NewBasicType("Accumulator")),
				ast.NewParam("x", ast.NewBasicType("float")),
			},
			[]slang.Type{},
			statements.NewAssign(
				expressions.NewAccessMember(
					expressions.NewIdentifier("a"), "total"),
				expressions.NewAdd(
					expressions.NewAccessMember(
						expressions.NewIdentifier("a"), "total"),
					expressions.NewIdentifier("x"),
				),
			),
		),
		statements.NewFunc(
			"Mul",
			[]slang.Param{
				ast.NewParam("a", ast.NewBasicType("Accumulator")),
				ast.NewParam("x", ast.NewBasicType("float")),
			},
			[]slang.Type{},
			statements.NewAssign(
				expressions.NewAccessMember(
					expressions.NewIdentifier("a"), "total"),
				expressions.NewMultiply(
					expressions.NewAccessMember(
						expressions.NewIdentifier("a"), "total"),
					expressions.NewIdentifier("x"),
				),
			),
		),
		statements.NewFunc(
			"Total",
			[]slang.Param{
				ast.NewParam("a", ast.NewBasicType("Accumulator")),
			},
			[]slang.Type{ast.NewBasicType("float")},
			statements.NewReturnVal(
				expressions.NewAccessMember(
					expressions.NewIdentifier("a"), "total"),
			),
		),
		statements.NewFunc("TestAccumulator",
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
		),
	}
	testFileParserSuccess(t, "struct with test", file, expected, 0)
}

func TestFileParserStruct(t *testing.T) {
	file := strings.NewReader(`
		struct X {
		  a, b int
		  c string
		}
	`)
	expected := []slang.Statement{
		statements.NewStruct("X",
			statements.NewStructField([]string{"a", "b"}, ast.NewBasicType("int")),
			statements.NewStructField([]string{"c"}, ast.NewBasicType("string")),
		),
	}
	testFileParserSuccess(t, "struct", file, expected, 0)
}

func TestFileParserWithComments(t *testing.T) {
	file := strings.NewReader(`
		// this is a leading
		// standalone comment
		
		// my struct comment
		struct X {
			a, b int
			c string
		}

		// this is a
		// standalone comment

		// also not empty
		const y = 10

		// this is a trailing
		// standalone comment
	`)
	expected := []slang.Statement{
		statements.NewComment("this is a leading", "standalone comment"),
		statements.WithComment(
			statements.NewStruct("X",
				statements.NewStructField([]string{"a", "b"}, ast.NewBasicType("int")),
				statements.NewStructField([]string{"c"}, ast.NewBasicType("string")),
			),
			"my struct comment",
		),
		statements.NewComment("this is a", "standalone comment"),
		statements.WithComment(
			statements.NewConst("y", expressions.NewInteger(10)),
			"also not empty",
		),
		statements.NewComment("this is a trailing", "standalone comment"),
	}

	testFileParserSuccess(t, "with comments", file, expected, 0)
}

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
