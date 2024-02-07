package parsing

import (
	"github.com/jamestunnell/slang"
)

type FileParser struct {
	*ParserBase

	Statements []slang.Statement
}

func NewFileParser() *FileParser {
	return &FileParser{
		ParserBase: NewParserBase(),
		Statements: []slang.Statement{},
	}
}

func (p *FileParser) Run(toks slang.TokenSeq) bool {
	p.Statements = []slang.Statement{}

	for !toks.Current().Is(slang.TokenEOF) {
		toks.Skip(slang.TokenNEWLINE)

		if st := p.parseStatement(toks); st != nil {
			p.Statements = append(p.Statements, st)
		}

		toks.Skip(slang.TokenNEWLINE)
	}

	return true
}

func (p *FileParser) parseStatement(toks slang.TokenSeq) slang.Statement {
	var sp StatementParser

	switch toks.Current().Type() {
	case slang.TokenCLASS:
		sp = NewClassStatementParser()
	case slang.TokenFUNC:
		sp = NewFuncStatementParser()
	case slang.TokenVAR:
		sp = NewVarStatementParser()
	case slang.TokenUSE:
		sp = NewUseStatementParser()
	default:
		p.TokenErr(
			toks.Current(), slang.TokenUSE, slang.TokenFUNC, slang.TokenCLASS, slang.TokenVAR)

		return nil
	}

	if !p.RunSubParser(toks, sp) {
		return nil
	}

	return sp.GetStatement()
}
