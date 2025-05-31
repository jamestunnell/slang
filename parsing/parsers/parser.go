package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/types"
	"github.com/jamestunnell/slang/customerrs"
	"github.com/jamestunnell/slang/parsing"
)

type ParserBase struct {
	errors []*parsing.ParseErr
}

func NewParserBase() *ParserBase {
	return &ParserBase{
		errors: []*parsing.ParseErr{},
	}
}

func (p *ParserBase) AddError(err *parsing.ParseErr) {
	p.errors = append(p.errors, err)
}

func (p *ParserBase) GetErrors() []*parsing.ParseErr {
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
	parseErr := parsing.NewParseError(err, tok)

	p.errors = append(p.errors, parseErr)
}

func (p *ParserBase) RunSubParser(
	toks parsing.TokenSeq,
	sub parsing.Parser,
) bool {
	sub.Run(toks)

	if len(sub.GetErrors()) > 0 {
		p.errors = append(p.errors, sub.GetErrors()...)

		return false
	}

	return true
}

func (p *ParserBase) RunSubStmtParser(
	toks parsing.TokenSeq,
	comment string,
	sub StatementParser,
) bool {
	sub.Run(toks, comment)

	if len(sub.GetErrors()) > 0 {
		p.errors = append(p.errors, sub.GetErrors()...)

		return false
	}

	return true
}

func (p *ParserBase) ParseType(toks parsing.TokenSeq) (*types.Type, bool) {
	switch toks.Current().Type() {
	case slang.TokenARY:
		return p.ParseArrayType(toks)
	case slang.TokenBOOL:
		toks.Advance()

		return types.NewBool(), true
	case slang.TokenERR:
		toks.Advance()

		return types.NewErr(), true
	case slang.TokenINT:
		toks.Advance()

		return types.NewInt(), true
	case slang.TokenMAP:
		return p.ParseMapType(toks)
	case slang.TokenSTR:
		toks.Advance()

		return types.NewStr(), true
	case slang.TokenFLT:
		toks.Advance()

		return types.NewFlt(), true
	case slang.TokenSYMBOL:
		nameFirst := toks.Current().Value()

		if !toks.Next().Is(slang.TokenDOT) {
			toks.Advance()

			return types.NewStruct("", nameFirst), true
		}

		toks.Advance()
		toks.Advance()

		if p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
			nameSecond := toks.Current().Value()

			toks.Advance()

			return types.NewStruct(nameFirst, nameSecond), true
		}
	}

	return nil, false
}

// func (p *ParserBase) ParseBasicType(toks parsing.TokenSeq) (slang.Type, bool) {
// 	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
// 		return nil, false
// 	}

// 	parts := []string{toks.Current().Value()}

// 	toks.Advance()

// 	for toks.Current().Is(slang.TokenDOT) {
// 		if !p.ExpectToken(toks.Next(), slang.TokenSYMBOL) {
// 			return nil, false
// 		}

// 		parts = append(parts, toks.Next().Value())

// 		toks.Advance()
// 		toks.Advance()
// 	}

// 	return ast.NewBasicType(parts...), true
// }

func (p *ParserBase) ParseArrayType(toks parsing.TokenSeq) (*types.Type, bool) {
	if !p.ExpectToken(toks.Next(), slang.TokenLESS) {
		return nil, false
	}

	toks.Advance()
	toks.Advance()

	valType, ok := p.ParseType(toks)
	if !ok {
		return nil, false
	}

	if !p.ExpectToken(toks.Current(), slang.TokenGREATER) {
		return nil, false
	}

	toks.Advance()

	return types.NewArray(valType), true
}

func (p *ParserBase) ParseMapType(toks parsing.TokenSeq) (*types.Type, bool) {
	if !p.ExpectToken(toks.Next(), slang.TokenLESS) {
		return nil, false
	}

	toks.Advance()
	toks.Advance()

	keyType, ok := p.ParseType(toks)
	if !ok {
		return nil, false
	}

	if !p.ExpectToken(toks.Current(), slang.TokenCOMMA) {
		return nil, false
	}

	toks.Advance()

	valType, ok := p.ParseType(toks)
	if !ok {
		return nil, false
	}

	if !p.ExpectToken(toks.Current(), slang.TokenGREATER) {
		return nil, false
	}

	toks.Advance()

	return types.NewMap(keyType, valType), true
}

func (p *ParserBase) ParseNameTypePair(toks parsing.TokenSeq) (string, *types.Type, bool) {
	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
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

func (p *ParserBase) ParseNamesType(toks parsing.TokenSeq) ([]string, *types.Type, bool) {
	if !toks.Current().Is(slang.TokenSYMBOL) {
		return []string{}, nil, false
	}

	names := []string{toks.Current().Value()}

	toks.Advance()

	for toks.Current().Is(slang.TokenCOMMA) {
		if !p.ExpectToken(toks.Next(), slang.TokenSYMBOL) {
			return []string{}, nil, false
		}

		toks.Advance()

		names = append(names, toks.Current().Value())

		toks.Advance()
	}

	typ, ok := p.ParseType(toks)
	if !ok {
		return []string{}, nil, false
	}

	return names, typ, true
}
