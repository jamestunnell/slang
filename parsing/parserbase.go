package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
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

func (p *ParserBase) GetErrors() []*ParseErr {
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
	err := customerrs.NewErrWrongTokenType(tok, expectedTypes...)
	parseErr := NewParseError(err, tok)

	p.errors = append(p.errors, parseErr)
}

func (p *ParserBase) RunSubParser(toks slang.TokenSeq, sub Parser) bool {
	sub.Run(toks)

	if len(sub.GetErrors()) > 0 {
		p.errors = append(p.errors, sub.GetErrors()...)

		return false
	}

	return true
}

func (p *ParserBase) ParseType(toks slang.TokenSeq) (slang.Type, bool) {
	switch toks.Current().Type() {
	case slang.TokenSYMBOL:
		return p.ParseBasicType(toks)
	case slang.TokenLBRACKET:
		if toks.Next().Is(slang.TokenRBRACKET) {
			return p.ParseArrayType(toks)
		}

		return p.ParseMapType(toks)
	}

	p.TokenErr(toks.Current(), slang.TokenSYMBOL, slang.TokenLBRACKET)

	return nil, false
}

func (p *ParserBase) ParseBasicType(toks slang.TokenSeq) (slang.Type, bool) {
	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return nil, false
	}

	parts := []string{toks.Current().Value()}

	toks.Advance()

	for toks.Current().Is(slang.TokenDOT) {
		if !p.ExpectToken(toks.Next(), slang.TokenSYMBOL) {
			return nil, false
		}

		parts = append(parts, toks.Next().Value())

		toks.Advance()
		toks.Advance()
	}

	return ast.NewBasicType(parts...), true
}

func (p *ParserBase) ParseArrayType(toks slang.TokenSeq) (*ast.ArrayType, bool) {
	if !p.ExpectToken(toks.Current(), slang.TokenLBRACKET) ||
		!p.ExpectToken(toks.Next(), slang.TokenRBRACKET) {
		return nil, false
	}

	toks.Advance()
	toks.Advance()

	valType, ok := p.ParseType(toks)
	if !ok {
		return nil, false
	}

	return ast.NewArrayType(valType), true
}

func (p *ParserBase) ParseMapType(toks slang.TokenSeq) (*ast.MapType, bool) {
	if !p.ExpectToken(toks.Current(), slang.TokenLBRACKET) {
		return nil, false
	}

	toks.Advance()

	keyType, ok := p.ParseType(toks)
	if !ok {
		return nil, false
	}

	if !p.ExpectToken(toks.Current(), slang.TokenRBRACKET) {
		return nil, false
	}

	toks.Advance()

	valType, ok := p.ParseType(toks)
	if !ok {
		return nil, false
	}

	return ast.NewMapType(keyType, valType), true
}

func (p *ParserBase) ParseNameTypePair(toks slang.TokenSeq) (string, slang.Type, bool) {
	if !toks.Current().Is(slang.TokenSYMBOL) {
		return "", nil, false
	}

	name := toks.Current().Value()

	toks.Advance()

	typ, ok := p.ParseType(toks)
	if !ok {
		return "", nil, false
	}

	return name, typ, true
}
