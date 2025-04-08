package lexing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/tokens"
)

func TestLexer_Example1(t *testing.T) {
	input := `Person struct {
		names []str
		birthdate date
	}`
	expected := []*slang.Token{
		tok(tokens.SYMBOL("Person"), 1, 1),
		tok(tokens.STRUCT(), 1, 8),
		tok(tokens.LBRACE(), 1, 15),
		tok(tokens.NEWLINE(), 1, 16),

		tok(tokens.SYMBOL("names"), 2, 3),
		tok(tokens.LBRACKET(), 2, 9),
		tok(tokens.RBRACKET(), 2, 10),
		tok(tokens.STR(), 2, 11),
		tok(tokens.NEWLINE(), 2, 14),

		tok(tokens.SYMBOL("birthdate"), 3, 3),
		tok(tokens.SYMBOL("date"), 3, 13),
		tok(tokens.NEWLINE(), 3, 17),

		tok(tokens.RBRACE(), 4, 2),
	}

	testLexer(t, input, expected...)
}

func TestLexer_Example2(t *testing.T) {
	input := `run func(r Runner, run Run) {
		r.runs << run
	}`

	expected := []*slang.Token{
		tok(tokens.SYMBOL("run"), 1, 1),
		tok(tokens.FUNC(), 1, 5),
		tok(tokens.LPAREN(), 1, 9),
		tok(tokens.SYMBOL("r"), 1, 10),
		tok(tokens.SYMBOL("Runner"), 1, 12),
		tok(tokens.COMMA(), 1, 18),
		tok(tokens.SYMBOL("run"), 1, 20),
		tok(tokens.SYMBOL("Run"), 1, 24),
		tok(tokens.RPAREN(), 1, 27),
		tok(tokens.LBRACE(), 1, 29),
		tok(tokens.NEWLINE(), 1, 30),

		tok(tokens.SYMBOL("r"), 2, 3),
		tok(tokens.DOT(), 2, 4),
		tok(tokens.SYMBOL("runs"), 2, 5),
		tok(tokens.LESSLESS(), 2, 10),
		tok(tokens.SYMBOL("run"), 2, 13),
		tok(tokens.NEWLINE(), 2, 16),

		tok(tokens.RBRACE(), 3, 2),
	}

	testLexer(t, input, expected...)
}

func TestLexer_FuncDefinition(t *testing.T) {
	input := `func mul5(x flt) (y flt) {y = 5.0*x}`
	expected := []*slang.Token{
		tok(tokens.FUNC(), 1, 1),
		tok(tokens.SYMBOL("mul5"), 1, 6),
		tok(tokens.LPAREN(), 1, 10),
		tok(tokens.SYMBOL("x"), 1, 11),
		tok(tokens.FLT(), 1, 13),
		tok(tokens.RPAREN(), 1, 16),
		tok(tokens.LPAREN(), 1, 18),
		tok(tokens.SYMBOL("y"), 1, 19),
		tok(tokens.FLT(), 1, 21),
		tok(tokens.RPAREN(), 1, 24),
		tok(tokens.LBRACE(), 1, 26),
		tok(tokens.SYMBOL("y"), 1, 27),
		tok(tokens.EQUAL(), 1, 29),
		tok(tokens.FLTVAL("5.0"), 1, 31),
		tok(tokens.STAR(), 1, 34),
		tok(tokens.SYMBOL("x"), 1, 35),
		tok(tokens.RBRACE(), 1, 36),
	}

	testLexer(t, input, expected...)
}

func TestLexer_IndentWithDigits(t *testing.T) {
	testLexer(t, "var_1", tok(tokens.SYMBOL("var_1"), 1, 1))
}

func tok(info slang.TokenInfo, line, col int) *slang.Token {
	return slang.NewToken(info, slang.NewLoc(line, col))
}

func TestLexer_StringPlusString(t *testing.T) {
	expected := []*slang.Token{
		tok(tokens.STRVAL("abc"), 1, 1),
		tok(tokens.PLUS(), 1, 7),
		tok(tokens.STRVAL("xyz"), 1, 9),
	}

	testLexer(t, `"abc" + "xyz"`, expected...)

	expected = []*slang.Token{
		tok(tokens.STRVAL("foo bar"), 1, 1),
	}

	testLexer(t, `"foo bar"`, expected...)
}

func TestLexer_AssignVerbatimString(t *testing.T) {
	input := "myvar = `this \nis verbatim`"
	expected := []*slang.Token{
		tok(tokens.SYMBOL("myvar"), 1, 1),
		tok(tokens.EQUAL(), 1, 7),
		tok(tokens.VERBATIMSTRING("this \nis verbatim"), 1, 9),
	}

	testLexer(t, input, expected...)
}

