package lexer

import (
	"io"
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

	for isSpaceOrTab(l.cur) {
		l.advance()
	}

	loc := slang.SourceLocation{
		Line:   l.line,
		Column: l.col,
	}

	switch l.cur {
	case '\n', '\v', '\f':
		tokInfo = tokens.LINE()

		l.line++
		l.col = 0
	case '!', '>', '<', '=', '.', ',', ';', '(', ')', '{', '}', '+', '-', '*', '/', '[', ']':
		tokInfo = l.readSymbol()
	case 0:
		tokInfo = tokens.EOF()
	case '"':
		tokInfo = l.readString()
	default:
		if isLetterOrUnderscore(l.cur) {
			tokInfo = l.readIdentOrKeyword()
		} else if unicode.IsDigit(l.cur) {
			tokInfo = l.readNumber()
		} else {
			tokInfo = tokens.ILLEGAL(l.cur)
		}
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

	for l.cur != 0 && l.cur != '"' {
		runes = append(runes, l.cur)

		l.advance()
	}

	if l.cur == 0 {
		return tokens.ILLEGAL(l.cur)
	}

	return tokens.STRING(string(runes))
}

func (l *Lexer) readIdentOrKeyword() slang.TokenInfo {
	runes := []rune{l.cur}

	for unicode.IsDigit(l.next) || isLetterOrUnderscore(l.next) {
		l.advance()

		runes = append(runes, l.cur)
	}

	str := string(runes)

	switch str {
	case tokens.StrELSE:
		return tokens.ELSE()
	case tokens.StrFALSE:
		return tokens.FALSE()
	case tokens.StrFUNC:
		return tokens.FUNC()
	case tokens.StrIF:
		return tokens.IF()
	case tokens.StrRETURN:
		return tokens.RETURN()
	case tokens.StrTRUE:
		return tokens.TRUE()
	}

	return tokens.IDENT(str)
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

func isSpaceOrTab(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}
