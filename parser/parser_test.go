package parser_test

import (
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/expressions"
	"github.com/jamestunnell/slang/parser"
	"github.com/jamestunnell/slang/statements"
)

func TestParserExprStatement(t *testing.T) {
	testCases := map[string]slang.Statement{
		// plain values
		"x":     se(id("x")),
		"5":     se(expressions.NewInteger(5)),
		"25.7":  se(expressions.NewFloat(25.7)),
		"false": se(expressions.NewBool(false)),
		"true":  se(expressions.NewBool(true)),

		// prefix operators
		"-15":   se(expressions.NewNegative(expressions.NewInteger(15))),
		"!true": se(expressions.NewNot(expressions.NewBool(true))),

		// infix operators
		"a + b":  se(add(id("a"), id("b"))),
		"a - b":  se(sub(id("a"), id("b"))),
		"a * b":  se(mul(id("a"), id("b"))),
		"a / b":  se(div(id("a"), id("b"))),
		"a > b":  se(expressions.NewGreater(id("a"), id("b"))),
		"a < b":  se(expressions.NewLess(id("a"), id("b"))),
		"a == b": se(expressions.NewEqual(id("a"), id("b"))),
		"a != b": se(expressions.NewNotEqual(id("a"), id("b"))),

		// more elaborate expressions
		"10 + 7 - 3":    se(sub(add(i(10), i(7)), i(3))),
		"15 + 2 * 12":   se(add(i(15), mul(i(2), i(12)))),
		"6 * 6 - 3 * 3": se(sub(mul(i(6), i(6)), mul(i(3), i(3)))),

		// grouped expression
		"(15 + 2) * 12": se(mul(add(i(15), i(2)), i(12))),

		// func calls
		"sum(1,2,3)":     se(call(id("sum"), i(1), i(2), i(3))),
		"5 * sub(10, 5)": se(mul(i(5), call(id("sub"), i(10), i(5)))),

		// strings
		`"abc" + "123"`: se(add(str("abc"), str("123"))),

		// assignment
		`x = 12`: statements.NewAssign(id("x"), i(12)),

		// comment
		`x + 5 # this part should be ignored`: se(add(id("x"), i(5))),
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			testParser(t, input, expected)
		})
	}
}

func TestParserErrsAssignMissingValue(t *testing.T) {
	testParserErrs(t, "x = ")
}

func TestParserErrsAssignMissingEqual(t *testing.T) {
	testParserErrs(t, "x 5")
}

func TestParserGroupedExpr(t *testing.T) {
	s := statements.NewAssign(
		expressions.NewIdentifier("x"),
		expressions.NewAdd(
			expressions.NewInteger(4),
			expressions.NewIdentifier("y"),
		),
	)

	testParser(t, "x = (4 + y)", s)
}

func TestParserOneAssignStatement(t *testing.T) {
	identExpr := expressions.NewIdentifier("x")
	intExpr := expressions.NewInteger(5)
	assign := statements.NewAssign(identExpr, intExpr)

	testParser(t, "x = 5", assign)
}

func TestParserArrayLiteral(t *testing.T) {
	input := `x = [1, "abc", true]`

	a := expressions.NewInteger(1)
	b := expressions.NewString("abc")
	c := expressions.NewBool(true)
	ary := expressions.NewArray(a, b, c)
	x := expressions.NewIdentifier("x")
	s := statements.NewAssign(x, ary)

	testParser(t, input, s)
}

func TestParserArrayIndx(t *testing.T) {
	input := `ary[1+3]`

	ary := expressions.NewIdentifier("ary")
	add := expressions.NewAdd(
		expressions.NewInteger(1),
		expressions.NewInteger(3))
	idx := expressions.NewIndex(ary, add)
	s := statements.NewExpression(idx)

	testParser(t, input, s)
}

func TestParserThreeAssignStatements(t *testing.T) {
	input := `
	a = 77
	b = 100.0
	longer_name = 75.0 - 22.2`

	a := expressions.NewIdentifier("a")
	b := expressions.NewIdentifier("b")
	c := expressions.NewIdentifier("longer_name")
	aVal := expressions.NewInteger(77)
	bVal := expressions.NewFloat(100.0)
	cVal := expressions.NewSubtract(
		expressions.NewFloat(75.0),
		expressions.NewFloat(22.2))
	s1 := statements.NewAssign(a, aVal)
	s2 := statements.NewAssign(b, bVal)
	s3 := statements.NewAssign(c, cVal)

	testParser(t, input, s1, s2, s3)
}

func TestParserFunctionReturnStatement(t *testing.T) {
	l := expressions.NewFloat(12.77)
	r := expressions.NewIdentifier("num")
	add := expressions.NewAdd(l, r)
	ret := statements.NewReturn(add)
	fnParams := []*slang.Param{slang.NewParam("num", "Integer")}
	fnBody := statements.NewBlock(ret)
	fn := expressions.NewFunctionLiteral(fnParams, fnBody)
	assign := statements.NewAssign(expressions.NewIdentifier("x"), fn)

	testParser(t, "x = func(num Integer){return 12.77 + num}", assign)
}

func TestParserErrsIfMissingRBrace(t *testing.T) {
	const input = `if true {
		x = 2
	`

	testParserErrs(t, input)
}

