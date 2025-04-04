package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type ExprOrAssignStatementParser struct {
	*ParserBase

	Stmt *statements.Statement
}

func NewExprOrAssignStatementParser() *ExprOrAssignStatementParser {
	return &ExprOrAssignStatementParser{ParserBase: NewParserBase()}
}

func (p *ExprOrAssignStatementParser) GetStatement() *statements.Statement {
	return p.Stmt
}

func (p *ExprOrAssignStatementParser) Run(
	toks slang.TokenSeq,
	comment string,
) bool {
	exprParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	if toks.Current().Is(slang.TokenEQUAL) {
		toks.AdvanceSkip(slang.TokenNEWLINE)

		valueParser := NewExprParser(PrecedenceLOWEST)
		if !p.RunSubParser(toks, valueParser) {
			return false
		}

		p.Stmt = statements.NewAssign(exprParser.Expr, valueParser.Expr)
	} else {
		p.Stmt = statements.NewExpression(exprParser.Expr)
	}

	p.Stmt.SetComment(comment)

	return true
}
