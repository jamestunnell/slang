package statements_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/field"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/ast/types"
)

func TestMarshalJSON(t *testing.T) {
	testMarshalJSON(t, statements.NewAssign(expressions.NewIdentifier("xyz"), expressions.NewInt(5)))
	testMarshalJSON(t, statements.NewStruct("myclass",
		field.New("x", types.NewInt()),
		field.New("y", types.NewStr()),
		field.New("z", types.NewStr()),
	))
	testMarshalJSON(t, statements.NewFunc("myfunc", []*field.Field{}, []*field.Field{}))
	testMarshalJSON(t, statements.NewUse("", []string{"my", "path"}))
}

func testMarshalJSON(t *testing.T, stmt slang.Statement) {
	t.Run(stmt.GetType().String(), func(t *testing.T) {
		d, err := json.Marshal(stmt)

		require.NoError(t, err)

		var stmt2 statements.Statement

		err = json.Unmarshal(d, &stmt2)

		if assert.NoError(t, err) {
			assert.True(t, stmt2.IsEqual(stmt))
		}
	})
}
