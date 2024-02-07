package parsing

import (
	"github.com/jamestunnell/slang"
)

type BodyParserBase struct {
	*ParserBase

	Statements []slang.Statement

	parseStatement ParseStmtFunc
}

type ParseStmtFunc func(slang.TokenSeq) slang.Statement

func NewBodyParserBase(parseStatement ParseStmtFunc) *BodyParserBase {
	return &BodyParserBase{
		ParserBase:     NewParserBase(),
		Statements:     []slang.Statement{},
		parseStatement: parseStatement,
	}
}

func (p *BodyParserBase) GetStatements() []slang.Statement {
	return p.Statements
}

func (p *BodyParserBase) ParseStatement(
	toks slang.TokenSeq,
	sp StatementParser,
) slang.Statement {
	if !p.RunSubParser(toks, sp) {
		return nil
	}

	return sp.GetStatement()
}

func (p *BodyParserBase) Run(toks slang.TokenSeq) bool {
	p.Statements = []slang.Statement{}

	if !p.ExpectToken(toks.Current(), slang.TokenLBRACE) {
		return false
	}

	toks.AdvanceSkip(slang.TokenNEWLINE)

	for !toks.Current().Is(slang.TokenRBRACE) {
		if st := p.parseStatement(toks); st != nil {
			p.Statements = append(p.Statements, st)
		}

		if !p.ExpectToken(toks.Current(), slang.TokenNEWLINE, slang.TokenRBRACE) {
			return false
		}

		toks.Skip(slang.TokenNEWLINE)
	}

	toks.Advance()

	return true
}
