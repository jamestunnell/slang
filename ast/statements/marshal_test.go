package statements_test

import (
	"encoding/json"
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestMarshalJSON(t *testing.T) {
	testMarshalJSON(t, statements.NewAssign(expressions.NewIdentifier("xyz"), expressions.NewInteger(5)))
	testMarshalJSON(t, statements.NewClass("myclass", ""))
	testMarshalJSON(t, statements.NewField("myfield", ast.NewBasicType("int")))
	testMarshalJSON(t, statements.NewMethod("mymethod", ast.NewFunction([]slang.Param{}, []slang.Type{})))
	testMarshalJSON(t, statements.NewReturnVal(expressions.NewInteger(7)))
	testMarshalJSON(t, statements.NewUse("my/path"))
}

func testMarshalJSON(t *testing.T, stmt slang.Statement) {
	t.Run(stmt.Type().String(), func(t *testing.T) {
		d, err := json.Marshal(stmt)

		require.NoError(t, err)

		result := gjson.GetBytes(d, "type")

		require.True(t, result.Exists())
		require.Equal(t, gjson.String, result.Type)

		assert.Equal(t, stmt.Type().String(), result.String())
	})
}
