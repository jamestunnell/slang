package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/statements"
)

type ClassBodyParser struct {
	*ParserBase

	Statements []slang.Statement
}

func NewClassBodyParser() *ClassBodyParser {
	return &ClassBodyParser{
		ParserBase: NewParserBase(),
	}
}

func (p *ClassBodyParser) GetStatements() []slang.Statement {
	return p.Statements
}

func (p *ClassBodyParser) Run(toks slang.TokenSeq) {
	p.Statements = []slang.Statement{}

	if !p.ExpectToken(toks.Current(), slang.TokenLBRACE) {
		return
	}

	toks.AdvanceSkip(slang.TokenNEWLINE)

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

		p.Statements = append(p.Statements, statements.NewField(name, typ))
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

		fn := ast.NewFunction(
			sigParser.Params, sigParser.ReturnTypes, bodyParser.Statements...)
		p.Statements = append(p.Statements, statements.NewMethod(name, fn))
	}

	return true
}
