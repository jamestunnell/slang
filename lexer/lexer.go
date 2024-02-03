package lexer

import (
	"io"
	"strings"
	"unicode"

	"github.com/oleiade/lane/v2"
	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/tokens"
)

type Lexer struct {
	scanner       io.RuneScanner
	cur, next     rune
	line, col     int
	interpContext *lane.Stack[slang.TokenType]
	tokens        *lane.Queue[*slang.Token]
}

const eof = 0

func New(scanner io.RuneScanner) slang.Lexer {
	l := &Lexer{
		scanner:       scanner,
		cur:           0,
		next:          0,
		line:          1,
		col:           -1, //will be 1 after advancing twice
		interpContext: lane.NewStack[slang.TokenType](),
		tokens:        lane.NewQueue[*slang.Token](),
	}

	// read runes for cur and next
	l.advance()
	l.advance()

	return l
}

func (l *Lexer) NextToken() *slang.Token {
	if tok, ok := l.tokens.Dequeue(); ok {
		return tok
	}

	l.skipWhitespace()

	loc := l.curLoc()

	switch {
	case l.cur == '#':
		l.readComment(loc)
	case l.cur == '\n':
		l.readNewline(loc)
	case l.cur == '"':
		l.advance()

		l.readString(loc)
	case l.cur == '`':
		l.readVerbatimString(loc)
	case isSymbol(l.cur):
		l.readSymbol(loc)
	case l.cur == eof:
		l.emit(tokens.EOF(), loc)
	case isLetterOrUnderscore(l.cur):
		l.readNameOrKeyword(loc)
	case unicode.IsDigit(l.cur):
		l.readNumber(loc)
	default:
		l.emit(tokens.ILLEGAL(l.cur), loc)

		l.advance()
	}

	if next, ok := l.tokens.Dequeue(); ok {
		return next
	}

	log.Fatal().Msg("no token to dequeue")

	return nil
}

func (l *Lexer) emit(info slang.TokenInfo, loc slang.SourceLocation) {
	l.tokens.Enqueue(slang.NewToken(info, loc))
}

func (l *Lexer) advance() {
	r, _, _ := l.scanner.ReadRune()

	l.col++

	l.cur = l.next
	l.next = r
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

func (l *Lexer) readComment(loc slang.SourceLocation) {
	var b strings.Builder

	b.WriteRune('#')

	l.advance()

	for l.cur != eof && l.cur != '\n' {
		b.WriteRune(l.cur)

		l.advance()
	}

	l.advance()

	l.emit(tokens.COMMENT(b.String()), loc)
}

func (l *Lexer) advanceLine() {
	l.line++
	l.col = 0

	l.advance()
}

func (l *Lexer) readNewline(loc slang.SourceLocation) {
	l.advanceLine()

	l.emit(tokens.NEWLINE(), loc)
}

func (l *Lexer) readSymbol(loc slang.SourceLocation) {
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

	l.emit(tok, loc)

	l.advance()

	if l.interpContext.Size() == 0 {
		return
	}

	if tok.Type() == slang.TokenLBRACE {
		l.interpContext.Push(slang.TokenLBRACE)

		return
	}

	if tok.Type() != slang.TokenRBRACE {
		return
	}

	contextType, _ := l.interpContext.Pop()
	if contextType == slang.TokenDOLLARLBRACE {
		// completed a string interpolation expression, now resume reading string until "
		l.readString(l.curLoc())
	}
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

func (l *Lexer) curLoc() slang.SourceLocation {
	return slang.SourceLocation{
		Line:   l.line,
		Column: l.col,
	}
}

func (l *Lexer) readString(loc slang.SourceLocation) {
	var b strings.Builder

	for l.cur != '\n' && l.cur != eof && l.cur != '"' {
		if l.cur == '$' && l.next == '{' {
			l.emit(tokens.STRING(b.String()), loc)
			l.emit(tokens.DOLLARLBRACE(), l.curLoc())

			l.advance()
			l.advance()

			l.interpContext.Push(slang.TokenDOLLARLBRACE)

			return
		}

		b.WriteRune(l.cur)

		l.advance()
	}

	if l.cur == eof || l.cur == '\n' {
		l.emit(tokens.ILLEGAL(l.cur), l.curLoc())
	}

	l.advance()

	l.emit(tokens.STRING(b.String()), loc)
}

func (l *Lexer) readVerbatimString(loc slang.SourceLocation) {
	l.advance()

	var b strings.Builder

	for l.cur != eof && l.cur != '`' {
		b.WriteRune(l.cur)

		// check for newline in string
		if l.cur == '\n' {
			l.advanceLine()
		} else {
			l.advance()
		}
	}

	if l.cur == eof {
		l.emit(tokens.ILLEGAL(l.cur), l.curLoc())

		return
	}

	l.advance()

	l.emit(tokens.VERBATIMSTRING(b.String()), loc)
}

func (l *Lexer) readNameOrKeyword(loc slang.SourceLocation) {
	var b strings.Builder

	b.WriteRune(l.cur)

	for unicode.IsDigit(l.next) || isLetterOrUnderscore(l.next) {
		l.advance()

		b.WriteRune(l.cur)
	}

	l.advance()

	str := b.String()

	switch str {
	case tokens.StrAND:
		l.emit(tokens.AND(), loc)
	case tokens.StrCLASS:
		l.emit(tokens.CLASS(), loc)
	case tokens.StrELSE:
		l.emit(tokens.ELSE(), loc)
	case tokens.StrFALSE:
		l.emit(tokens.FALSE(), loc)
	case tokens.StrFIELD:
		l.emit(tokens.FIELD(), loc)
	case tokens.StrVAR:
		l.emit(tokens.VAR(), loc)
	case tokens.StrFUNC:
		l.emit(tokens.FUNC(), loc)
	case tokens.StrIF:
		l.emit(tokens.IF(), loc)
	case tokens.StrMETHOD:
		l.emit(tokens.METHOD(), loc)
	case tokens.StrOR:
		l.emit(tokens.OR(), loc)
	case tokens.StrRETURN:
		l.emit(tokens.RETURN(), loc)
	case tokens.StrTRUE:
		l.emit(tokens.TRUE(), loc)
	case tokens.StrUSE:
		l.emit(tokens.USE(), loc)
	default:
		l.emit(tokens.SYMBOL(str), loc)
	}
}

func (l *Lexer) readNumber(loc slang.SourceLocation) {
	var b strings.Builder

	for unicode.IsDigit(l.cur) {
		b.WriteRune(l.cur)

		l.advance()
	}

	if isLetterOrUnderscore(l.cur) {
		l.emit(tokens.ILLEGAL(l.cur), l.curLoc())

		return
	}

	if l.cur == '.' {
		if unicode.IsDigit(l.next) {
			b.WriteRune('.')

			l.advance()

			for unicode.IsDigit(l.cur) {
				b.WriteRune(l.cur)

				l.advance()
			}

			l.emit(tokens.FLOAT(b.String()), loc)

		} else {
			l.emit(tokens.INT(b.String()), loc)
			l.emit(tokens.DOT(), l.curLoc())

			l.advance()
		}
	} else {
		l.emit(tokens.INT(b.String()), loc)
	}
}

func isLetterOrUnderscore(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}
