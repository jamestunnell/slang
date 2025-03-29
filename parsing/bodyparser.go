package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type BodyParserBase struct {
	*ParserBase

	Statements []slang.Statement

	makeStmtParser MakeStmtParserFunc
}

type MakeStmtParserFunc func(cur *slang.Token) StatementParser

func NewBodyParserBase(makeStmtParser MakeStmtParserFunc) *BodyParserBase {
	return &BodyParserBase{
		ParserBase:     NewParserBase(),
		Statements:     []slang.Statement{},
		makeStmtParser: makeStmtParser,
	}
}

func (p *BodyParserBase) GetStatements() []slang.Statement {
	return p.Statements
}

func (p *BodyParserBase) parseStatement(
	toks slang.TokenSeq,
) bool {
	commentLines := []string{}

	for toks.Current().Is(slang.TokenCOMMENT) {
		commentLines = append(commentLines, toks.Current().Value())

		if toks.AdvanceSkip(slang.TokenNEWLINE) > 1 || toks.Current().Is(slang.TokenRBRACE) {
			p.Statements = append(p.Statements, statements.NewComment(commentLines...))

			return true
		}
	}

	stmtParser := p.makeStmtParser(toks.Current())
	if stmtParser == nil {
		return false
	}

	if !p.RunSubParser(toks, stmtParser) {
		return false
	}

	stmt := stmtParser.GetStatement()

	if len(commentLines) > 0 {
		stmt.SetComment(commentLines)
	}

	p.Statements = append(p.Statements, stmt)

	return true
}

func (p *BodyParserBase) Run(toks slang.TokenSeq) bool {
	p.Statements = []slang.Statement{}

	if !p.ExpectToken(toks.Current(), slang.TokenLBRACE) {
		return false
	}

	_ = toks.AdvanceSkip(slang.TokenNEWLINE)

	for !toks.Current().Is(slang.TokenRBRACE) {
		if !p.parseStatement(toks) {
			return false
		}

		// if !p.ExpectToken(toks.Current(), slang.TokenNEWLINE, slang.TokenRBRACE) {
		// 	return false
		// }

		_ = toks.Skip(slang.TokenNEWLINE)
	}

	toks.Advance()

	return true
}
