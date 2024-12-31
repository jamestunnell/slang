package lexing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/slang/lexing"
)

func TestLexer_IndentWithDigits(t *testing.T) {
	testLexer(t, "var_1", tok(lexing.SYMBOL("var_1"), 1, 1))
}

func tok(info *lexing.TokenInfo, line, col int) *lexing.Token {
	return lexing.NewToken(info, lexing.NewLoc(line, col))
}

func TestLexer_StringPlusString(t *testing.T) {
	expected := []*lexing.Token{
		tok(lexing.STRING("abc"), 1, 1),
		tok(lexing.PLUS(), 1, 7),
		tok(lexing.STRING("xyz"), 1, 9),
	}

	testLexer(t, `"abc" + "xyz"`, expected...)

	expected = []*lexing.Token{
		tok(lexing.STRING("foo bar"), 1, 1),
	}

	testLexer(t, `"foo bar"`, expected...)
}

func TestLexer_AssignVerbatimString(t *testing.T) {
	input := "myvar = `this \nis verbatim`"
	expected := []*lexing.Token{
		tok(lexing.SYMBOL("myvar"), 1, 1),
		tok(lexing.EQUAL(), 1, 7),
		tok(lexing.VERBATIMSTRING("this \nis verbatim"), 1, 9),
	}

	testLexer(t, input, expected...)
}

func TestLexer_SimpleStringInterp(t *testing.T) {
	expected := []*lexing.Token{
		tok(lexing.STRING("this string has a "), 1, 1),
		tok(lexing.DOLLARLBRACE(), 1, 20),
		tok(lexing.INT("5"), 1, 22),
		tok(lexing.RBRACE(), 1, 23),
		tok(lexing.STRING(" in it"), 1, 24),
	}

	testLexer(t, `"this string has a ${5} in it"`, expected...)
}

func TestLexer_AssignNestedStringInterp(t *testing.T) {
	expected := []*lexing.Token{
		tok(lexing.SYMBOL("myVar"), 1, 1),
		tok(lexing.EQUAL(), 1, 7),
		tok(lexing.STRING("a "), 1, 9),
		tok(lexing.DOLLARLBRACE(), 1, 12),
		tok(lexing.IF(), 1, 14),
		tok(lexing.SYMBOL("x"), 1, 17),
		tok(lexing.GREATER(), 1, 19),
		tok(lexing.INT("3"), 1, 21),
		tok(lexing.LBRACE(), 1, 23),
		tok(lexing.STRING("small"), 1, 24),
		tok(lexing.RBRACE(), 1, 31),
		tok(lexing.ELSE(), 1, 33),
		tok(lexing.LBRACE(), 1, 38),
		tok(lexing.STRING("med"), 1, 39),
		tok(lexing.RBRACE(), 1, 44),
		tok(lexing.RBRACE(), 1, 45),
		tok(lexing.STRING(" word"), 1, 46),
	}

	testLexer(t, `myVar = "a ${if x > 3 {"small"} else {"med"}} word"`, expected...)
}

func TestLexer_WholeLineComment(t *testing.T) {
	expected := []*lexing.Token{
		tok(lexing.COMMENT("# not gonna lie..."), 1, 3),
	}

	testLexer(t, `  # not gonna lie...`, expected...)
}

func TestLexer_InlineComment(t *testing.T) {
	expected := []*lexing.Token{
		tok(lexing.SYMBOL("x"), 1, 1),
		tok(lexing.EQUAL(), 1, 3),
		tok(lexing.INT("10"), 1, 5),
		tok(lexing.COMMENT("# this is why"), 1, 8),
	}

	testLexer(t, `x = 10 # this is why`, expected...)
}

func TestLexer_AssignInt(t *testing.T) {
	expected := []*lexing.Token{
		tok(lexing.SYMBOL("x"), 1, 4),
		tok(lexing.EQUAL(), 1, 5),
		tok(lexing.INT("5"), 1, 6),
	}

	testLexer(t, "   x=5", expected...)
}

func TestLexer_StatementsWithNewline(t *testing.T) {
	const str = "x = 5\ny = 10"

	expected := []*lexing.Token{
		tok(lexing.SYMBOL("x"), 1, 1),
		tok(lexing.EQUAL(), 1, 3),
		tok(lexing.INT("5"), 1, 5),
		tok(lexing.NEWLINE(), 1, 6),
		tok(lexing.SYMBOL("y"), 2, 1),
		tok(lexing.EQUAL(), 2, 3),
		tok(lexing.INT("10"), 2, 5),
	}

	testLexer(t, str, expected...)
}

func TestLexer_Use(t *testing.T) {
	const str = `use "lib/console"`

	expected := []*lexing.Token{
		tok(lexing.USE(), 1, 1),
		tok(lexing.STRING("lib/console"), 1, 5),
	}

	testLexer(t, str, expected...)
}

func TestLexer_StructBlock(t *testing.T) {
	const str = `class X {
		Y string
	}`

	expected := []*lexing.Token{
		tok(lexing.CLASS(), 1, 1),
		tok(lexing.SYMBOL("X"), 1, 7),
		tok(lexing.LBRACE(), 1, 9),
		tok(lexing.NEWLINE(), 1, 10),
		tok(lexing.SYMBOL("Y"), 2, 3),
		tok(lexing.SYMBOL("string"), 2, 5),
		tok(lexing.NEWLINE(), 2, 11),
		tok(lexing.RBRACE(), 3, 2),
	}

	testLexer(t, str, expected...)
}

