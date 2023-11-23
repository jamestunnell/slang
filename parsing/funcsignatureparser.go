package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
)

type FuncSignatureParser struct {
	*ParserBase

	Params      []*ast.Param
	ReturnTypes []string
}

func NewFuncSignatureParser() *FuncSignatureParser {
	return &FuncSignatureParser{
		ParserBase:  NewParserBase(),
		Params:      []*ast.Param{},
		ReturnTypes: []string{},
	}
}

func (p *FuncSignatureParser) Run(toks slang.TokenSeq) {
	p.Params = []*ast.Param{}
	p.ReturnTypes = []string{}

	if !p.ExpectToken(toks.Current(), slang.TokenLPAREN) {
		return
	}

	toks.AdvanceSkip(slang.TokenNEWLINE)

	p.Params = []*ast.Param{}
	p.ReturnTypes = []string{}

	if !toks.Current().Is(slang.TokenRPAREN) && !p.parseParam(toks) {
		return
	}

	for !toks.Current().Is(slang.TokenRPAREN) {
		if !p.ExpectToken(toks.Current(), slang.TokenCOMMA) {
			return
		}

		toks.Advance()

		if !p.parseParam(toks) {
			return
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
			return
		}
	case slang.TokenLPAREN:
		toks.Advance()

		if !addRetType() {
			return
		}

		for !toks.Current().Is(slang.TokenRPAREN) {
			if !p.ExpectToken(toks.Current(), slang.TokenCOMMA) {
				return
			}

			toks.Advance()

			if !addRetType() {
				return
			}
		}

		toks.Advance()
	}
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

	p.Params = append(p.Params, &ast.Param{Name: name, Type: typ})

	return true
}
