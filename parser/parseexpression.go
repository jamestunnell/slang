package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/expressions"
	"github.com/jamestunnell/slang/statements"
)

var errBadExpressionStart = errors.New("bad expression start")

func (p *Parser) parseExpression(prec Precedence) slang.Expression {
	prefixParse, found := p.prefixParseFns[p.curToken.Info.Type()]
	if !found {
		err := NewErrMissingPrefixParseFn(p.curToken.Info.Type())

		p.Errors = append(p.Errors, p.NewParseErr(err))

		return nil
	}
	leftExp := prefixParse()

	for prec < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Info.Type()]
		if infix == nil {
			return leftExp
		}

		p.nextTokenSkipComments()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseGroupedExpression() slang.Expression {
	p.nextTokenSkipComments()

	exp := p.parseExpression(PrecedenceLOWEST)

	if !p.expectPeekAndAdvance(slang.TokenRPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseArray() slang.Expression {
	elems := []slang.Expression{}

	noElems := p.peekTokenIs(slang.TokenRBRACKET)

	p.nextTokenSkipComments()

	if noElems {
		return expressions.NewArray(elems...)
	}

	elems = append(elems, p.parseExpression(PrecedenceLOWEST))

	for p.peekTokenIs(slang.TokenCOMMA) {
		p.nextTokenSkipComments()
		p.nextTokenSkipComments()

		elems = append(elems, p.parseExpression(PrecedenceLOWEST))
	}

	if !p.expectPeekAndAdvance(slang.TokenRBRACKET) {
		return nil
	}

	return expressions.NewArray(elems...)
}

func (p *Parser) parseIfExpression() slang.Expression {
	p.nextTokenSkipComments()

	cond := p.parseExpression(PrecedenceLOWEST)

	if !p.expectPeekAndAdvance(slang.TokenLBRACE) {
		return nil
	}

	conseq := p.parseBlock()

	var altern *statements.Block

	if p.peekTokenIs(slang.TokenELSE) {
		p.nextTokenSkipComments()

		if !p.expectPeekAndAdvance(slang.TokenLBRACE) {
			return nil
		}

		altern = p.parseBlock()
	}

	if altern == nil {
		return expressions.NewIf(cond, conseq)
	}

	return expressions.NewIfElse(cond, conseq, altern)
}

func (p *Parser) parseFuncLiteral() slang.Expression {
	if !p.expectPeekAndAdvance(slang.TokenLPAREN) {
		return nil
	}

	params := p.parseFuncParams()

	if !p.expectPeekAndAdvance(slang.TokenLBRACE) {
		return nil
	}

	body := p.parseBlock()

	return expressions.NewFunctionLiteral(params, body)
}

func (p *Parser) parseFuncParams() []*expressions.Identifier {
	params := []*expressions.Identifier{}
	addCur := func() {
		params = append(params,
			expressions.NewIdentifier(p.curToken.Info.Value()))
	}

	if p.peekTokenIs(slang.TokenRPAREN) {
		p.nextTokenSkipComments() // end on RPAREN

		return params
	}

	if !p.expectPeekAndAdvance(slang.TokenSYMBOL) {
		return nil
	}

	addCur()

	for p.peekTokenIs(slang.TokenCOMMA) {
		p.nextTokenSkipComments()

		if !p.expectPeekAndAdvance(slang.TokenSYMBOL) {
			return nil
		}

		addCur()
	}

	if !p.expectPeekAndAdvance(slang.TokenRPAREN) {
		return nil
	}

	return params
}

func (p *Parser) parseBlock() *statements.Block {
	p.nextTokenSkipComments()

	stmts := p.parseBlockStatements()

	return statements.NewBlock(stmts...)
}

func (p *Parser) parseIdentifier() slang.Expression {
	return expressions.NewIdentifier(p.curToken.Info.Value())
}

func (p *Parser) parseTrue() slang.Expression {
	return expressions.NewBool(true)
}

func (p *Parser) parseFalse() slang.Expression {
	return expressions.NewBool(false)
}

func (p *Parser) parseNegative() slang.Expression {
	return p.parsePrefixExpr(expressions.NewNegative)
}

func (p *Parser) parseNot() slang.Expression {
	return p.parsePrefixExpr(expressions.NewNot)
}

type newPrefixExprFn func(slang.Expression) slang.Expression

func (p *Parser) parsePrefixExpr(fn newPrefixExprFn) slang.Expression {
	p.nextTokenSkipComments()

	val := p.parseExpression(PrecedencePREFIX)

	return fn(val)
}

func (p *Parser) parseInteger() slang.Expression {
	str := p.curToken.Info.Value()

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s' as int: %w", str, err)

		p.Errors = append(p.Errors, p.NewParseErr(err))

		return nil
	}

	return expressions.NewInteger(i)
}

func (p *Parser) parseString() slang.Expression {
	return expressions.NewString(p.curToken.Info.Value())
}

func (p *Parser) parseFloat() slang.Expression {
	str := p.curToken.Info.Value()

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s' as float: %w", str, err)

		p.Errors = append(p.Errors, p.NewParseErr(err))

		return nil
	}

	return expressions.NewFloat(f)
}

