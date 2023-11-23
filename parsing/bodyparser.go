package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type BodyParser struct {
	*ParserBase

	ParseStatement func(slang.TokenSeq) bool
	Statements     []slang.Statement
}

func NewBodyParser(parseStatement func(slang.TokenSeq) bool) *BodyParser {
	return &BodyParser{
		ParserBase:     NewParserBase(),
		ParseStatement: parseStatement,
		Statements:     []slang.Statement{},
	}
}

func (p *BodyParser) Run(toks slang.TokenSeq) {
	p.Statements = []slang.Statement{}

	if !p.ExpectToken(toks.Current(), slang.TokenLBRACE) {
		return
	}

	toks.AdvanceSkip(slang.TokenNEWLINE)

	for !toks.Current().Is(slang.TokenRBRACE) {
		if !p.ParseStatement(toks) {
			return
		}

		toks.Skip(slang.TokenNEWLINE)
	}

	toks.Advance()
}

func (p *BodyParser) ParseReturn(toks slang.TokenSeq) bool {
	toks.Advance()

	exprParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	p.Statements = append(p.Statements, statements.NewReturn(exprParser.Expr))

	return true
}

func (p *BodyParser) ParseAssign(toks slang.TokenSeq) bool {
	names := []string{toks.Current().Value()}

	toks.Advance()

	for toks.Current().Is(slang.TokenCOMMA) {
		toks.Advance()

		if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
			return false
		}

		names = append(names, toks.Current().Value())

		toks.Advance()
	}

	if !p.ExpectToken(toks.Current(), slang.TokenASSIGN) {
		return false
	}

	exprParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	p.Statements = append(p.Statements, statements.NewAssign(exprParser.Expr, names...))

	return true
}

func (p *BodyParser) ParseExpression(toks slang.TokenSeq) bool {
	exprParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	p.Statements = append(p.Statements, statements.NewExpression(exprParser.Expr))

	return true
}
