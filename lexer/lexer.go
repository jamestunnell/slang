package lexer

import (
	"io"
	"strings"
	"unicode"

	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/tokens"
)

type Lexer struct {
	scanner             io.RuneScanner
	cur, next, nextNext rune
	line, col           int
}

const eof = 0

func New(scanner io.RuneScanner) slang.Lexer {
	l := &Lexer{
		scanner:  scanner,
		cur:      0,
		next:     0,
		nextNext: 0,
		line:     1,
		col:      -2, //will be 1 after advancing thrice
	}

	// read runes for cur, next, and nextNext
	l.advance()
	l.advance()
	l.advance()

	return l
}

func (l *Lexer) NextToken() *slang.Token {
	var tokInfo slang.TokenInfo

	l.skipWhitespace()

	loc := slang.SourceLocation{
		Line:   l.line,
		Column: l.col,
	}

	switch {
	case l.cur == '#':
		tokInfo = l.readComment()
	case l.cur == '\n':
		tokInfo = l.readNewline()
	case isSymbol(l.cur):
		tokInfo = l.readSymbol()
	case l.cur == eof:
		tokInfo = tokens.EOF()
	case l.cur == '"':
		tokInfo = l.readString()
	case l.cur == '`':
		tokInfo = l.readVerbatimString()
	case isLetterOrUnderscore(l.cur):
		tokInfo = l.readSymbolOrKeyword()
	case unicode.IsDigit(l.cur):
		tokInfo = l.readNumber()
	default:
		tokInfo = tokens.ILLEGAL(l.cur)
	}

	l.advance()

	return slang.NewToken(tokInfo, loc)
}

func (l *Lexer) advance() {
	r, _, _ := l.scanner.ReadRune()

	l.col++

	l.cur = l.next
	l.next = l.nextNext
	l.nextNext = r
}

func (l *Lexer) skipWhitespace() {
	for l.cur == ' ' || l.cur == '\t' || l.cur == '\r' {
		l.advance()
	}
}

func isSymbol(r rune) bool {
	switch r {
	case '!', '>', '<', '=', '.', ',', ':', ';', '(', ')', '{', '}', '+', '-', '*', '/', '[', ']':
		return true
	}

	return false
}

func (l *Lexer) readComment() slang.TokenInfo {
	var b strings.Builder

	b.WriteRune('#')

	l.advance()

	for l.cur != eof && l.cur != '\n' {
		b.WriteRune(l.cur)

		l.advance()
	}

	return tokens.COMMENT(b.String())
}

func (l *Lexer) readNewline() slang.TokenInfo {
	l.line++
	l.col = 0

	return tokens.NEWLINE()
}

func (l *Lexer) readSymbol() slang.TokenInfo {
	var tok slang.TokenInfo

	switch l.cur {
	case '!':
		tok = l.readNot()
	case '<':
		tok = l.readLess()
	case '>':
		tok = l.readGreater()
	case '=':
		tok = l.readEqual()
	case '+':
		tok = l.readPlus()
	case '-':
		tok = l.readMinus()
	case '*':
		tok = l.readStar()
	case '/':
		tok = l.readSlash()
	case '.':
		tok = tokens.DOT()
	case ',':
		tok = tokens.COMMA()
	case ':':
		tok = tokens.COLON()
	case ';':
		tok = tokens.SEMICOLON()
	case '(':
		tok = tokens.LPAREN()
	case ')':
		tok = tokens.RPAREN()
	case '{':
		tok = tokens.LBRACE()
	case '[':
		tok = tokens.LBRACKET()
	case '}':
		tok = tokens.RBRACE()
	case ']':
		tok = tokens.RBRACKET()
	default:
		log.Fatal().
			Str("rune", string([]rune{l.cur})).
			Msg("unexpected symbol rune")
	}

	return tok
}

func (l *Lexer) readNot() slang.TokenInfo {
	if l.next == '=' {
		l.advance()

		return tokens.NOTEQUAL()
	}

	return tokens.BANG()
}

