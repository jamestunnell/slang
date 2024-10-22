package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
)

type ExprOrAssignStatementParser struct {
	*ParserBase

	Stmt slang.Statement
}

func NewExprOrAssignStatementParser() *ExprOrAssignStatementParser {
	return &ExprOrAssignStatementParser{ParserBase: NewParserBase()}
}

func (p *ExprOrAssignStatementParser) GetStatement() slang.Statement {
	return p.Stmt
}

func (p *ExprOrAssignStatementParser) Run(toks lexing.TokenSeq) bool {
	exprParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	if toks.Current().Is(lexing.TokenASSIGN) {
		toks.AdvanceSkip(lexing.TokenNEWLINE)

		valueParser := NewExprParser(PrecedenceLOWEST)
		if !p.RunSubParser(toks, valueParser) {
			return false
		}

		p.Stmt = statements.NewAssign(exprParser.Expr, valueParser.Expr)
	} else {
		p.Stmt = statements.NewExpression(exprParser.Expr)
	}

	return true
}
