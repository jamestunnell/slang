package parsing

import (
	"fmt"
	"strconv"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/customerrs"
	"github.com/jamestunnell/slang/lexing"
)

func (p *ExprParser) parseExpression(toks lexing.TokenSeq, prec Precedence) slang.Expression {
	prefixParse, foundPrefix := p.findPrefixParseFn(toks.Current().Type)
	if !foundPrefix {
		err := customerrs.NewErrMissingPrefixParseFn(toks.Current().Type)

		p.errors = append(p.errors, NewParseError(err, toks.Current()))

		return nil
	}

	leftExpr := prefixParse(toks)

	for prec < TokenPrecedence(toks.Current().Type) {
		infix, foundInfix := p.findInfixParseFn(toks.Current().Type)
		if !foundInfix {
			break
		}

		leftExpr = infix(toks, leftExpr)
	}

	return leftExpr
}

func (p *ExprParser) parseGroupedExpression(toks lexing.TokenSeq) slang.Expression {
	toks.Advance()

	expr := p.parseExpression(toks, PrecedenceLOWEST)

	if !p.ExpectToken(toks.Current(), lexing.TokenRPAREN) {
		return nil
	}

	toks.Advance()

	return expr
}

func (p *ExprParser) parseArrayOrMap(toks lexing.TokenSeq) slang.Expression {
	if toks.Next().Is(lexing.TokenRBRACKET) {
		return p.parseArray(toks)
	}

	return p.parseMap(toks)
}

func (p *ExprParser) parseArray(toks lexing.TokenSeq) slang.Expression {
	// type includes brackets
	typ, ok := p.ParseArrayType(toks)
	if !ok {
		return nil
	}

	if !p.ExpectToken(toks.Current(), lexing.TokenLBRACE) {
		return nil
	}

	toks.Advance()

	vals := []slang.Expression{}

	// check for empty array
	if toks.Current().Is(lexing.TokenRBRACE) {
		return expressions.NewArray(typ.ValueType)
	}

	// first value
	v := p.parseExpression(toks, PrecedenceLOWEST)

	vals = append(vals, v)

	// more values
	for toks.Current().Is(lexing.TokenCOMMA) {
		toks.Advance()

		v := p.parseExpression(toks, PrecedenceLOWEST)

		vals = append(vals, v)
	}

	if !p.ExpectToken(toks.Current(), lexing.TokenRBRACE) {
		return nil
	}

	toks.Advance()

	return expressions.NewArray(typ.ValueType, vals...)
}

func (p *ExprParser) parseMap(toks lexing.TokenSeq) slang.Expression {
	typ, ok := p.ParseMapType(toks)
	if !ok {
		return nil
	}

	if !p.ExpectToken(toks.Current(), lexing.TokenLBRACE) {
		return nil
	}

	toks.Advance()

	keys := []slang.Expression{}
	vals := []slang.Expression{}

	// check for empty map
	if toks.Current().Is(lexing.TokenRBRACKET) {
		return expressions.NewMap(typ.KeyType, keys, typ.ValueType, vals)
	}

	// first KV pair
	k, v, ok := p.parseMapKVPair(toks)
	if !ok {
		return nil
	}

	keys = append(keys, k)
	vals = append(vals, v)

	// more KV pairs
	for toks.Current().Is(lexing.TokenCOMMA) {
		toks.Advance()

		k, v, ok := p.parseMapKVPair(toks)
		if !ok {
			return nil
		}

		keys = append(keys, k)
		vals = append(vals, v)
	}

	if !p.ExpectToken(toks.Current(), lexing.TokenRBRACE) {
		return nil
	}

	return expressions.NewMap(typ.KeyType, keys, typ.ValueType, vals)
}

func (p *ExprParser) parseMapKVPair(toks lexing.TokenSeq) (key, val slang.Expression, ok bool) {
	k := p.parseExpression(toks, PrecedenceLOWEST)

	if !p.ExpectToken(toks.Current(), lexing.TokenCOLON) {
		return nil, nil, false
	}

	toks.Advance()

	v := p.parseExpression(toks, PrecedenceLOWEST)

	return k, v, true
}

func (p *ExprParser) parseFuncLiteral(toks lexing.TokenSeq) slang.Expression {
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

func (p *ExprParser) parseIdentifier(toks lexing.TokenSeq) slang.Expression {
	name := toks.Current().Value

	toks.Advance()

	return expressions.NewIdentifier(name)
}

func (p *ExprParser) parseTrue(toks lexing.TokenSeq) slang.Expression {
	toks.Advance()

	return expressions.NewBool(true)
}

func (p *ExprParser) parseFalse(toks lexing.TokenSeq) slang.Expression {
	toks.Advance()

	return expressions.NewBool(false)
}

func (p *ExprParser) parseNegative(toks lexing.TokenSeq) slang.Expression {
	toks.Advance()

	return expressions.NewNegative(p.parseExpression(toks, PrecedencePREFIX))
}

func (p *ExprParser) parseNot(toks lexing.TokenSeq) slang.Expression {
	toks.Advance()

	return expressions.NewNot(p.parseExpression(toks, PrecedencePREFIX))
}

func (p *ExprParser) parseInteger(toks lexing.TokenSeq) slang.Expression {
	str := toks.Current().Value

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s' as int: %w", str, err)

		p.errors = append(p.errors, NewParseError(err, toks.Current()))

		return nil
	}

	toks.Advance()

	return expressions.NewInteger(i)
}

