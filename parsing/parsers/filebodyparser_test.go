package parsers_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/field"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/ast/types"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
	"github.com/stretchr/testify/assert"
)

func TestFileParserGlobalVars(t *testing.T) {
	const input = `
		use "rand"

		var x 0
		var y 0

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
	`

	expected := []*statements.Statement{
		statements.NewUse("", []string{"rand"}),
		statements.NewVar("x", expressions.NewInt(0)),
		statements.NewVar("y", expressions.NewInt(0)),
		statements.NewFunc("init",
			[]*field.Field{},
			[]*field.Field{},
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
			[]*field.Field{},
			[]*field.Field{
				field.New("result", types.NewInt()),
			},
			statements.NewAssign(
				expressions.NewIdentifier("result"),
				expressions.NewIdentifier("x"),
			),
		),
		statements.NewFunc("GetY",
			[]*field.Field{},
			[]*field.Field{
				field.New("result", types.NewInt()),
			},
			statements.NewAssign(
				expressions.NewIdentifier("result"),
				expressions.NewIdentifier("y"),
			),
		),
	}

	testFileParserSuccess(t, input, expected)
}

func TestFileParserGlobalConst(t *testing.T) {
	const input = `
		const myConst 25.7
		var x 12
	`

	expected := []*statements.Statement{
		statements.NewConst("myConst", expressions.NewFloat(25.7)),
		statements.NewVar("x", expressions.NewInt(12)),
	}

	testFileParserSuccess(t, input, expected)
}

func TestFileParserStructWithTest(t *testing.T) {
	const input = `
		use "test"

		struct Accumulator(total flt)

		func Add(
			a Accumulator
			x flt) {
			a.total = a.total + x
		}

		func Mul(
			a Accumulator
			x flt) {
			a.total = a.total * x
		}

		func Total(a Accumulator) (result flt) {
			result = a.total
		}

		func TestAccumulator(t test.Test) {
			accum = Accumulator(0.0)
			
			accum.Add(2.0)
			accum.Mul(2.0)

			t.AssertAlmostEq(accum.Total() 4.0)

			accum.Add(1.0)
			accum.Mul(0.5)

			t.AssertAlmostEq(accum.Total() 2.5)
		}
	`

	expected := []*statements.Statement{
		statements.NewUse("", []string{"test"}),
		statements.NewStruct("Accumulator",
			field.New("total", types.NewFlt())),
		statements.NewFunc("Add",
			[]*field.Field{
				field.New("a", types.NewStruct("", "Accumulator")),
				field.New("x", types.NewFlt()),
			},
			[]*field.Field{},
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
			[]*field.Field{
				field.New("a", types.NewStruct("", "Accumulator")),
				field.New("x", types.NewFlt()),
			},
			[]*field.Field{},
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
			[]*field.Field{
				field.New("a", types.NewStruct("", "Accumulator")),
			},
			[]*field.Field{
				field.New("result", types.NewFlt()),
			},
			statements.NewAssign(
				expressions.NewIdentifier("result"),
				expressions.NewAccessMember(
					expressions.NewIdentifier("a"), "total"),
			),
		),
		statements.NewFunc("TestAccumulator",
			[]*field.Field{field.New("t", types.NewStruct("test", "Test"))},
			[]*field.Field{},
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

	testFileParserSuccess(t, input, expected)
}

func TestFileParserStructOneLine(t *testing.T) {
	testFileParserFail(t, `struct X (a int b str)`)
}

func TestFileParserStructMultiline(t *testing.T) {
	const input = `
		struct X (
		  a, b int
		  c str
		)
	`

	expected := []*statements.Statement{
		statements.NewStruct("X",
			field.New("a", types.NewInt()),
			field.New("b", types.NewInt()),
			field.New("c", types.NewStr()),
		),
	}
	testFileParserSuccess(t, input, expected)
}

func TestFileParserWithComments(t *testing.T) {
	const input = `
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
		const y 10

		// this is a trailing
		// standalone comment
	`

	expected := []*statements.Statement{
		withComment(statements.NewComment(), "this is a leading standalone comment"),
		withComment(
			statements.NewStruct("X",
				field.New("a", types.NewInt()),
				field.New("b", types.NewInt()),
				field.New("c", types.NewStr()),
			),
			"my struct comment",
		),
		withComment(statements.NewComment(), "this is a standalone comment"),
		withComment(statements.NewConst("y", expressions.NewInt(10)), "also not empty"),
		withComment(statements.NewComment(), "this is a trailing standalone comment"),
	}

	testFileParserSuccess(t, input, expected)
}

func testFileParserSuccess(
	t *testing.T,
	input string,
	expectedStmts []*statements.Statement,
) {
	p := parsers.NewFileParser()
	err := parsing.RunParser(p, strings.NewReader(input))

	assert.NoError(t, err)

	verifyStatements(t, expectedStmts, p.GetStatements())
}

func testFileParserFail(t *testing.T, input string) {
	p := parsers.NewFileParser()
	err := parsing.RunParser(p, strings.NewReader(input))

	assert.Error(t, err)
}

func verifyStatements(t *testing.T, expected, actual []*statements.Statement) {
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
		t.Logf("unxpected parse err: %v", parseErr)
	}
}

func withComment(s *statements.Statement, comment string) *statements.Statement {
	s.SetComment(comment)

	return s
}
