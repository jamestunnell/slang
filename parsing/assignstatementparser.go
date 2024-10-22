package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
)

type AssignStatementParser struct {
	*ParserBase

	Stmt slang.Statement
}

func NewAssignStatementParser() *AssignStatementParser {
	return &AssignStatementParser{ParserBase: NewParserBase()}
}

func (p *AssignStatementParser) GetStatement() slang.Statement {
	return p.Stmt
}

func (p *AssignStatementParser) Run(toks lexing.TokenSeq) bool {
	exprParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	if !p.ExpectToken(toks.Current(), lexing.TokenASSIGN) {
		return false
	}

	toks.AdvanceSkip(lexing.TokenNEWLINE)

	valueParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, valueParser) {
		return false
	}

	p.Stmt = statements.NewAssign(exprParser.Expr, valueParser.Expr)

	return true
}