func (p *ExprParser) parseString(toks lexing.TokenSeq) slang.Expression {
	strExprs := []slang.Expression{expressions.NewString(toks.Current().Value)}

	toks.Advance()

	for toks.Current().Is(lexing.TokenDOLLARLBRACE) {
		toks.AdvanceSkip(lexing.TokenNEWLINE)

		expr := p.parseExpression(toks, PrecedenceLOWEST)
		if expr == nil {
			return nil
		}

		strExprs = append(strExprs, expr)

		if !p.ExpectToken(toks.Current(), lexing.TokenRBRACE) {
			return nil
		}

		toks.Advance()

		if !p.ExpectToken(toks.Current(), lexing.TokenSTRING) {
			return nil
		}

		strExprs = append(strExprs, expressions.NewString(toks.Current().Value))

		toks.Advance()
	}

	if len(strExprs) == 1 {
		return strExprs[0]
	}

	return expressions.NewConcat(strExprs...)
}

func (p *ExprParser) parseVerbatimString(toks lexing.TokenSeq) slang.Expression {
	val := toks.Current().Value

	toks.Advance()

	return expressions.NewString(val)
}

func (p *ExprParser) parseFloat(toks lexing.TokenSeq) slang.Expression {
	str := toks.Current().Value

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s' as float: %w", str, err)

		p.errors = append(p.errors, NewParseError(err, toks.Current()))

		return nil
	}

	toks.Advance()

	return expressions.NewFloat(f)
}

func (p *ExprParser) parseAccessMember(toks lexing.TokenSeq, object slang.Expression) slang.Expression {
	toks.Advance() // past the DOT

	if !p.ExpectToken(toks.Current(), lexing.TokenSYMBOL) {
		return nil
	}

	member := toks.Current().Value

	toks.Advance()

	return expressions.NewAccessMember(object, member)
}

func (p *ExprParser) parseCall(toks lexing.TokenSeq, fn slang.Expression) slang.Expression {
	toks.Advance() // past the LPAREN

	args, ok := p.parseCallArgs(toks)
	if !ok {
		return nil
	}

	return expressions.NewCall(fn, args...)
}

func (p *ExprParser) parseAccessElem(toks lexing.TokenSeq, ary slang.Expression) slang.Expression {
	toks.Advance()

	keyExpr := p.parseExpression(toks, PrecedenceLOWEST)
	if keyExpr == nil {
		return nil
	}

	if !p.ExpectToken(toks.Current(), lexing.TokenRBRACKET) {
		return nil
	}

	toks.Advance()

	return expressions.NewAccessElem(ary, keyExpr)
}

func (p *ExprParser) parseCallArgs(
	toks lexing.TokenSeq,
) ([]*expressions.Arg, bool) {
	if toks.Current().Is(lexing.TokenRPAREN) {
		toks.Advance()

		return []*expressions.Arg{}, true
	}

	args := []*expressions.Arg{}

	addArg := func() bool {
		var nameTok *lexing.Token

		if toks.Current().Is(lexing.TokenSYMBOL) && toks.Next().Is(lexing.TokenCOLON) {
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

		arg := expressions.NewKeywordArg(nameTok.Value, argExpr)

		args = append(args, arg)

		return true
	}

	if !addArg() {
		return nil, false
	}

	for toks.Current().Is(lexing.TokenCOMMA) {
		toks.Advance()

		if !addArg() {
			return nil, false
		}
	}

	if !p.ExpectToken(toks.Current(), lexing.TokenRPAREN) {
		return nil, false
	}

	toks.Advance()

	return args, true
}

func (p *ExprParser) parseAnd(toks lexing.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewAnd)
}

func (p *ExprParser) parseOr(toks lexing.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewOr)
}

func (p *ExprParser) parseAdd(toks lexing.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewAdd)
}

func (p *ExprParser) parseSubtract(toks lexing.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewSubtract)
}

func (p *ExprParser) parseMultiply(toks lexing.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewMultiply)
}

func (p *ExprParser) parseDivide(toks lexing.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewDivide)
}

func (p *ExprParser) parseEqual(toks lexing.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewEqual)
}

func (p *ExprParser) parseNotEqual(toks lexing.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewNotEqual)
}

func (p *ExprParser) parseLess(toks lexing.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewLess)
}

func (p *ExprParser) parseLessEqual(toks lexing.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewLessEqual)
}

func (p *ExprParser) parseGreater(toks lexing.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewGreater)
}

func (p *ExprParser) parseGreaterEqual(toks lexing.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewGreaterEqual)
}

type newInfixExprFn func(left, right slang.Expression) slang.Expression

func (p *ExprParser) parseInfixExpr(toks lexing.TokenSeq, left slang.Expression, fn newInfixExprFn) slang.Expression {
	prec := TokenPrecedence(toks.Current().Type)

	toks.Advance()

	right := p.parseExpression(toks, prec)

	return fn(left, right)
}
