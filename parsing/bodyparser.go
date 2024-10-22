package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/lexing"
)

type BodyParserBase struct {
	*ParserBase

	Statements []slang.Statement

	parseStatement ParseStmtFunc
}

type ParseStmtFunc func(lexing.TokenSeq) slang.Statement

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
	toks lexing.TokenSeq,
	sp StatementParser,
) slang.Statement {
	if !p.RunSubParser(toks, sp) {
		return nil
	}

	return sp.GetStatement()
}

func (p *BodyParserBase) Run(toks lexing.TokenSeq) bool {
	p.Statements = []slang.Statement{}

	if !p.ExpectToken(toks.Current(), lexing.TokenLBRACE) {
		return false
	}

	toks.AdvanceSkip(lexing.TokenNEWLINE)

	for !toks.Current().Is(lexing.TokenRBRACE) {
		if st := p.parseStatement(toks); st != nil {
			p.Statements = append(p.Statements, st)
		}

		if !p.ExpectToken(toks.Current(), lexing.TokenNEWLINE, lexing.TokenRBRACE) {
			return false
		}

		toks.Skip(lexing.TokenNEWLINE)
	}

	toks.Advance()

	return true
}