func TestParserErrsIfMissingRBraceInline(t *testing.T) {
	const input = `if true { x = 2`

	testParserErrs(t, input)
}

func TestParserIfExpr(t *testing.T) {
	input := `y = if a == 2 {
		x + 10
	}`
	cond := expressions.NewEqual(
		expressions.NewIdentifier("a"),
		expressions.NewInteger(2))
	assign := statements.NewExpression(
		expressions.NewAdd(
			expressions.NewIdentifier("x"),
			expressions.NewInteger(10)))
	conseq := statements.NewBlock(assign)
	ifExpr := expressions.NewIf(cond, conseq)

	testParser(t, input, statements.NewAssign(id("y"), ifExpr))

	input += ` else {
		76
	}`

	altern := statements.NewBlock(
		statements.NewExpression(
			expressions.NewInteger(76),
		),
	)
	ifElseExpr := expressions.NewIfElse(cond, conseq, altern)

	testParser(t, input, statements.NewAssign(id("y"), ifElseExpr))
}

func TestParserIfEquivalents(t *testing.T) {
	inputs := []string{
		`y = if a == 2 { x + 10 }`,
		`y = if a == 2 {
			x + 10
		}`,
		`y = if a == 2 {
			x + 10	# comments are fine
		}`,
	}
	cond := expressions.NewEqual(
		expressions.NewIdentifier("a"),
		expressions.NewInteger(2))
	assign := statements.NewExpression(
		expressions.NewAdd(
			expressions.NewIdentifier("x"),
			expressions.NewInteger(10)))
	conseq := statements.NewBlock(assign)
	ifExpr := expressions.NewIf(cond, conseq)
	st := statements.NewAssign(id("y"), ifExpr)

	for _, input := range inputs {
		testParser(t, input, st)
	}
}

func TestParserFuncLiteralNoParams(t *testing.T) {
	const input = `myvar = func(){
		return 7
	}`

	body := statements.NewBlock(
		statements.NewReturn(expressions.NewInteger(7)),
	)
	af := expressions.NewFunctionLiteral([]*slang.Param{}, body)
	assign := statements.NewAssign(
		expressions.NewIdentifier("myvar"), af)

	testParser(t, input, assign)
}

func TestParserFuncLiteralOneParams(t *testing.T) {
	const input = `myvar = func(x){
		return 7
	}`

	body := statements.NewBlock(
		statements.NewReturn(expressions.NewInteger(7)),
	)
	params := []*slang.Param{slang.NewParam("x", "Integer")}
	af := expressions.NewFunctionLiteral(params, body)
	assign := statements.NewAssign(
		expressions.NewIdentifier("myvar"), af)

	testParser(t, input, assign)
}

func TestParserFuncLiteralTwoParams(t *testing.T) {
	const input = `myvar = func(x Integer, y Float){
		return 7
	}`

	body := statements.NewBlock(
		statements.NewReturn(expressions.NewInteger(7)),
	)
	params := []*slang.Param{
		slang.NewParam("x", "Integer"),
		slang.NewParam("y", "Float"),
	}
	af := expressions.NewFunctionLiteral(params, body)
	assign := statements.NewAssign(
		expressions.NewIdentifier("myvar"), af)

	testParser(t, input, assign)
}

func testParser(t *testing.T, input string, expected ...slang.Statement) {
	results := parser.Parse(input)

	for i, err := range results.Errors {
		log.Debug().
			Err(err.Error).
			Int("line", err.Token.Location.Line).
			Int("column", err.Token.Location.Column).
			Str("token", err.Token.Info.Value()).
			Msgf("parse error #%d", i+1)
	}

	require.Empty(t, results.Errors)

	require.Len(t, results.Statements, len(expected))

	for i := 0; i < len(results.Statements); i++ {
		s := results.Statements[i]

		assert.Equal(t, expected[i].Type(), s.Type())
		if !assert.True(t, s.Equal(expected[i])) {
			t.Logf("statements not equal: expected %#v, got %#v", expected[i], s)
		}
	}
}

func testParserErrs(t *testing.T, input string) {
	results := parser.Parse(input)

	for _, pErr := range results.Errors {
		t.Logf("parse error at %s: %v", pErr.Token.Location, pErr.Error)
	}

	require.NotEmpty(t, results.Errors)
}

func se(expr slang.Expression) slang.Statement {
	return statements.NewExpression(expr)
}

func id(name string) *expressions.Identifier {
	return expressions.NewIdentifier(name)
}

func add(left, right slang.Expression) slang.Expression {
	return expressions.NewAdd(left, right)
}

func sub(left, right slang.Expression) slang.Expression {
	return expressions.NewSubtract(left, right)
}

func mul(left, right slang.Expression) slang.Expression {
	return expressions.NewMultiply(left, right)
}

func div(left, right slang.Expression) slang.Expression {
	return expressions.NewDivide(left, right)
}

func i(val int64) slang.Expression {
	return expressions.NewInteger(val)
}

func str(val string) slang.Expression {
	return expressions.NewString(val)
}

func call(fn slang.Expression, args ...slang.Expression) slang.Expression {
	return expressions.NewFunctionCall(fn, args...)
}