func TestLexer_SimpleStringInterp(t *testing.T) {
	expected := []*slang.Token{
		tok(tokens.STRVAL("this string has a "), 1, 1),
		tok(tokens.DOLLARLBRACE(), 1, 20),
		tok(tokens.INTVAL("5"), 1, 22),
		tok(tokens.RBRACE(), 1, 23),
		tok(tokens.STRVAL(" in it"), 1, 24),
	}

	testLexer(t, `"this string has a ${5} in it"`, expected...)
}

func TestLexer_AssignNestedStringInterp(t *testing.T) {
	expected := []*slang.Token{
		tok(tokens.SYMBOL("myVar"), 1, 1),
		tok(tokens.EQUAL(), 1, 7),
		tok(tokens.STRVAL("a "), 1, 9),
		tok(tokens.DOLLARLBRACE(), 1, 12),
		tok(tokens.IF(), 1, 14),
		tok(tokens.SYMBOL("x"), 1, 17),
		tok(tokens.GREATER(), 1, 19),
		tok(tokens.INTVAL("3"), 1, 21),
		tok(tokens.LBRACE(), 1, 23),
		tok(tokens.STRVAL("small"), 1, 24),
		tok(tokens.RBRACE(), 1, 31),
		tok(tokens.ELSE(), 1, 33),
		tok(tokens.LBRACE(), 1, 38),
		tok(tokens.STRVAL("med"), 1, 39),
		tok(tokens.RBRACE(), 1, 44),
		tok(tokens.RBRACE(), 1, 45),
		tok(tokens.STRVAL(" word"), 1, 46),
	}

	testLexer(t, `myVar = "a ${if x > 3 {"small"} else {"med"}} word"`, expected...)
}

func TestLexer_WholeLineComment(t *testing.T) {
	expected := []*slang.Token{
		tok(tokens.COMMENT("not gonna lie..."), 1, 3),
	}

	testLexer(t, `  // not gonna lie... `, expected...)
}

func TestLexer_CommentLines(t *testing.T) {
	const input = `
		// x
		// y
		// z`

	expected := []*slang.Token{
		tok(tokens.NEWLINE(), 1, 1),
		tok(tokens.COMMENT("x"), 2, 3),
		tok(tokens.NEWLINE(), 2, 7),
		tok(tokens.COMMENT("y"), 3, 3),
		tok(tokens.NEWLINE(), 3, 7),
		tok(tokens.COMMENT("z"), 4, 3),
	}

	testLexer(t, input, expected...)
}

func TestLexer_AssignInt(t *testing.T) {
	expected := []*slang.Token{
		tok(tokens.SYMBOL("x"), 1, 4),
		tok(tokens.EQUAL(), 1, 5),
		tok(tokens.INTVAL("5"), 1, 6),
	}

	testLexer(t, "   x=5", expected...)
}

func TestLexer_StatementsWithNewline(t *testing.T) {
	const str = "x = 5\ny = 10"

	expected := []*slang.Token{
		tok(tokens.SYMBOL("x"), 1, 1),
		tok(tokens.EQUAL(), 1, 3),
		tok(tokens.INTVAL("5"), 1, 5),
		tok(tokens.NEWLINE(), 1, 6),
		tok(tokens.SYMBOL("y"), 2, 1),
		tok(tokens.EQUAL(), 2, 3),
		tok(tokens.INTVAL("10"), 2, 5),
	}

	testLexer(t, str, expected...)
}

func TestLexer_Use(t *testing.T) {
	const str = `use "lib/console"`

	expected := []*slang.Token{
		tok(tokens.USE(), 1, 1),
		tok(tokens.STRVAL("lib/console"), 1, 5),
	}

	testLexer(t, str, expected...)
}

func TestLexer_StructBlock(t *testing.T) {
	const str = `struct X {
		Y string
	}`

	expected := []*slang.Token{
		tok(tokens.STRUCT(), 1, 1),
		tok(tokens.SYMBOL("X"), 1, 8),
		tok(tokens.LBRACE(), 1, 10),
		tok(tokens.NEWLINE(), 1, 11),
		tok(tokens.SYMBOL("Y"), 2, 3),
		tok(tokens.SYMBOL("string"), 2, 5),
		tok(tokens.NEWLINE(), 2, 11),
		tok(tokens.RBRACE(), 3, 2),
	}

	testLexer(t, str, expected...)
}

func TestLexer_AssignClassField(t *testing.T) {
	const str = `this.MyField = 7`

	expected := []*slang.Token{
		tok(tokens.SYMBOL("this"), 1, 1),
		tok(tokens.DOT(), 1, 5),
		tok(tokens.SYMBOL("MyField"), 1, 6),
		tok(tokens.EQUAL(), 1, 14),
		tok(tokens.INTVAL("7"), 1, 16),
	}

	testLexer(t, str, expected...)
}

