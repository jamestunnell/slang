package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/types"
)

type DataSignatureParser struct {
	*ParserBase

	NameTypes []slang.NameType
}

func NewDataSignatureParser() *DataSignatureParser {
	return &DataSignatureParser{
		ParserBase: NewParserBase(),
		NameTypes:  []slang.NameType{},
	}
}

func (p *DataSignatureParser) Run(toks slang.TokenSeq) bool {
	p.NameTypes = []slang.NameType{}

	if !p.ExpectToken(toks.Current(), slang.TokenLPAREN) {
		return false
	}

	toks.Advance()

	for {
		toks.Skip(slang.TokenNEWLINE, slang.TokenCOMMENT)

		if toks.Current().Is(slang.TokenRPAREN) {
			toks.Advance()

			break
		}

		if !p.parseNamesType(toks) {
			return false
		}

		toks.Skip(slang.TokenCOMMA)
	}

	return true

	// for toks.Current().Is(slang.TokenCOMMA, slang.TokenNEWLINE) {
	// 	toks.AdvanceSkip(slang.TokenNEWLINE)

	// 	if !p.parseNamesType(toks) {
	// 		return false
	// 	}
	// }

	// if !p.ExpectToken(toks.Current(), slang.TokenRPAREN) {
	// 	return false
	// }

	// toks.Advance()

	// return true
}

func (p *DataSignatureParser) parseNamesType(toks slang.TokenSeq) bool {
	names, typ, ok := p.ParseNamesType(toks)
	if !ok {
		return false
	}

	for _, name := range names {
		p.NameTypes = append(p.NameTypes, types.NewNameType(name, typ))
	}

	return true
}
