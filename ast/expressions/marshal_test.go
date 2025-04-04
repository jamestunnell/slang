package expressions_test

import (
	"encoding/json"
	"testing"

	"github.com/jamestunnell/slang"
	e "github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestMarshalJSON(t *testing.T) {
	a := e.NewIdentifier("a")
	b := e.NewIdentifier("b")

	// testMarshalJSON(t, e.NewAccessElem(a, b))
	testMarshalJSON(t, e.NewAccessMember(a, "myMember"))
	testMarshalJSON(t, e.NewAdd(a, b))
	// testMarshalJSON(t, e.NewArray(types.NewInt(), a, b))
	testMarshalJSON(t, e.NewBool(true))
	testMarshalJSON(t, e.NewDivide(a, b))
	testMarshalJSON(t, e.NewEqual(a, b))
	testMarshalJSON(t, e.NewFloat(0.0))
	testMarshalJSON(t, e.NewFunc(
		[]*types.NameType{types.NewNameType("x", types.NewInt())},
		[]*types.NameType{types.NewNameType("result", types.NewBool())},
	))
	testMarshalJSON(t, e.NewInvoke(a, e.NewInvokeArgPos(e.NewInt(10))))
	testMarshalJSON(t, e.NewInvoke(a, e.NewInvokeArgKW("b", e.NewInt(10))))
	testMarshalJSON(t, e.NewGreater(a, b))
	testMarshalJSON(t, e.NewGreaterEqual(a, b))
	testMarshalJSON(t, e.NewIdentifier("x"))
	testMarshalJSON(t, e.NewInt(0))
	testMarshalJSON(t, e.NewLess(a, b))
	testMarshalJSON(t, e.NewLessEqual(a, b))

	testMarshalJSON(t, e.NewMultiply(a, b))
	testMarshalJSON(t, e.NewNegative(a))
	testMarshalJSON(t, e.NewNot(a))
	testMarshalJSON(t, e.NewNotEqual(a, b))
	testMarshalJSON(t, e.NewStr("hello"))
	testMarshalJSON(t, e.NewStruct(
		types.NewNameType("x", types.NewInt()),
		types.NewNameType("y", types.NewStr()),
	))
	testMarshalJSON(t, e.NewSubtract(a, b))
}

func testMarshalJSON(t *testing.T, expr slang.Expression) {
	t.Run(expr.GetType().String(), func(t *testing.T) {
		d, err := json.Marshal(expr)

		require.NoError(t, err)

		result := gjson.GetBytes(d, "type")

		require.True(t, result.Exists())
		require.Equal(t, gjson.String, result.Type)

		assert.Equal(t, expr.GetType().String(), result.String())
	})
}
