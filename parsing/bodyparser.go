package parsing

import (
	"github.com/jamestunnell/slang"
)

type BodyParserBase struct {
	*ParserBase

	Statements []slang.Statement

	parseStatement func(slang.TokenSeq) bool
}

func NewBodyParserBase(parseStatement func(slang.TokenSeq) bool) *BodyParserBase {
	return &BodyParserBase{
		ParserBase:     NewParserBase(),
		Statements:     []slang.Statement{},
		parseStatement: parseStatement,
	}
}

func (p *BodyParserBase) GetStatements() []slang.Statement {
	return p.Statements
}

func (p *BodyParserBase) ParseStatement(toks slang.TokenSeq, sp StatementParser) bool {
	if !p.RunSubParser(toks, sp) {
		return false
	}

	p.Statements = append(p.Statements, sp.GetStatement())

	return true
}

func (p *BodyParserBase) Run(toks slang.TokenSeq) bool {
	p.Statements = []slang.Statement{}

	if !p.ExpectToken(toks.Current(), slang.TokenLBRACE) {
		return false
	}

	toks.AdvanceSkip(slang.TokenNEWLINE)

	for !toks.Current().Is(slang.TokenRBRACE) {
		if !p.parseStatement(toks) {
			return false
		}

		if !p.ExpectToken(toks.Current(), slang.TokenNEWLINE, slang.TokenRBRACE) {
			return false
		}

		toks.Skip(slang.TokenNEWLINE)
	}

	toks.Advance()

	return true
}
