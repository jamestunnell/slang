package parsers_test

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/ast/types"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
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

		func GetX() (result int) {
			result = x
		}

		func GetY() (result int) {
			result = y
		}
	`)
	expected := []*statements.Statement{
		statements.NewUse("rand"),
		statements.NewVar("x", types.NewInt()),
		statements.NewVar("y", types.NewInt()),
		statements.NewFunc("init",
			[]slang.Param{},
			[]slang.Param{},
			statements.NewAssign(
				expressions.NewIdentifier("x"),
				expressions.NewInvoke(
					expressions.NewAccessMember(expressions.NewIdentifier("rand"), "Int"),
				),
			),
			statements.NewAssign(
				expressions.NewIdentifier("y"),
				expressions.NewInvoke(
					expressions.NewAccessMember(expressions.NewIdentifier("rand"), "Int"),
				),
			),
		),
		statements.NewFunc("GetX",
			[]slang.Param{},
			[]slang.Param{
				types.NewNameType("result", types.NewInt()),
			},
			statements.NewAssign(
				expressions.NewIdentifier("result"),
				expressions.NewIdentifier("x"),
			),
		),
		statements.NewFunc("GetY",
			[]slang.Param{},
			[]slang.Param{
				types.NewNameType("result", types.NewInt()),
			},
			statements.NewAssign(
				expressions.NewIdentifier("result"),
				expressions.NewIdentifier("y"),
			),
		),
	}
	testFileParserSuccess(t, "global vars and funcs", file, expected)
}

func TestFileParserGlobalConst(t *testing.T) {
	file := strings.NewReader(`
		const myConst = 25.7
		var x int
	`)
	expected := []*statements.Statement{
		statements.NewConst("myConst", expressions.NewFloat(25.7)),
		statements.NewVar("x", types.NewInt()),
	}
	testFileParserSuccess(t, "global const", file, expected)
}

func TestFileParserStructWithTest(t *testing.T) {
	file := strings.NewReader(`
		use "test"

		struct Accumulator(total flt)

		func Add(a Accumulator, x flt) {
			a.total = a.total + x
		}

		func Mul(a Accumulator, x flt) {
			a.total = a.total * x
		}

		func Total(a Accumulator) (result flt) {
			result = a.total
		}

		func TestAccumulator(t test.Test) {
			accum = Accumulator(0.0)
			
			accum.Add(2.0)
			accum.Mul(2.0)

			t.AssertAlmostEq(accum.Total(), 4.0)

			accum.Add(1.0)
			accum.Mul(0.5)

			t.AssertAlmostEq(accum.Total(), 2.5)
		}
	`)
	expected := []*statements.Statement{
		statements.NewUse("test"),
		statements.NewStruct("Accumulator",
			types.NewNameType("total", types.NewFlt())),
		statements.NewFunc("Add",
			[]slang.Param{
				types.NewNameType("a", types.NewStruct("", "Accumulator")),
				types.NewNameType("x", types.NewFlt()),
			},
			[]slang.Param{},
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
		statements.NewFunc("Mul",
			[]slang.Param{
				types.NewNameType("a", types.NewStruct("", "Accumulator")),
				types.NewNameType("x", types.NewFlt()),
			},
			[]slang.Param{},
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
		statements.NewFunc("Total",
			[]slang.Param{
				types.NewNameType("a", types.NewStruct("", "Accumulator")),
			},
			[]slang.Param{
				types.NewNameType("result", types.NewFlt()),
			},
			statements.NewAssign(
				expressions.NewIdentifier("result"),
				expressions.NewAccessMember(
					expressions.NewIdentifier("a"), "total"),
			),
		),
		statements.NewFunc("TestAccumulator",
			[]slang.Param{types.NewNameType("t", types.NewStruct("test", "Test"))},
			[]slang.Param{},
			statements.NewAssign(
				expressions.NewIdentifier("accum"),
				expressions.NewInvoke(
					expressions.NewIdentifier("Accumulator"),
					expressions.NewInvokeArgPos(expressions.NewFloat(0.0)),
				),
			),
			statements.NewExpression(
				expressions.NewInvoke(
					expressions.NewAccessMember(expressions.NewIdentifier("accum"), "Add"),
					expressions.NewInvokeArgPos(expressions.NewFloat(2.0)),
				),
			),
			statements.NewExpression(
				expressions.NewInvoke(
					expressions.NewAccessMember(expressions.NewIdentifier("accum"), "Mul"),
					expressions.NewInvokeArgPos(expressions.NewFloat(2.0)),
				),
			),
			statements.NewExpression(
				expressions.NewInvoke(
					expressions.NewAccessMember(expressions.NewIdentifier("t"), "AssertAlmostEq"),
					expressions.NewInvokeArgPos(expressions.NewInvoke(
						expressions.NewAccessMember(expressions.NewIdentifier("accum"), "Total"),
					)),
					expressions.NewInvokeArgPos(expressions.NewFloat(4.0)),
				),
			),
			statements.NewExpression(
				expressions.NewInvoke(
					expressions.NewAccessMember(expressions.NewIdentifier("accum"), "Add"),
					expressions.NewInvokeArgPos(expressions.NewFloat(1.0)),
				),
			),
			statements.NewExpression(
				expressions.NewInvoke(
					expressions.NewAccessMember(expressions.NewIdentifier("accum"), "Mul"),
					expressions.NewInvokeArgPos(expressions.NewFloat(0.5)),
				),
			),
			statements.NewExpression(
				expressions.NewInvoke(
					expressions.NewAccessMember(expressions.NewIdentifier("t"), "AssertAlmostEq"),
					expressions.NewInvokeArgPos(expressions.NewInvoke(
						expressions.NewAccessMember(expressions.NewIdentifier("accum"), "Total"),
					)),
					expressions.NewInvokeArgPos(expressions.NewFloat(2.5)),
				),
			),
		),
	}
	testFileParserSuccess(t, "struct with test", file, expected)
}

func TestFileParserStructOneline(t *testing.T) {
	file := strings.NewReader(`struct X (a, b int, c str)`)
	expected := []*statements.Statement{
		statements.NewStruct("X",
			types.NewNameType("a", types.NewInt()),
			types.NewNameType("b", types.NewInt()),
			types.NewNameType("c", types.NewStr()),
		),
	}
	testFileParserSuccess(t, "struct multiline", file, expected)
}

func TestFileParserStructMultiline(t *testing.T) {
	file := strings.NewReader(`
		struct X (
		  a, b int
		  c str
		)
	`)
	expected := []*statements.Statement{
		statements.NewStruct("X",
			types.NewNameType("a", types.NewInt()),
			types.NewNameType("b", types.NewInt()),
			types.NewNameType("c", types.NewStr()),
		),
	}
	testFileParserSuccess(t, "struct multiline", file, expected)
}

func TestFileParserWithComments(t *testing.T) {
	file := strings.NewReader(`
		// this is a leading
		// standalone comment
		
		// my struct comment
		struct X (
			a, b int
			c str
		)

		// this is a
		// standalone comment

		// also not empty
		const y = 10

		// this is a trailing
		// standalone comment
	`)
	expected := []*statements.Statement{
		withComment(statements.NewComment(), "this is a leading standalone comment"),
		withComment(
			statements.NewStruct("X",
				types.NewNameType("a", types.NewInt()),
				types.NewNameType("b", types.NewInt()),
				types.NewNameType("c", types.NewStr()),
			),
			"my struct comment",
		),
		withComment(statements.NewComment(), "this is a standalone comment"),
		withComment(statements.NewConst("y", expressions.NewInt(10)), "also not empty"),
		withComment(statements.NewComment(), "this is a trailing standalone comment"),
	}

	testFileParserSuccess(t, "with comments", file, expected)
}

func testFileParserSuccess(
	t *testing.T,
	name string,
	file io.Reader,
	expectedStmts []*statements.Statement,
) {
	t.Run(name, func(t *testing.T) {
		p := parsers.NewFileParser()

		err := parsing.RunParser(p, file)

		assert.NoError(t, err)

		verifyStatemnts(t, expectedStmts, p.Statements)
	})
}

func verifyStatemnts(t *testing.T, expected, actual []*statements.Statement) {
	if !assert.Equal(t, len(expected), len(actual)) {
		return
	}

	for i, stmt := range expected {
		if !assert.True(t, stmt.IsEqual(actual[i])) {
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

func withComment(s *statements.Statement, comment string) *statements.Statement {
	s.SetComment(comment)

	return s
}
