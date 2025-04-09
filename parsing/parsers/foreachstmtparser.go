package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
)

type ForEachStmtParser struct {
	*ParserBase

	Stmt *statements.Statement
}

func NewForEachStmtParser() *ForEachStmtParser {
	return &ForEachStmtParser{
		ParserBase: NewParserBase(),
	}
}

func (p *ForEachStmtParser) GetStatement() *statements.Statement {
	return p.Stmt
}

func (p *ForEachStmtParser) Run(
	toks parsing.TokenSeq,
	comment string,
) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenFOREACH) {
		return false
	}

	toks.Advance()

	// expect one or more var names
	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return false
	}

	vars := []string{toks.Current().Value()}

	toks.Advance()

	for toks.Current().Is(slang.TokenCOMMA) {
		toks.Advance()

		if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
			return false
		}

		vars = append(vars, toks.Current().Value())

		toks.Advance()
	}

	if !p.ExpectToken(toks.Current(), slang.TokenIN) {
		return false
	}

	toks.Advance()

	exprParser := NewExprParser(parsing.PrecedenceLOWEST)

	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	bodyParser := NewCondBodyParser()

	if !p.RunSubParser(toks, bodyParser) {
		return false
	}

	p.Stmt = statements.NewForEach(vars, exprParser.Expr, bodyParser.GetStatements())

	p.Stmt.SetComment(comment)

	return true
}
