package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/lexing"
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

func (p *FuncSignatureParser) Run(toks lexing.TokenSeq) bool {
	p.Params = []slang.Param{}
	p.ReturnTypes = []slang.Type{}

	if !p.ExpectToken(toks.Current(), lexing.TokenLPAREN) {
		return false
	}

	toks.AdvanceSkip(lexing.TokenNEWLINE)

	p.Params = []slang.Param{}
	p.ReturnTypes = []slang.Type{}

	if !toks.Current().Is(lexing.TokenRPAREN) && !p.parseParam(toks) {
		return false
	}

	for !toks.Current().Is(lexing.TokenRPAREN) {
		if !p.ExpectToken(toks.Current(), lexing.TokenCOMMA) {
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

	switch toks.Current().Type {
	case lexing.TokenSYMBOL:
		if !addRetType() {
			return false
		}
	case lexing.TokenLPAREN:
		toks.Advance()

		if !addRetType() {
			return false
		}

		for !toks.Current().Is(lexing.TokenRPAREN) {
			if !p.ExpectToken(toks.Current(), lexing.TokenCOMMA) {
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

func (p *FuncSignatureParser) parseParam(toks lexing.TokenSeq) bool {
	if !toks.Current().Is(lexing.TokenSYMBOL) {
		return false
	}

	name := toks.Current().Value

	toks.Advance()

	typ, ok := p.ParseType(toks)
	if !ok {
		return false
	}

	p.Params = append(p.Params, ast.NewParam(name, typ))

	return true
}
