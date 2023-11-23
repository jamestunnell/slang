package parsing

import (
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type ParserBase struct {
	errors []*ParseErr
}

func NewParserBase() *ParserBase {
	return &ParserBase{
		errors: []*ParseErr{},
	}
}

func (p *ParserBase) Errors() []*ParseErr {
	return p.errors
}

func (p *ParserBase) ExpectToken(
	tok *slang.Token,
	expectedTypes ...slang.TokenType) bool {
	if !tok.Is(expectedTypes...) {
		p.TokenErr(tok, expectedTypes...)

		return false
	}

	return true
}

func (p *ParserBase) TokenErr(tok *slang.Token, expectedTypes ...slang.TokenType) {
	err := customerrs.NewErrWrongTokenType(tok.Type(), expectedTypes)
	parseErr := NewParseError(err, tok)

	p.errors = append(p.errors, parseErr)
}

func (p *ParserBase) RunSubParser(toks slang.TokenSeq, sub Parser) bool {
	sub.Run(toks)

	if len(sub.Errors()) > 0 {
		p.errors = append(p.errors, sub.Errors()...)

		return false
	}

	return true
}

func (p *ParserBase) ParseType(toks slang.TokenSeq) (string, bool) {
	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return "", false
	}

	typ := toks.Current().Value()

	toks.Advance()

	if toks.Current().Is(slang.TokenDOT) {
		if !p.ExpectToken(toks.Next(), slang.TokenSYMBOL) {
			return "", false
		}

		typ = fmt.Sprintf("%s.%s", typ, toks.Next().Value())

		toks.Advance()
		toks.Advance()
	}

	return typ, true
}

func (p *ParserBase) ParseNameTypePair(toks slang.TokenSeq) (string, string, bool) {
	if !toks.Current().Is(slang.TokenSYMBOL) {
		return "", "", false
	}

	name := toks.Current().Value()

	toks.Advance()

	typ, ok := p.ParseType(toks)
	if !ok {
		return "", "", false
	}

	return name, typ, true
}