func (p *Parser) parseMethodCall(obj slang.Expression) slang.Expression {
	if !p.expectPeekAndAdvance(slang.TokenSYMBOL) {
		return nil
	}

	methodName := expressions.NewIdentifier(p.curToken.Info.Value())
	args := []slang.Expression{}

	if p.peekTokenIs(slang.TokenLPAREN) {
		p.nextTokenSkipComments()

		args = p.parseCallArgs()
	}

	return expressions.NewMethodCall(obj, methodName, args...)
}

func (p *Parser) parseFunctionCall(fn slang.Expression) slang.Expression {
	args := p.parseCallArgs()

	return expressions.NewFunctionCall(fn, args...)
}

func (p *Parser) parseIndex(ary slang.Expression) slang.Expression {
	p.nextTokenSkipComments()

	idx := p.parseExpression(PrecedenceLOWEST)

	if !p.expectPeekAndAdvance(slang.TokenRBRACKET) {
		return nil
	}

	return expressions.NewIndex(ary, idx)
}

func (p *Parser) parseCallArgs() []slang.Expression {
	args := []slang.Expression{}

	noArgs := p.peekTokenIs(slang.TokenRPAREN)

	p.nextTokenSkipComments()

	if noArgs {
		return args
	}

	args = append(args, p.parseExpression(PrecedenceLOWEST))

	for p.peekTokenIs(slang.TokenCOMMA) {
		p.nextTokenSkipComments()
		p.nextTokenSkipComments()

		args = append(args, p.parseExpression(PrecedenceLOWEST))
	}

	if !p.expectPeekAndAdvance(slang.TokenRPAREN) {
		return nil
	}

	return args
}

func (p *Parser) parseAdd(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewAdd)
}

func (p *Parser) parseSubtract(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewSubtract)
}

func (p *Parser) parseMultiply(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewMultiply)
}

func (p *Parser) parseDivide(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewDivide)
}

func (p *Parser) parseEqual(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewEqual)
}

func (p *Parser) parseNotEqual(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewNotEqual)
}

func (p *Parser) parseLess(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewLess)
}

func (p *Parser) parseLessEqual(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewLessEqual)
}

func (p *Parser) parseGreater(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewGreater)
}

func (p *Parser) parseGreaterEqual(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewGreaterEqual)
}

type newInfixExprFn func(left, right slang.Expression) slang.Expression

func (p *Parser) parseInfixExpr(left slang.Expression, fn newInfixExprFn) slang.Expression {
	prec := p.curPrecedence()

	p.nextTokenSkipComments()

	right := p.parseExpression(prec)

	return fn(left, right)
}
