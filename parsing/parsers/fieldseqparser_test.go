package parsers_test

import (
	"testing"

	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/ast/types"
	"github.com/jamestunnell/slang/parsing/parsers"
)

func TestFieldSeqParser_NoInput(t *testing.T) {
	testFieldSeqParserFail(t, "")
}

func TestFieldSeqParser_EmptySeq(t *testing.T) {
	testFieldSeqParserSuccess(t, "()")
	testFieldSeqParserSuccess(t, `(
	
	)`)
}

func TestFieldSeqParser_OneField(t *testing.T) {
	stmts := []*statements.Statement{
		statements.NewField(types.NewFlt(), "x"),
	}

	// testFieldSeqParserSuccess(t, "(x float)", stmts...)
	testFieldSeqParserSuccess(t, `(
			x float
		)`, stmts...)
}

func TestFieldSeqParser_TwoFields(t *testing.T) {

	testFieldSeqParserFail(t, "(a int b mymodule.MyType)")
	testFieldSeqParserSuccess(t, `(
			a int
			b mymodule.MyType
		)`,
		statements.NewField(types.NewInt(), "a"),
		statements.NewField(types.NewStruct("mymodule", "MyType"), "b"),
	)
}

func TestFieldSeqParser_TwoMultiName(t *testing.T) {
	// testFieldSeqParserFail(t, "(a, b int c, d str)")
	testFieldSeqParserSuccess(t, `(
			a, b int
			c, d str
		)`,
		statements.NewField(types.NewInt(), "a", "b"),
		statements.NewField(types.NewStr(), "c", "d"),
	)
}

func TestFieldSeqParser_CommentedFields(t *testing.T) {
	testBodyParserSuccess(t, parsers.NewFieldSeqParser(), `(
		// a comment
		a int
	)`, withComment(statements.NewField(types.NewInt(), "a"), "a comment"))
}

func testFieldSeqParserSuccess(
	t *testing.T,
	input string,
	expected ...*statements.Statement,
) {
	testBodyParserSuccess(t, parsers.NewFieldSeqParser(), input, expected...)
}

func testFieldSeqParserFail(
	t *testing.T,
	input string,
) {
	p := parsers.NewFieldSeqParser()

	testBodyParserFail(t, p, input)
}
