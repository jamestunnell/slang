package parsing

import (
	"fmt"
	"strconv"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/customerrs"
)

func (p *ExprParser) parseExpression(toks slang.TokenSeq, prec Precedence) slang.Expression {
	prefixParse, foundPrefix := p.findPrefixParseFn(toks.Current().Type())
	if !foundPrefix {
		err := customerrs.NewErrMissingPrefixParseFn(toks.Current().Type())

		p.errors = append(p.errors, NewParseError(err, toks.Current()))

		return nil
	}

	leftExpr := prefixParse(toks)

	for prec < TokenPrecedence(toks.Current().Type()) {
		infix, foundInfix := p.findInfixParseFn(toks.Current().Type())
		if !foundInfix {
			break
		}

		leftExpr = infix(toks, leftExpr)
	}

	return leftExpr
}

func (p *ExprParser) parseGroupedExpression(toks slang.TokenSeq) slang.Expression {
	toks.Advance()

	expr := p.parseExpression(toks, PrecedenceLOWEST)

	if !p.ExpectToken(toks.Current(), slang.TokenRPAREN) {
		return nil
	}

	toks.Advance()

	return expr
}

func (p *ExprParser) parseArray(toks slang.TokenSeq) slang.Expression {
	toks.Advance()

	// check for empty array
	if toks.Current().Is(slang.TokenRBRACKET) {
		return expressions.NewArray()
	}

	elems := []slang.Expression{
		p.parseExpression(toks, PrecedenceLOWEST),
	}

	for toks.Next().Is(slang.TokenCOMMA) {
		toks.Advance()
		toks.Advance()

		elems = append(elems, p.parseExpression(toks, PrecedenceLOWEST))
	}

	if !p.ExpectToken(toks.Next(), slang.TokenRBRACKET) {
		return nil
	}

	toks.Advance()

	return expressions.NewArray(elems...)
}

func (p *ExprParser) parseIfExpression(toks slang.TokenSeq) slang.Expression {
	toks.Advance()

	cond := p.parseExpression(toks, PrecedenceLOWEST)

	conseqParser := NewCondBodyParser()
	if !p.RunSubParser(toks, conseqParser) {
		return nil
	}

	if !toks.Current().Is(slang.TokenELSE) {
		return expressions.NewIf(cond, conseqParser.Statements)
	}

	toks.Advance()

	alternParser := NewCondBodyParser()
	if !p.RunSubParser(toks, alternParser) {
		return nil
	}

	return expressions.NewIfElse(cond, conseqParser.Statements, alternParser.Statements)
}

func (p *ExprParser) parseFuncLiteral(toks slang.TokenSeq) slang.Expression {
	toks.Advance()

	sigParser := NewFuncSignatureParser()
	if !p.RunSubParser(toks, sigParser) {
		return nil
	}

	bodyParser := NewFuncBodyParser()
	if !p.RunSubParser(toks, bodyParser) {
		return nil
	}

	fn := ast.NewFunction(
		sigParser.Params, sigParser.ReturnTypes, bodyParser.Statements...)

	return expressions.NewFunc(fn)
}

func (p *ExprParser) parseIdentifier(toks slang.TokenSeq) slang.Expression {
	name := toks.Current().Value()

	toks.Advance()

	return expressions.NewIdentifier(name)
}

func (p *ExprParser) parseTrue(toks slang.TokenSeq) slang.Expression {
	toks.Advance()

	return expressions.NewBool(true)
}

func (p *ExprParser) parseFalse(toks slang.TokenSeq) slang.Expression {
	toks.Advance()

	return expressions.NewBool(false)
}

func (p *ExprParser) parseNegative(toks slang.TokenSeq) slang.Expression {
	toks.Advance()

	return expressions.NewNegative(p.parseExpression(toks, PrecedencePREFIX))
}

func (p *ExprParser) parseNot(toks slang.TokenSeq) slang.Expression {
	toks.Advance()

	return expressions.NewNot(p.parseExpression(toks, PrecedencePREFIX))
}

func (p *ExprParser) parseInteger(toks slang.TokenSeq) slang.Expression {
	str := toks.Current().Value()

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s' as int: %w", str, err)

		p.errors = append(p.errors, NewParseError(err, toks.Current()))

		return nil
	}

	toks.Advance()

	return expressions.NewInteger(i)
}

