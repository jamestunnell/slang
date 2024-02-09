package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
)

type FuncSignatureParser struct {
	*ParserBase

	Params      []slang.Param
	ReturnTypes []slang.Type
}

func NewFuncSignatureParser() *FuncSignatureParser {
	return &FuncSignatureParser{
		ParserBase:  NewParserBase(),
		Params:      []slang.Param{},
		ReturnTypes: []slang.Type{},
	}
}

func (p *FuncSignatureParser) Run(toks slang.TokenSeq) bool {
	p.Params = []slang.Param{}
	p.ReturnTypes = []slang.Type{}

	if !p.ExpectToken(toks.Current(), slang.TokenLPAREN) {
		return false
	}

	toks.AdvanceSkip(slang.TokenNEWLINE)

	p.Params = []slang.Param{}
	p.ReturnTypes = []slang.Type{}

	if !toks.Current().Is(slang.TokenRPAREN) && !p.parseParam(toks) {
		return false
	}

	for !toks.Current().Is(slang.TokenRPAREN) {
		if !p.ExpectToken(toks.Current(), slang.TokenCOMMA) {
			return false
		}

		toks.Advance()

		if !p.parseParam(toks) {
			return false
		}
	}

	toks.Advance()

	// parse return type(s), if any

	addRetType := func() bool {
		typ, ok := p.ParseType(toks)
		if !ok {
			return false
		}

		p.ReturnTypes = append(p.ReturnTypes, typ)

		return true
	}

	switch toks.Current().Type() {
	case slang.TokenSYMBOL:
		if !addRetType() {
			return false
		}
	case slang.TokenLPAREN:
		toks.Advance()

		if !addRetType() {
			return false
		}

		for !toks.Current().Is(slang.TokenRPAREN) {
			if !p.ExpectToken(toks.Current(), slang.TokenCOMMA) {
				return false
			}

			toks.Advance()

			if !addRetType() {
				return false
			}
		}

		toks.Advance()
	}

	return true
}

func (p *FuncSignatureParser) parseParam(toks slang.TokenSeq) bool {
	if !toks.Current().Is(slang.TokenSYMBOL) {
		return false
	}

	name := toks.Current().Value()

	toks.Advance()

	typ, ok := p.ParseType(toks)
	if !ok {
		return false
	}

	p.Params = append(p.Params, ast.NewParam(name, typ))

	return true
}
