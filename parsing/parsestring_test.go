package parsing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/parsing"
)

func TestParseString(t *testing.T) {
	testParseStringSuccess(t, "just a string", expressions.NewString("just a string"))
	testParseStringSuccess(t, "string with {5} in the middle",
		expressions.NewString("string with "),
		expressions.NewInteger(5),
		expressions.NewString(" in the middle"),
	)
}

func testParseStringSuccess(t *testing.T, input string, expected ...slang.Expression) {
	exprs, err := parsing.ParseString(input)

	require.NoError(t, err)

	assert.True(t, slang.ExpressionsEqual(expected, exprs))
}