func (p *ExprParser) parseString(toks slang.TokenSeq) slang.Expression {
	strExprs := []slang.Expression{expressions.NewString(toks.Current().Value())}

	toks.Advance()

	for toks.Current().Is(slang.TokenDOLLARLBRACE) {
		toks.AdvanceSkip(slang.TokenNEWLINE)

		expr := p.parseExpression(toks, PrecedenceLOWEST)
		if expr == nil {
			return nil
		}

		strExprs = append(strExprs, expr)

		if !p.ExpectToken(toks.Current(), slang.TokenRBRACE) {
			return nil
		}

		toks.Advance()

		if !p.ExpectToken(toks.Current(), slang.TokenSTRING) {
			return nil
		}

		strExprs = append(strExprs, expressions.NewString(toks.Current().Value()))

		toks.Advance()
	}

	if len(strExprs) == 1 {
		return strExprs[0]
	}

	return expressions.NewConcat(strExprs...)
}

func (p *ExprParser) parseVerbatimString(toks slang.TokenSeq) slang.Expression {
	val := toks.Current().Value()

	toks.Advance()

	return expressions.NewString(val)
}

func (p *ExprParser) parseFloat(toks slang.TokenSeq) slang.Expression {
	str := toks.Current().Value()

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s' as float: %w", str, err)

		p.errors = append(p.errors, NewParseError(err, toks.Current()))

		return nil
	}

	toks.Advance()

	return expressions.NewFloat(f)
}

func (p *ExprParser) parseMemberAccess(toks slang.TokenSeq, object slang.Expression) slang.Expression {
	toks.Advance() // past the DOT

	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return nil
	}

	member := toks.Current().Value()

	toks.Advance()

	return expressions.NewMemberAccess(object, member)
}

func (p *ExprParser) parseCall(toks slang.TokenSeq, fn slang.Expression) slang.Expression {
	toks.Advance() // past the LPAREN

	args, ok := p.parseCallArgs(toks)
	if !ok {
		return nil
	}

	return expressions.NewCall(fn, args...)
}

func (p *ExprParser) parseIndex(toks slang.TokenSeq, ary slang.Expression) slang.Expression {
	toks.Advance()

	indexExprParser := NewExprParser(PrecedenceLOWEST)
	if !p.RunSubParser(toks, indexExprParser) {
		return nil
	}

	toks.Advance()

	if !p.ExpectToken(toks.Current(), slang.TokenRBRACKET) {
		return nil
	}

	return expressions.NewIndex(ary, indexExprParser.Expr)
}

func (p *ExprParser) parseCallArgs(
	toks slang.TokenSeq,
) ([]*expressions.Arg, bool) {
	if toks.Current().Is(slang.TokenRPAREN) {
		toks.Advance()

		return []*expressions.Arg{}, true
	}

	args := []*expressions.Arg{}

	addArg := func() bool {
		var nameTok *slang.Token

		if toks.Current().Is(slang.TokenSYMBOL) && toks.Next().Is(slang.TokenCOLON) {
			nameTok = toks.Current()

			toks.Advance()
			toks.Advance()
		}

		argExpr := p.parseExpression(toks, PrecedenceLOWEST)

		if argExpr == nil {
			return false
		}

		if nameTok == nil {
			args = append(args, expressions.NewPositionalArg(argExpr))

			return true
		}

		arg := expressions.NewKeywordArg(nameTok.Value(), argExpr)

		args = append(args, arg)

		return true
	}

	if !addArg() {
		return nil, false
	}

	for toks.Current().Is(slang.TokenCOMMA) {
		toks.Advance()

		if !addArg() {
			return nil, false
		}
	}

	if !p.ExpectToken(toks.Current(), slang.TokenRPAREN) {
		return nil, false
	}

	toks.Advance()

	return args, true
}

func (p *ExprParser) parseAdd(toks slang.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewAdd)
}

func (p *ExprParser) parseSubtract(toks slang.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewSubtract)
}

func (p *ExprParser) parseMultiply(toks slang.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewMultiply)
}

func (p *ExprParser) parseDivide(toks slang.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewDivide)
}

func (p *ExprParser) parseEqual(toks slang.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewEqual)
}

func (p *ExprParser) parseNotEqual(toks slang.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewNotEqual)
}

func (p *ExprParser) parseLess(toks slang.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewLess)
}

func (p *ExprParser) parseLessEqual(toks slang.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewLessEqual)
}

func (p *ExprParser) parseGreater(toks slang.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewGreater)
}

func (p *ExprParser) parseGreaterEqual(toks slang.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewGreaterEqual)
}

type newInfixExprFn func(left, right slang.Expression) slang.Expression

func (p *ExprParser) parseInfixExpr(toks slang.TokenSeq, left slang.Expression, fn newInfixExprFn) slang.Expression {
	prec := TokenPrecedence(toks.Current().Type())

	toks.Advance()

	right := p.parseExpression(toks, prec)

	return fn(left, right)
}
