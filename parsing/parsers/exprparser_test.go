package parsing_test

import (
	"strings"
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/parsing"
	"github.com/stretchr/testify/assert"
)

func TestExprParser(t *testing.T) {
	testCases := map[string]*expressions.Expression{
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
		"sum(1,2,3)":     invokePos(id("sum"), i(1), i(2), i(3)),
		"5 * sub(10, 5)": mul(i(5), invokePos(id("sub"), i(10), i(5))),

		// strings
		`"abc" + "123"`: add(str("abc"), str("123")),

		// string interpolation
		`"${x} is a ${y}"`: expressions.NewConcat(str(""), id("x"), str(" is a "), id("y"), str("")),

		// // map literal value
		// `[string]int{"a": 1, "b": 2}`: m(
		// 	types.NewStr(),
		// 	exprs(str("a"), str("b")),
		// 	types.NewInt(),
		// 	exprs(i(1), i(2)),
		// ),

		// // nested map values
		// `[string][string]int{"a": [string]int{"b": 2}}`: m(
		// 	types.NewStr(),
		// 	exprs(str("a")),
		// 	ast.NewMapType(types.NewStr(), types.NewInt()),
		// 	exprs(m(types.NewStr(), exprs(str("b")), types.NewInt(), exprs(i(2)))),
		// ),

		// // array literal value
		// `[]int{1, 2, 3}`: ary(types.NewInt(), i(1), i(2), i(3)),

		// // nested array values
		// `[][]string{[]string{"a", "b", "c"}, []string{"x", "y", "z"}}`: ary(
		// 	ast.NewArrayType(types.NewStr()),
		// 	ary(types.NewStr(), str("a"), str("b"), str("c")),
		// 	ary(types.NewStr(), str("x"), str("y"), str("z")),
		// ),

		// // access map/array element
		// `myContainer[myKey]`: elem(id("myContainer"), id("myKey")),
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			testExprParser(t, input, expected)
		})
	}
}

func testExprParser(t *testing.T, input string, expected *expressions.Expression) {
	t.Run(input, func(t *testing.T) {
		l := lexing.NewLexer(strings.NewReader(input))
		toks := parsing.NewTokenSeq(l)
		p := parsing.NewExprParser(parsing.PrecedenceLOWEST)

		assert.True(t, p.Run(toks))

		if !assert.Empty(t, p.GetErrors()) {
			logParseErrs(t, p.GetErrors())

			t.FailNow()
		}

		if !assert.True(t, p.Expr.IsEqual(expected)) {
			t.Logf("%#v (actual) != %#v (expected)", p.Expr, expected)
		}
	})
}

func id(name string) *expressions.Expression {
	return expressions.NewIdentifier(name)
}

func add(left, right *expressions.Expression) *expressions.Expression {
	return expressions.NewAdd(left, right)
}

func sub(left, right *expressions.Expression) *expressions.Expression {
	return expressions.NewSubtract(left, right)
}

func mul(left, right *expressions.Expression) *expressions.Expression {
	return expressions.NewMultiply(left, right)
}

func div(left, right *expressions.Expression) *expressions.Expression {
	return expressions.NewDivide(left, right)
}

func i(val int64) *expressions.Expression {
	return expressions.NewInt(val)
}

func b(val bool) *expressions.Expression {
	return expressions.NewBool(val)
}

func f(val float64) *expressions.Expression {
	return expressions.NewFloat(val)
}

func str(val string) *expressions.Expression {
	return expressions.NewStr(val)
}

func invokePos(fn slang.Expression, argVals ...*expressions.Expression) *expressions.Expression {
	args := make([]*expressions.InvokeArg, len(argVals))

	for i, val := range argVals {
		args[i] = expressions.NewInvokeArgPos(val)
	}

	return expressions.NewInvoke(fn, args...)
}

func gt(left, right *expressions.Expression) *expressions.Expression {
	return expressions.NewGreater(left, right)
}

func lt(left, right *expressions.Expression) *expressions.Expression {
	return expressions.NewLess(left, right)
}

func and(left, right *expressions.Expression) *expressions.Expression {
	return expressions.NewAnd(left, right)
}

func or(left, right *expressions.Expression) *expressions.Expression {
	return expressions.NewOr(left, right)
}

func not(val *expressions.Expression) *expressions.Expression {
	return expressions.NewNot(val)
}

// func ary(valType slang.Type, vals ...*expressions.Expression) slang.Expression {
// 	return expressions.NewArray(valType, vals...)
// }

func exprs(vals ...*expressions.Expression) []*expressions.Expression {
	return vals
}

// func m(keyType slang.Type, keys []slang.Expression,
// 	valType slang.Type, vals []*expressions.Expression) slang.Expression {
// 	return expressions.NewMap(keyType, keys, valType, vals)
// }

// func elem(a, b *expressions.Expression) slang.Expression {
// 	return expressions.NewAccessElem(a, b)
// }
