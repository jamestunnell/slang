package lexer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/lexer"
	"github.com/jamestunnell/slang/tokens"
)

func TestLexer_IndentWithDigits(t *testing.T) {
	testLexer(t, "var_1", tok(tokens.SYMBOL("var_1"), 1, 1))
}

func tok(info slang.TokenInfo, line, col int) *slang.Token {
	return slang.NewToken(info, slang.NewLoc(line, col))
}

func TestLexer_StringPlusString(t *testing.T) {
	expected := []*slang.Token{
		tok(tokens.STRING("abc"), 1, 1),
		tok(tokens.PLUS(), 1, 7),
		tok(tokens.STRING("xyz"), 1, 9),
	}

	testLexer(t, `"abc" + "xyz"`, expected...)

	expected = []*slang.Token{
		tok(tokens.STRING("foo bar"), 1, 1),
	}

	testLexer(t, `"foo bar"`, expected...)
}

func TestLexer_WholeLineComment(t *testing.T) {
	expected := []*slang.Token{
		tok(tokens.COMMENT("# not gonna lie..."), 1, 3),
	}

	testLexer(t, `  # not gonna lie...`, expected...)
}

func TestLexer_InlineComment(t *testing.T) {
	expected := []*slang.Token{
		tok(tokens.SYMBOL("x"), 1, 1),
		tok(tokens.ASSIGN(), 1, 3),
		tok(tokens.INT("10"), 1, 5),
		tok(tokens.COMMENT("# this is why"), 1, 8),
	}

	testLexer(t, `x = 10 # this is why`, expected...)
}

func TestLexer_AssignInt(t *testing.T) {
	expected := []*slang.Token{
		tok(tokens.SYMBOL("x"), 1, 4),
		tok(tokens.ASSIGN(), 1, 5),
		tok(tokens.INT("5"), 1, 6),
	}

	testLexer(t, "   x=5", expected...)
}

func TestLexer_StatementsWithNewline(t *testing.T) {
	const str = "x = 5\ny = 10"

	expected := []*slang.Token{
		tok(tokens.SYMBOL("x"), 1, 1),
		tok(tokens.ASSIGN(), 1, 3),
		tok(tokens.INT("5"), 1, 5),
		tok(tokens.NEWLINE(), 1, 6),
		tok(tokens.SYMBOL("y"), 2, 1),
		tok(tokens.ASSIGN(), 2, 3),
		tok(tokens.INT("10"), 2, 5),
	}

	testLexer(t, str, expected...)
}

func TestLexer_Use(t *testing.T) {
	const str = `use "lib/console"`

	expected := []*slang.Token{
		tok(tokens.USE(), 1, 1),
		tok(tokens.STRING("lib/console"), 1, 5),
	}

	testLexer(t, str, expected...)
}

func TestLexer_StructBlock(t *testing.T) {
	const str = `class X {
		Y string
	}`

	expected := []*slang.Token{
		tok(tokens.CLASS(), 1, 1),
		tok(tokens.SYMBOL("X"), 1, 7),
		tok(tokens.LBRACE(), 1, 9),
		tok(tokens.NEWLINE(), 1, 10),
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
		tok(tokens.ASSIGN(), 1, 14),
		tok(tokens.INT("7"), 1, 16),
	}

	testLexer(t, str, expected...)
}

func TestLexer_IntMethodCall(t *testing.T) {
	const str = "25.add(12)"

	expected := []*slang.Token{
		tok(tokens.INT("25"), 1, 1),
		tok(tokens.DOT(), 1, 3),
		tok(tokens.SYMBOL("add"), 1, 4),
		tok(tokens.LPAREN(), 1, 7),
		tok(tokens.INT("12"), 1, 8),
		tok(tokens.RPAREN(), 1, 10),
	}

	testLexer(t, str, expected...)
}

func TestLexer_FloatMath(t *testing.T) {
	expected := []*slang.Token{
		tok(tokens.SYMBOL("my_num"), 1, 1),
		tok(tokens.ASSIGN(), 1, 8),
		tok(tokens.LPAREN(), 1, 10),
		tok(tokens.FLOAT("2.5"), 1, 11),
		tok(tokens.PLUS(), 1, 15),
		tok(tokens.FLOAT("7.7"), 1, 17),
		tok(tokens.RPAREN(), 1, 20),
		tok(tokens.STAR(), 1, 22),
		tok(tokens.LPAREN(), 1, 24),
		tok(tokens.SYMBOL("otherNum"), 1, 25),
		tok(tokens.SLASH(), 1, 34),
		tok(tokens.FLOAT("33.5"), 1, 36),
		tok(tokens.RPAREN(), 1, 40),
	}

	testLexer(t, "my_num = (2.5 + 7.7) * (otherNum / 33.5)", expected...)
}

func TestLexer_AssignFunc(t *testing.T) {
	input := "y = func(myName: uint) {\n\treturn 7\n}"
	expected := []*slang.Token{
		tok(tokens.SYMBOL("y"), 1, 1),
		tok(tokens.ASSIGN(), 1, 3),
		tok(tokens.FUNC(), 1, 5),
		tok(tokens.LPAREN(), 1, 9),
		tok(tokens.SYMBOL("myName"), 1, 10),
		tok(tokens.COLON(), 1, 16),
		tok(tokens.SYMBOL("uint"), 1, 18),
		tok(tokens.RPAREN(), 1, 22),
		tok(tokens.LBRACE(), 1, 24),
		tok(tokens.NEWLINE(), 1, 25),
		tok(tokens.RETURN(), 2, 2),
		tok(tokens.INT("7"), 2, 9),
		tok(tokens.NEWLINE(), 2, 10),
		tok(tokens.RBRACE(), 3, 1),
	}

	testLexer(t, input, expected...)
}

func testLexer(t *testing.T, input string, expected ...*slang.Token) {
	toks := lexer.ScanString(input)

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
		t.Errorf("expected token %v does not equal actual %v", exp, act)
	}
}
