package parsing_test

import (
	"strings"
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
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
		"!true": not(b(true)),

		// infix operators
		"a + b":   add(id("a"), id("b")),
		"a - b":   sub(id("a"), id("b")),
		"a * b":   mul(id("a"), id("b")),
		"a / b":   div(id("a"), id("b")),
		"a > b":   gt(id("a"), id("b")),
		"a < b":   lt(id("a"), id("b")),
		"a == b":  expressions.NewEqual(id("a"), id("b")),
		"a != b":  expressions.NewNotEqual(id("a"), id("b")),
		"a or b":  or(id("a"), id("b")),
		"a and b": and(id("a"), id("b")),

		// more elaborate expressions
		"10 + 7 - 3":         sub(add(i(10), i(7)), i(3)),
		"15 + 2 * 12":        add(i(15), mul(i(2), i(12))),
		"6 * 6 - 3 * 3":      sub(mul(i(6), i(6)), mul(i(3), i(3))),
		"12 < 7 and a or !b": or(and(lt(i(12), i(7)), id("a")), not(id("b"))),

		// grouped expression
		"(15 + 2) * 12": mul(add(i(15), i(2)), i(12)),

		// func calls
		"sum(1,2,3)":     callPosArgsOnly(id("sum"), i(1), i(2), i(3)),
		"5 * sub(10, 5)": mul(i(5), callPosArgsOnly(id("sub"), i(10), i(5))),

		// strings
		`"abc" + "123"`: add(str("abc"), str("123")),

		// string interpolation
		`"${x} is a ${y}"`: expressions.NewConcat(str(""), id("x"), str(" is a "), id("y"), str("")),

		// map value
		`{"a": 1, "b": 2}`: m(exprs(str("a"), str("b")), exprs(i(1), i(2))),

		// nested map values
		`{"a": {"b": 2}}`: m(exprs(str("a")), exprs(m(exprs(str("b")), exprs(i(2))))),

		// access map with key
		`map{key}`: expressions.NewKey(id("map"), id("key")),

		// array value
		`[1, 2, 3]`: ary(i(1), i(2), i(3)),

		// nested array values
		`["abc", ["xyz"]]`: ary(str("abc"), ary(str("xyz"))),

		// access array with index
		`ary[idx]`: expressions.NewIndex(id("ary"), id("idx")),
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			testExprParser(t, input, expected)
		})
	}
}

func testExprParser(t *testing.T, input string, expected slang.Expression) {
	t.Run(input, func(t *testing.T) {
		l := lexer.New(strings.NewReader(input))
		toks := parsing.NewTokenSeq(l)
		p := parsing.NewExprParser(parsing.PrecedenceLOWEST)

		assert.True(t, p.Run(toks))

		if !assert.Empty(t, p.GetErrors()) {
			logParseErrs(t, p.GetErrors())

			t.FailNow()
		}

		if !assert.True(t, p.Expr.Equal(expected)) {
			t.Logf("%#v (actual) != %#v (expected)", p.Expr, expected)
		}
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

func callPosArgsOnly(fn slang.Expression, argVals ...slang.Expression) slang.Expression {
	args := make([]*expressions.Arg, len(argVals))

	for i, val := range argVals {
		args[i] = expressions.NewPositionalArg(val)
	}

	return expressions.NewCall(fn, args...)
}

func gt(left, right slang.Expression) slang.Expression {
	return expressions.NewGreater(left, right)
}

func lt(left, right slang.Expression) slang.Expression {
	return expressions.NewLess(left, right)
}

func and(left, right slang.Expression) slang.Expression {
	return expressions.NewAnd(left, right)
}

func or(left, right slang.Expression) slang.Expression {
	return expressions.NewOr(left, right)
}

func not(val slang.Expression) slang.Expression {
	return expressions.NewNot(val)
}

func ary(vals ...slang.Expression) slang.Expression {
	return expressions.NewArray(vals...)
}

func exprs(vals ...slang.Expression) []slang.Expression {
	return vals
}

func m(keys, vals []slang.Expression) slang.Expression {
	return expressions.NewMap(keys, vals)
}
