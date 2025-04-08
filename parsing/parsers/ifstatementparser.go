package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
)

type IfStatementParser struct {
	*ParserBase

	Stmt *statements.Statement
}

func NewIfStatementParser() *IfStatementParser {
	return &IfStatementParser{
		ParserBase: NewParserBase(),
	}
}

func (p *IfStatementParser) GetStatement() *statements.Statement {
	return p.Stmt
}

func (p *IfStatementParser) Run(
	toks slang.TokenSeq,
	comment string,
) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenIF) {
		return false
	}

	toks.Advance()

	condParser := NewExprParser(parsing.PrecedenceLOWEST)
	if !p.RunSubParser(toks, condParser) {
		return false
	}

	ifBodyParser := NewCondBodyParser()
	if !p.RunSubParser(toks, ifBodyParser) {
		return false
	}

	if !toks.Current().Is(slang.TokenELSE) {
		p.Stmt = statements.NewIf(condParser.Expr, ifBodyParser.GetStatements())

		p.Stmt.SetComment(comment)

		return true
	}

	toks.Advance()

	elseBodyParser := NewCondBodyParser()
	if !p.RunSubParser(toks, elseBodyParser) {
		return false
	}

	p.Stmt = statements.NewIfElse(
		condParser.Expr, ifBodyParser.GetStatements(), elseBodyParser.GetStatements())

	p.Stmt.SetComment(comment)

	return true
}
