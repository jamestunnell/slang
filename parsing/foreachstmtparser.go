package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
)

type ForEachStmtParser struct {
	*ParserBase

	Stmt slang.Statement
}

func NewForEachStmtParser() *ForEachStmtParser {
	return &ForEachStmtParser{
		ParserBase: NewParserBase(),
	}
}

func (p *ForEachStmtParser) GetStatement() slang.Statement {
	return p.Stmt
}

func (p *ForEachStmtParser) Run(toks lexing.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), lexing.TokenFOREACH) {
		return false
	}

	toks.Advance()

	// expect one or more var names
	if !p.ExpectToken(toks.Current(), lexing.TokenSYMBOL) {
		return false
	}

	vars := []string{toks.Current().Value}

	toks.Advance()

	for toks.Current().Is(lexing.TokenCOMMA) {
		toks.Advance()

		if !p.ExpectToken(toks.Current(), lexing.TokenSYMBOL) {
			return false
		}

		vars = append(vars, toks.Current().Value)

		toks.Advance()
	}

	if !p.ExpectToken(toks.Current(), lexing.TokenIN) {
		return false
	}

	toks.Advance()

	exprParser := NewExprParser(PrecedenceLOWEST)

	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	bodyParser := NewCondBodyParser()

	if !p.RunSubParser(toks, bodyParser) {
		return false
	}

	block := statements.NewBlock(bodyParser.Statements...)

	p.Stmt = statements.NewForEach(vars, exprParser.Expr, block)

	return true
}
