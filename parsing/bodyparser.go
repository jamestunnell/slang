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

func (p *BodyParser) GetStatements() []slang.Statement {
	return p.Statements
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

		if !p.ExpectToken(toks.Current(), slang.TokenNEWLINE, slang.TokenRBRACE) {
			return
		}

		toks.Skip(slang.TokenNEWLINE)
	}

	toks.Advance()
}

func (p *BodyParser) ParseReturnStatment(toks slang.TokenSeq) bool {
	toks.Advance()

	exprParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	p.Statements = append(p.Statements, statements.NewReturn(exprParser.Expr))

	return true
}

func (p *BodyParser) ParseExpressionOrAssignStatement(toks slang.TokenSeq) bool {
	exprParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, exprParser) {
		return false
	}

	if toks.Current().Is(slang.TokenASSIGN) {
		toks.AdvanceSkip(slang.TokenNEWLINE)

		valueParser := NewExprParser(PrecedenceLOWEST)
		if !p.RunSubParser(toks, valueParser) {
			return false
		}

		p.Statements = append(p.Statements, statements.NewAssign(exprParser.Expr, valueParser.Expr))
	} else {
		p.Statements = append(p.Statements, statements.NewExpression(exprParser.Expr))
	}

	return true
}
