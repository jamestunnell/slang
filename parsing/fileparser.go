package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/lexing"
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

func (p *FileParser) Run(toks lexing.TokenSeq) bool {
	p.Statements = []slang.Statement{}

	for !toks.Current().Is(lexing.TokenEOF) {
		toks.Skip(lexing.TokenNEWLINE)

		if st := p.parseStatement(toks); st != nil {
			p.Statements = append(p.Statements, st)
		}

		toks.Skip(lexing.TokenNEWLINE)
	}

	return true
}

func (p *FileParser) parseStatement(toks lexing.TokenSeq) slang.Statement {
	var sp StatementParser

	switch toks.Current().Type {
	case lexing.TokenCLASS:
		sp = NewClassStatementParser()
	case lexing.TokenCONST:
		sp = NewConstStatementParser()
	case lexing.TokenFUNC:
		sp = NewFuncStatementParser()
	case lexing.TokenVAR:
		sp = NewVarStatementParser()
	case lexing.TokenUSE:
		sp = NewUseStatementParser()
	default:
		p.TokenErr(
			toks.Current(), lexing.TokenUSE, lexing.TokenFUNC, lexing.TokenCLASS, lexing.TokenVAR)

		return nil
	}

	if !p.RunSubParser(toks, sp) {
		return nil
	}

	return sp.GetStatement()
}
