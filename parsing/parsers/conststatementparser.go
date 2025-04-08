package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
)

type ConstStatementParser struct {
	*ParserBase

	ConstStmt *statements.Statement
}

func NewConstStatementParser() *ConstStatementParser {
	return &ConstStatementParser{ParserBase: NewParserBase()}
}

func (p *ConstStatementParser) GetStatement() *statements.Statement {
	return p.ConstStmt
}

func (p *ConstStatementParser) Run(toks slang.TokenSeq, comment string) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenCONST) {
		return false
	}

	toks.Advance()

	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return false
	}

	name := toks.Current().Value()

	toks.Advance()

	exprParser := NewExprParser(parsing.PrecedenceLOWEST)

	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	p.ConstStmt = statements.NewConst(name, exprParser.Expr)

	p.ConstStmt.SetComment(comment)

	return true
}
