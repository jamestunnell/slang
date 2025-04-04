package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
)

type BodyParser interface {
	parsing.Parser

	GetStatements() []*statements.Statement
}

type BodyParserBase struct {
	*ParserBase

	Statements []*statements.Statement

	makeStmtParser MakeStmtParserFunc
}

type MakeStmtParserFunc func(cur *slang.Token) StatementParser

func NewBodyParserBase(makeStmtParser MakeStmtParserFunc) *BodyParserBase {
	return &BodyParserBase{
		ParserBase:     NewParserBase(),
		Statements:     []*statements.Statement{},
		makeStmtParser: makeStmtParser,
	}
}

func (p *BodyParserBase) GetStatements() []*statements.Statement {
	return p.Statements
}

func (p *BodyParserBase) parseStatement(
	toks slang.TokenSeq,
) bool {
	commentLines := []string{}

	for toks.Current().Is(slang.TokenCOMMENT) {
		commentLines = append(commentLines, toks.Current().Value())

		if toks.AdvanceSkip(slang.TokenNEWLINE) > 1 || toks.Current().Is(slang.TokenRBRACE) {
			stmt := statements.NewComment()

			stmt.SetComment(makeComment(commentLines))

			p.Statements = append(p.Statements, stmt)

			return true
		}
	}

	stmtParser := p.makeStmtParser(toks.Current())
	if stmtParser == nil {
		return false
	}

	if !p.RunSubStmtParser(toks, makeComment(commentLines), stmtParser) {
		return false
	}

	p.Statements = append(p.Statements, stmtParser.GetStatement())

	return true
}

func (p *BodyParserBase) Run(toks slang.TokenSeq) bool {
	p.Statements = []*statements.Statement{}

	if !p.ExpectToken(toks.Current(), slang.TokenLBRACE) {
		return false
	}

	_ = toks.AdvanceSkip(slang.TokenNEWLINE)

	for !toks.Current().Is(slang.TokenRBRACE) {
		if !p.parseStatement(toks) {
			return false
		}

		_ = toks.Skip(slang.TokenNEWLINE)
	}

	toks.Advance()

	return true
}
