package expressions_test

import (
	"encoding/json"
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestMarshalJSON(t *testing.T) {
	a := expressions.NewIdentifier("a")
	b := expressions.NewIdentifier("b")

	testMarshalJSON(t, expressions.NewAccessElem(a, b))
	testMarshalJSON(t, expressions.NewAccessMember(a, "myMember"))
	testMarshalJSON(t, expressions.NewAdd(a, b))
	testMarshalJSON(t, expressions.NewArray(ast.NewBasicType("int"), a, b))
	testMarshalJSON(t, expressions.NewBool(true))
	testMarshalJSON(t, expressions.NewDivide(a, b))
	testMarshalJSON(t, expressions.NewEqual(a, b))
	testMarshalJSON(t, expressions.NewFloat(0.0))
	testMarshalJSON(t, expressions.NewFunc(
		ast.NewFunction([]slang.Param{}, []slang.Type{ast.NewBasicType("bool")})))
	testMarshalJSON(t, expressions.NewCall(a))
	testMarshalJSON(t, expressions.NewGreater(a, b))
	testMarshalJSON(t, expressions.NewGreaterEqual(a, b))
	testMarshalJSON(t, expressions.NewIdentifier("x"))
	testMarshalJSON(t, expressions.NewInteger(0))
	testMarshalJSON(t, expressions.NewLess(a, b))
	testMarshalJSON(t, expressions.NewLessEqual(a, b))

	testMarshalJSON(t, expressions.NewMultiply(a, b))
	testMarshalJSON(t, expressions.NewNegative(a))
	testMarshalJSON(t, expressions.NewNot(a))
	testMarshalJSON(t, expressions.NewNotEqual(a, b))
	testMarshalJSON(t, expressions.NewString("hello"))
	testMarshalJSON(t, expressions.NewSubtract(a, b))
}

func testMarshalJSON(t *testing.T, expr slang.Expression) {
	t.Run(expr.Type().String(), func(t *testing.T) {
		d, err := json.Marshal(expr)

		require.NoError(t, err)

		result := gjson.GetBytes(d, "type")

		require.True(t, result.Exists())
		require.Equal(t, gjson.String, result.Type)

		assert.Equal(t, expr.Type().String(), result.String())
	})
}
