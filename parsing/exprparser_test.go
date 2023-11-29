package parsing_test

import (
	"strings"
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexer"
	"github.com/jamestunnell/slang/parsing"
	"github.com/stretchr/testify/assert"
)

func TestExprParser(t *testing.T) {
	testCases := map[string]slang.Expression{
		// plain values
		"x":     id("x"),
		"5":     i(5),
		"25.7":  f(25.7),
		"false": b(false),
		"true":  b(true),

		// prefix operators
		"-15":   expressions.NewNegative(i(15)),
		"!true": expressions.NewNot(b(true)),

		// infix operators
		"a + b":  add(id("a"), id("b")),
		"a - b":  sub(id("a"), id("b")),
		"a * b":  mul(id("a"), id("b")),
		"a / b":  div(id("a"), id("b")),
		"a > b":  gt(id("a"), id("b")),
		"a < b":  expressions.NewLess(id("a"), id("b")),
		"a == b": expressions.NewEqual(id("a"), id("b")),
		"a != b": expressions.NewNotEqual(id("a"), id("b")),

		// more elaborate expressions
		"10 + 7 - 3":    sub(add(i(10), i(7)), i(3)),
		"15 + 2 * 12":   add(i(15), mul(i(2), i(12))),
		"6 * 6 - 3 * 3": sub(mul(i(6), i(6)), mul(i(3), i(3))),

		// grouped expression
		"(15 + 2) * 12": mul(add(i(15), i(2)), i(12)),

		// func calls
		"sum(1,2,3)":     callPosArgsOnly(id("sum"), i(1), i(2), i(3)),
		"5 * sub(10, 5)": mul(i(5), callPosArgsOnly(id("sub"), i(10), i(5))),

		// strings
		`"abc" + "123"`: add(str("abc"), str("123")),
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			testExprParser(t, input, expected)
		})
	}
}

func TestExprParserIf(t *testing.T) {
	input := `if a > b {
			return 5
		}`

	cond := gt(id("a"), id("b"))
	conseqs := []slang.Statement{statements.NewReturn(i(5))}

	testExprParser(t, input, expressions.NewIf(cond, conseqs))
}

func testExprParser(t *testing.T, input string, expected slang.Expression) {
	t.Run(input, func(t *testing.T) {
		l := lexer.New(strings.NewReader(input))
		toks := parsing.NewTokenSeq(l)
		p := parsing.NewExprParser(parsing.PrecedenceLOWEST)

		p.Run(toks)

		if !assert.Empty(t, p.GetErrors()) {
			logParseErrs(t, p.GetErrors())

			t.FailNow()
		}

		assert.True(t, p.Expr.Equal(expected))
	})
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

func b(val bool) slang.Expression {
	return expressions.NewBool(val)
}

func f(val float64) slang.Expression {
	return expressions.NewFloat(val)
}

func str(val string) slang.Expression {
	return expressions.NewString(val)
}

func callPosArgsOnly(fn slang.Expression, args ...slang.Expression) slang.Expression {
	return expressions.NewCall(fn, args, map[string]slang.Expression{})
}

func gt(left, right slang.Expression) slang.Expression {
	return expressions.NewGreater(left, right)
}
