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

		if !p.parseStatement(toks) {
			return false
		}

		toks.Skip(slang.TokenNEWLINE)
	}

	return true
}

func (p *FileParser) parseStatement(toks slang.TokenSeq) bool {
	var sp StatementParser

	switch toks.Current().Type() {
	case slang.TokenCLASS:
		sp = NewClassStatementParser()
	case slang.TokenFUNC:
		sp = NewFuncStatementParser()
	case slang.TokenUSE:
		sp = NewUseStatementParser()
	default:
		sp = NewAssignStatementParser()
	}

	if !p.RunSubParser(toks, sp) {
		return false
	}

	p.Statements = append(p.Statements, sp.GetStatement())

	return true
}
