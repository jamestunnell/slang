package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
)

type ClassBodyParser struct {
	*ParserBase

	cls *ast.Class
}

func NewClassBodyParser() *ClassBodyParser {

	return &ClassBodyParser{
		ParserBase: NewParserBase(),
	}
}

func (p *ClassBodyParser) Class() *ast.Class {
	return p.cls
}

func (p *ClassBodyParser) Run(toks slang.TokenSeq) {
	if !p.ExpectToken(toks.Current(), slang.TokenLBRACE) {
		return
	}

	toks.AdvanceSkip(slang.TokenNEWLINE)

	p.cls = ast.NewClass()

	for !toks.Current().Is(slang.TokenRBRACE) {
		if !p.parseMember(toks) {
			return
		}

		toks.Skip(slang.TokenNEWLINE)
	}

	toks.Advance()
}

func (p *ClassBodyParser) parseMember(toks slang.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenFIELD, slang.TokenMETHOD) {
		return false
	}

	tokType := toks.Current().Type()

	toks.Advance()

	switch tokType {
	case slang.TokenFIELD:
		name, typ, ok := p.ParseNameTypePair(toks)
		if !ok {
			return false
		}

		p.cls.Fields = append(p.cls.Fields, ast.NewParam(name, typ))
	case slang.TokenMETHOD:
		if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
			return false
		}

		name := toks.Current().Value()

		toks.Advance()

		sigParser := NewFuncSignatureParser()
		if !p.RunSubParser(toks, sigParser) {
			return false
		}

		bodyParser := NewFuncBodyParser()
		if !p.RunSubParser(toks, bodyParser) {
			return false
		}

		p.cls.Methods[name] = ast.NewFunction(
			sigParser.Params, sigParser.ReturnTypes, bodyParser.Statements...)
	}

	return true
}