func TestLexer_IntSymbolIsIllegal(t *testing.T) {
	const str = "25Name"

	expected := []*slang.Token{
		tok(tokens.ILLEGAL('N'), 1, 3),
		tok(tokens.SYMBOL("Name"), 1, 3),
	}

	testLexer(t, str, expected...)
}

func TestLexer_IntMethodCall(t *testing.T) {
	const str = "25.add(12)"

	expected := []*slang.Token{
		tok(tokens.INTVAL("25"), 1, 1),
		tok(tokens.DOT(), 1, 3),
		tok(tokens.SYMBOL("add"), 1, 4),
		tok(tokens.LPAREN(), 1, 7),
		tok(tokens.INTVAL("12"), 1, 8),
		tok(tokens.RPAREN(), 1, 10),
	}

	testLexer(t, str, expected...)
}

func TestLexer_FloatMethodCall(t *testing.T) {
	const str = "25.5.add(12)"

	expected := []*slang.Token{
		tok(tokens.FLTVAL("25.5"), 1, 1),
		tok(tokens.DOT(), 1, 5),
		tok(tokens.SYMBOL("add"), 1, 6),
		tok(tokens.LPAREN(), 1, 9),
		tok(tokens.INTVAL("12"), 1, 10),
		tok(tokens.RPAREN(), 1, 12),
	}

	testLexer(t, str, expected...)
}

func TestLexer_AndExprCall(t *testing.T) {
	const str = "x and y"

	expected := []*slang.Token{
		tok(tokens.SYMBOL("x"), 1, 1),
		tok(tokens.AND(), 1, 3),
		tok(tokens.SYMBOL("y"), 1, 7),
	}

	testLexer(t, str, expected...)
}

func TestLexer_OrExprCall(t *testing.T) {
	const str = "x or y"

	expected := []*slang.Token{
		tok(tokens.SYMBOL("x"), 1, 1),
		tok(tokens.OR(), 1, 3),
		tok(tokens.SYMBOL("y"), 1, 6),
	}

	testLexer(t, str, expected...)
}

func TestLexer_FloatMath(t *testing.T) {
	expected := []*slang.Token{
		tok(tokens.SYMBOL("my_num"), 1, 1),
		tok(tokens.EQUAL(), 1, 8),
		tok(tokens.LPAREN(), 1, 10),
		tok(tokens.FLTVAL("2.5"), 1, 11),
		tok(tokens.PLUS(), 1, 15),
		tok(tokens.FLTVAL("7.7"), 1, 17),
		tok(tokens.RPAREN(), 1, 20),
		tok(tokens.STAR(), 1, 22),
		tok(tokens.LPAREN(), 1, 24),
		tok(tokens.SYMBOL("otherNum"), 1, 25),
		tok(tokens.SLASH(), 1, 34),
		tok(tokens.FLTVAL("33.5"), 1, 36),
		tok(tokens.RPAREN(), 1, 40),
	}

	testLexer(t, "my_num = (2.5 + 7.7) * (otherNum / 33.5)", expected...)
}

func TestLexer_AssignFunc(t *testing.T) {
	input := "y = func(myName: uint) {\n\treturn 7\n}"
	expected := []*slang.Token{
		tok(tokens.SYMBOL("y"), 1, 1),
		tok(tokens.EQUAL(), 1, 3),
		tok(tokens.FUNC(), 1, 5),
		tok(tokens.LPAREN(), 1, 9),
		tok(tokens.SYMBOL("myName"), 1, 10),
		tok(tokens.COLON(), 1, 16),
		tok(tokens.SYMBOL("uint"), 1, 18),
		tok(tokens.RPAREN(), 1, 22),
		tok(tokens.LBRACE(), 1, 24),
		tok(tokens.NEWLINE(), 1, 25),
		tok(tokens.RETURN(), 2, 2),
		tok(tokens.INTVAL("7"), 2, 9),
		tok(tokens.NEWLINE(), 2, 10),
		tok(tokens.RBRACE(), 3, 1),
	}

	testLexer(t, input, expected...)
}

func testLexer(t *testing.T, input string, expected ...*slang.Token) {
	toks := lexing.ScanString(input)

	require.Len(t, toks, len(expected))

	for i := 0; i < len(toks); i++ {
		testTokensEqual(t, expected[i], toks[i])
	}
}

func testTokensEqual(t *testing.T, exp, act *slang.Token) {
	result := assert.Equal(t, exp.Info.Type(), act.Info.Type()) &&
		assert.Equal(t, exp.Info.Value(), act.Info.Value()) &&
		assert.Equal(t, exp.Location.Line, act.Location.Line) &&
		assert.Equal(t, exp.Location.Column, act.Location.Column)

	if !result {
		t.Errorf("expected token %#v does not equal actual %#v", exp, act)
	}
}
