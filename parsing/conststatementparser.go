package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type ConstStatementParser struct {
	*ParserBase

	ConstStmt *statements.Const
}

func NewConstStatementParser() *ConstStatementParser {
	return &ConstStatementParser{ParserBase: NewParserBase()}
}

func (p *ConstStatementParser) GetStatement() slang.Statement {
	return p.ConstStmt
}

func (p *ConstStatementParser) Run(toks slang.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenCONST) {
		return false
	}

	toks.Advance()

	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return false
	}

	name := toks.Current().Value()

	toks.Advance()

	if !p.ExpectToken(toks.Current(), slang.TokenASSIGN) {
		return false
	}

	toks.Advance()

	exprParser := NewExprParser(PrecedenceLOWEST)

	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	p.ConstStmt = statements.NewConst(name, exprParser.Expr)

	return true
}
