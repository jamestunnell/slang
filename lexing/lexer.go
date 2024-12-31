package lexing

import (
	"io"
	"strings"
	"unicode"

	"github.com/oleiade/lane/v2"
	"github.com/rs/zerolog/log"
)

type Lexer interface {
	NextToken() *Token
}

type lexer struct {
	scanner       io.RuneScanner
	cur, next     rune
	line, col     int
	interpContext *lane.Stack[TokenType]
	toks          *lane.Queue[*Token]
}

const eof = 0

func NewLexer(scanner io.RuneScanner) Lexer {
	l := &lexer{
		scanner:       scanner,
		cur:           0,
		next:          0,
		line:          1,
		col:           -1, //will be 1 after advancing twice
		interpContext: lane.NewStack[TokenType](),
		toks:          lane.NewQueue[*Token](),
	}

	// read runes for cur and next
	l.advance()
	l.advance()

	return l
}

func (l *lexer) NextToken() *Token {
	if tok, ok := l.toks.Dequeue(); ok {
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
		l.emit(EOF(), loc)
	case isLetterOrUnderscore(l.cur):
		l.readNameOrKeyword(loc)
	case unicode.IsDigit(l.cur):
		l.readNumber(loc)
	default:
		l.emit(ILLEGAL(l.cur), loc)

		l.advance()
	}

	if next, ok := l.toks.Dequeue(); ok {
		return next
	}

	log.Fatal().Msg("no token to dequeue")

	return nil
}

func (l *lexer) emit(info *TokenInfo, loc SourceLocation) {
	l.toks.Enqueue(NewToken(info, loc))
}

func (l *lexer) advance() {
	r, _, _ := l.scanner.ReadRune()

	l.col++

	l.cur = l.next
	l.next = r
}

func (l *lexer) skipWhitespace() {
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

func (l *lexer) readComment(loc SourceLocation) {
	var b strings.Builder

	b.WriteRune('#')

	l.advance()

	for l.cur != eof && l.cur != '\n' {
		b.WriteRune(l.cur)

		l.advance()
	}

	l.advance()

	l.emit(COMMENT(b.String()), loc)
}

func (l *lexer) advanceLine() {
	l.line++
	l.col = 0

	l.advance()
}

func (l *lexer) readNewline(loc SourceLocation) {
	l.advanceLine()

	l.emit(NEWLINE(), loc)
}

func (l *lexer) readSymbol(loc SourceLocation) {
	var tok *TokenInfo

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
		tok = PLUS()
	case '-':
		tok = MINUS()
	case '*':
		tok = STAR()
	case '/':
		tok = SLASH()
	case '.':
		tok = DOT()
	case ',':
		tok = COMMA()
	case ':':
		tok = COLON()
	case ';':
		tok = SEMICOLON()
	case '(':
		tok = LPAREN()
	case ')':
		tok = RPAREN()
	case '{':
		tok = LBRACE()
	case '[':
		tok = LBRACKET()
	case '}':
		tok = RBRACE()
	case ']':
		tok = RBRACKET()
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

	if tok.Type == TokenLBRACE {
		l.interpContext.Push(TokenLBRACE)

		return
	}

	if tok.Type != TokenRBRACE {
		return
	}

	contextType, _ := l.interpContext.Pop()
	if contextType == TokenDOLLARLBRACE {
		// completed a string interpolation expression, now resume reading string until "
		l.readString(l.curLoc())
	}
}

func (l *lexer) readEqual() *TokenInfo {
	if l.next == '=' {
		l.advance()

		return EQUALEQUAL()
	}

	return EQUAL()
}

func (l *lexer) readNot() *TokenInfo {
	if l.next == '=' {
		l.advance()

		return NOTEQUAL()
	}

	return BANG()
}

func (l *lexer) readLess() *TokenInfo {
	if l.next == '=' {
		l.advance()

		return LESSEQUAL()
	}

	return LESS()
}

func (l *lexer) readGreater() *TokenInfo {
	if l.next == '=' {
		l.advance()

		return GREATEREQUAL()
	}

	return GREATER()
}

func (l *lexer) curLoc() SourceLocation {
	return SourceLocation{
		Line:   l.line,
		Column: l.col,
	}
}

func (l *lexer) readString(loc SourceLocation) {
	var b strings.Builder

	for l.cur != '\n' && l.cur != eof && l.cur != '"' {
		if l.cur == '$' && l.next == '{' {
			l.emit(STRING(b.String()), loc)
			l.emit(DOLLARLBRACE(), l.curLoc())

			l.advance()
			l.advance()

			l.interpContext.Push(TokenDOLLARLBRACE)

			return
		}

		b.WriteRune(l.cur)

		l.advance()
	}

	if l.cur == eof || l.cur == '\n' {
		l.emit(ILLEGAL(l.cur), l.curLoc())
	}

	l.advance()

	l.emit(STRING(b.String()), loc)
}

func (l *lexer) readVerbatimString(loc SourceLocation) {
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
		l.emit(ILLEGAL(l.cur), l.curLoc())

		return
	}

	l.advance()

	l.emit(VERBATIMSTRING(b.String()), loc)
}

func (l *lexer) readNameOrKeyword(loc SourceLocation) {
	var b strings.Builder

	b.WriteRune(l.cur)

	for unicode.IsDigit(l.next) || isLetterOrUnderscore(l.next) {
		l.advance()

		b.WriteRune(l.cur)
	}

	l.advance()

	str := b.String()

	switch str {
	case StrAND:
		l.emit(AND(), loc)
	case StrBREAK:
		l.emit(BREAK(), loc)
	case StrCLASS:
		l.emit(CLASS(), loc)
	case StrCONST:
		l.emit(CONST(), loc)
	case StrCONTINUE:
		l.emit(CONTINUE(), loc)
	case StrELSE:
		l.emit(ELSE(), loc)
	case StrFALSE:
		l.emit(FALSE(), loc)
	case StrFIELD:
		l.emit(FIELD(), loc)
	case StrFOREACH:
		l.emit(FOREACH(), loc)
	case StrFUNC:
		l.emit(FUNC(), loc)
	case StrIF:
		l.emit(IF(), loc)
	case StrIN:
		l.emit(IN(), loc)
	case StrMETHOD:
		l.emit(METHOD(), loc)
	case StrOR:
		l.emit(OR(), loc)
	case StrRETURN:
		l.emit(RETURN(), loc)
	case StrTRUE:
		l.emit(TRUE(), loc)
	case StrUSE:
		l.emit(USE(), loc)
	case StrVAR:
		l.emit(VAR(), loc)
	default:
		l.emit(SYMBOL(str), loc)
	}
}

func (l *lexer) readNumber(loc SourceLocation) {
	var b strings.Builder

	for unicode.IsDigit(l.cur) {
		b.WriteRune(l.cur)

		l.advance()
	}

	if isLetterOrUnderscore(l.cur) {
		l.emit(ILLEGAL(l.cur), l.curLoc())

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

			l.emit(FLOAT(b.String()), loc)

		} else {
			l.emit(INT(b.String()), loc)
			l.emit(DOT(), l.curLoc())

			l.advance()
		}
	} else {
		l.emit(INT(b.String()), loc)
	}
}

func isLetterOrUnderscore(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}
