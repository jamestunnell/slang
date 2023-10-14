package parser

import (
	"errors"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/expressions"
	"github.com/jamestunnell/slang/statements"
)

var errBadStatementStart = errors.New("bad statment start")

func (p *Parser) parseStatement() slang.Statement {
	var s slang.Statement

	switch p.curToken.Info.Type() {
	case slang.TokenRETURN:
		s = p.parseRetStatement()
	case slang.TokenIDENT:
		if p.peekTokenIs(slang.TokenASSIGN) {
			s = p.parseAssignStatement()
		} else {
			s = p.parseExprStatement()
		}
	default:
		s = p.parseExprStatement()
	}

	return s
}

func (p *Parser) parseRetStatement() slang.Statement {
	p.nextToken()

	expr := p.parseExpression(PrecedenceLOWEST)

	return statements.NewReturn(expr)
}

func (p *Parser) parseAssignStatement() slang.Statement {
	ident := expressions.NewIdentifier(p.curToken.Info.Value())

	p.nextToken()

	p.nextToken()

	expr := p.parseExpression(PrecedenceLOWEST)

	return statements.NewAssign(ident, expr)
}

func (p *Parser) parseExprStatement() slang.Statement {
	expr := p.parseExpression(PrecedenceLOWEST)

	return statements.NewExpression(expr)
}
