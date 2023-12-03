package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type IfStatementParser struct {
	*ParserBase

	Stmt slang.Statement
}

func NewIfStatementParser() *IfStatementParser {
	return &IfStatementParser{
		ParserBase: NewParserBase(),
	}
}

func (p *IfStatementParser) GetStatement() slang.Statement {
	return p.Stmt
}

func (p *IfStatementParser) Run(toks slang.TokenSeq) bool {
	toks.Advance()

	condParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, condParser) {
		return false
	}

	ifBodyParser := NewCondBodyParser()
	if !p.RunSubParser(toks, ifBodyParser) {
		return false
	}

	ifBlock := statements.NewBlock(ifBodyParser.Statements)

	if !toks.Current().Is(slang.TokenELSE) {
		p.Stmt = statements.NewIf(condParser.Expr, ifBlock)

		return true
	}

	toks.Advance()

	elseBodyParser := NewCondBodyParser()
	if !p.RunSubParser(toks, elseBodyParser) {
		return false
	}

	elseBlock := statements.NewBlock(elseBodyParser.Statements)

	p.Stmt = statements.NewIfElse(condParser.Expr, ifBlock, elseBlock)

	return true
}
