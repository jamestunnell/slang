package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
)

type AssignStatementParser struct {
	*ParserBase

	Stmt *statements.Statement
}

func NewAssignStatementParser() *AssignStatementParser {
	return &AssignStatementParser{ParserBase: NewParserBase()}
}

func (p *AssignStatementParser) GetStatement() slang.Statement {
	return p.Stmt
}

func (p *AssignStatementParser) Run(
	toks parsing.TokenSeq,
	comment string,
) bool {
	exprParser := NewExprParser(parsing.PrecedenceLOWEST)
	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	if !p.ExpectToken(toks.Current(), slang.TokenEQUAL) {
		return false
	}

	toks.AdvanceSkip(slang.TokenNEWLINE)

	valueParser := NewExprParser(parsing.PrecedenceLOWEST)
	if !p.RunSubParser(toks, valueParser) {
		return false
	}

	p.Stmt = statements.NewAssign(exprParser.Expr, valueParser.Expr)

	p.Stmt.SetComment(comment)

	return true
}