func TestLexer_AssignClassField(t *testing.T) {
	const str = `this.MyField = 7`

	expected := []*lexing.Token{
		tok(lexing.SYMBOL("this"), 1, 1),
		tok(lexing.DOT(), 1, 5),
		tok(lexing.SYMBOL("MyField"), 1, 6),
		tok(lexing.EQUAL(), 1, 14),
		tok(lexing.INT("7"), 1, 16),
	}

	testLexer(t, str, expected...)
}

func TestLexer_IntSymbolIsIllegal(t *testing.T) {
	const str = "25Name"

	expected := []*lexing.Token{
		tok(lexing.ILLEGAL('N'), 1, 3),
		tok(lexing.SYMBOL("Name"), 1, 3),
	}

	testLexer(t, str, expected...)
}

func TestLexer_IntMethodCall(t *testing.T) {
	const str = "25.add(12)"

	expected := []*lexing.Token{
		tok(lexing.INT("25"), 1, 1),
		tok(lexing.DOT(), 1, 3),
		tok(lexing.SYMBOL("add"), 1, 4),
		tok(lexing.LPAREN(), 1, 7),
		tok(lexing.INT("12"), 1, 8),
		tok(lexing.RPAREN(), 1, 10),
	}

	testLexer(t, str, expected...)
}

func TestLexer_FloatMethodCall(t *testing.T) {
	const str = "25.5.add(12)"

	expected := []*lexing.Token{
		tok(lexing.FLOAT("25.5"), 1, 1),
		tok(lexing.DOT(), 1, 5),
		tok(lexing.SYMBOL("add"), 1, 6),
		tok(lexing.LPAREN(), 1, 9),
		tok(lexing.INT("12"), 1, 10),
		tok(lexing.RPAREN(), 1, 12),
	}

	testLexer(t, str, expected...)
}

func TestLexer_AndExprCall(t *testing.T) {
	const str = "x and y"

	expected := []*lexing.Token{
		tok(lexing.SYMBOL("x"), 1, 1),
		tok(lexing.AND(), 1, 3),
		tok(lexing.SYMBOL("y"), 1, 7),
	}

	testLexer(t, str, expected...)
}

func TestLexer_OrExprCall(t *testing.T) {
	const str = "x or y"

	expected := []*lexing.Token{
		tok(lexing.SYMBOL("x"), 1, 1),
		tok(lexing.OR(), 1, 3),
		tok(lexing.SYMBOL("y"), 1, 6),
	}

	testLexer(t, str, expected...)
}

func TestLexer_FloatMath(t *testing.T) {
	expected := []*lexing.Token{
		tok(lexing.SYMBOL("my_num"), 1, 1),
		tok(lexing.EQUAL(), 1, 8),
		tok(lexing.LPAREN(), 1, 10),
		tok(lexing.FLOAT("2.5"), 1, 11),
		tok(lexing.PLUS(), 1, 15),
		tok(lexing.FLOAT("7.7"), 1, 17),
		tok(lexing.RPAREN(), 1, 20),
		tok(lexing.STAR(), 1, 22),
		tok(lexing.LPAREN(), 1, 24),
		tok(lexing.SYMBOL("otherNum"), 1, 25),
		tok(lexing.SLASH(), 1, 34),
		tok(lexing.FLOAT("33.5"), 1, 36),
		tok(lexing.RPAREN(), 1, 40),
	}

	testLexer(t, "my_num = (2.5 + 7.7) * (otherNum / 33.5)", expected...)
}

func TestLexer_AssignFunc(t *testing.T) {
	input := "y = func(myName: uint) {\n\treturn 7\n}"
	expected := []*lexing.Token{
		tok(lexing.SYMBOL("y"), 1, 1),
		tok(lexing.EQUAL(), 1, 3),
		tok(lexing.FUNC(), 1, 5),
		tok(lexing.LPAREN(), 1, 9),
		tok(lexing.SYMBOL("myName"), 1, 10),
		tok(lexing.COLON(), 1, 16),
		tok(lexing.SYMBOL("uint"), 1, 18),
		tok(lexing.RPAREN(), 1, 22),
		tok(lexing.LBRACE(), 1, 24),
		tok(lexing.NEWLINE(), 1, 25),
		tok(lexing.RETURN(), 2, 2),
		tok(lexing.INT("7"), 2, 9),
		tok(lexing.NEWLINE(), 2, 10),
		tok(lexing.RBRACE(), 3, 1),
	}

	testLexer(t, input, expected...)
}

func testLexer(t *testing.T, input string, expected ...*lexing.Token) {
	toks := lexing.ScanString(input)

	require.Len(t, toks, len(expected))

	for i := 0; i < len(toks); i++ {
		testTokensEqual(t, expected[i], toks[i])
	}
}

func testTokensEqual(t *testing.T, exp, act *lexing.Token) {
	result := assert.Equal(t, exp.Type, act.Type) &&
		assert.Equal(t, exp.Value, act.Value) &&
		assert.Equal(t, exp.Location.Line, act.Location.Line) &&
		assert.Equal(t, exp.Location.Column, act.Location.Column)

	if !result {
		t.Errorf("expected token %#v does not equal actual %#v", exp, act)
	}
}