func (l *Lexer) readLess() slang.TokenInfo {
	if l.next == '=' {
		l.advance()

		return tokens.LESSEQUAL()
	}

	return tokens.LESS()
}

func (l *Lexer) readGreater() slang.TokenInfo {
	if l.next == '=' {
		l.advance()

		return tokens.GREATEREQUAL()
	}

	return tokens.GREATER()
}

func (l *Lexer) readEqual() slang.TokenInfo {
	if l.next == '=' {
		l.advance()

		return tokens.EQUAL()
	}

	return tokens.ASSIGN()
}

func (l *Lexer) readPlus() slang.TokenInfo {
	switch l.next {
	case '=':
		l.advance()

		return tokens.PLUSEQUAL()
	case '+':
		l.advance()

		return tokens.PLUSPLUS()
	}

	return tokens.PLUS()
}

func (l *Lexer) readMinus() slang.TokenInfo {
	switch l.next {
	case '=':
		l.advance()

		return tokens.MINUSEQUAL()
	case '-':
		l.advance()

		return tokens.MINUSMINUS()
	}

	return tokens.MINUS()
}

func (l *Lexer) readStar() slang.TokenInfo {
	if l.next == '=' {
		l.advance()

		return tokens.STAREQUAL()
	}

	return tokens.STAR()
}

func (l *Lexer) readSlash() slang.TokenInfo {
	if l.next == '=' {
		l.advance()

		return tokens.SLASHEQUAL()
	}

	return tokens.SLASH()
}

func (l *Lexer) readString() slang.TokenInfo {
	l.advance()

	runes := []rune{}

	for l.cur != eof && l.cur != '"' {
		runes = append(runes, l.cur)

		l.advance()
	}

	if l.cur == eof {
		return tokens.ILLEGAL(l.cur)
	}

	return tokens.STRING(string(runes))
}

func (l *Lexer) readVerbatimString() slang.TokenInfo {
	l.advance()

	runes := []rune{}

	for l.cur != eof && l.cur != '`' {
		runes = append(runes, l.cur)

		l.advance()
	}

	if l.cur == eof {
		return tokens.ILLEGAL(l.cur)
	}

	return tokens.VERBATIMSTRING(string(runes))
}

func (l *Lexer) readSymbolOrKeyword() slang.TokenInfo {
	runes := []rune{l.cur}

	for unicode.IsDigit(l.next) || isLetterOrUnderscore(l.next) {
		l.advance()

		runes = append(runes, l.cur)
	}

	str := string(runes)

	switch str {
	case tokens.StrCLASS:
		return tokens.CLASS()
	case tokens.StrELSE:
		return tokens.ELSE()
	case tokens.StrFALSE:
		return tokens.FALSE()
	case tokens.StrFIELD:
		return tokens.FIELD()
	case tokens.StrFUNC:
		return tokens.FUNC()
	case tokens.StrIF:
		return tokens.IF()
	case tokens.StrMETHOD:
		return tokens.METHOD()
	case tokens.StrRETURN:
		return tokens.RETURN()
	case tokens.StrTRUE:
		return tokens.TRUE()
	case tokens.StrUSE:
		return tokens.USE()
	}

	return tokens.SYMBOL(str)
}

func (l *Lexer) readNumber() slang.TokenInfo {
	preDotDigits := []rune{l.cur}

	for unicode.IsDigit(l.next) {
		l.advance()

		preDotDigits = append(preDotDigits, l.cur)
	}

	if l.next != '.' || !unicode.IsDigit(l.nextNext) {
		return tokens.INT(string(preDotDigits))
	}

	l.advance()
	l.advance()

	postDotDigits := []rune{l.cur}

	if unicode.IsDigit(l.next) {
		l.advance()

		postDotDigits = append(postDotDigits, l.cur)
	}

	runes := append(preDotDigits, '.')
	runes = append(runes, postDotDigits...)

	return tokens.FLOAT(string(runes))
}

func isLetterOrUnderscore(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}
