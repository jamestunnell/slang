package parsing

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/customerrs"
)

var errNonIdentInTypeChain = errors.New("not a valid type chain, non-identifier found")

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

// func (p *ExprParser) parseArrayAuto(toks slang.TokenSeq) slang.Expression {
// 	toks.Advance() // past [

// 	// // type includes brackets
// 	// typ, ok := p.ParseArrayType(toks)
// 	// if !ok {
// 	// 	return nil
// 	// }

// 	// if !p.ExpectToken(toks.Current(), slang.TokenLBRACE) {
// 	// 	return nil
// 	// }

// 	// toks.Advance()

// 	vals := []slang.Expression{}

// 	// check for empty array
// 	if toks.Current().Is(slang.TokenRBRACKET) {
// 		return expressions.NewArrayAuto()
// 	}

// 	// first value
// 	v := p.parseExpression(toks, PrecedenceLOWEST)

// 	vals = append(vals, v)

// 	// more values
// 	for toks.Current().Is(slang.TokenCOMMA) {
// 		toks.Advance()

// 		v := p.parseExpression(toks, PrecedenceLOWEST)

// 		vals = append(vals, v)
// 	}

// 	if !p.ExpectToken(toks.Current(), slang.TokenRBRACKET) {
// 		return nil
// 	}

// 	toks.Advance()

// 	return expressions.NewArrayAuto(vals...)
// }

// func (p *ExprParser) parseMapAuto(toks slang.TokenSeq) slang.Expression {
// 	toks.Advance() // past <

// 	// typ, ok := p.ParseMapType(toks)
// 	// if !ok {
// 	// 	return nil
// 	// }

// 	// if !p.ExpectToken(toks.Current(), slang.TokenLBRACE) {
// 	// 	return nil
// 	// }

// 	keys := []slang.Expression{}
// 	vals := []slang.Expression{}

// 	// check for empty map
// 	if toks.Current().Is(slang.TokenGREATER) {
// 		return expressions.NewMapAuto(keys, vals)
// 	}

// 	// first KV pair
// 	k, v, ok := p.parseMapKVPair(toks)
// 	if !ok {
// 		return nil
// 	}

// 	keys = append(keys, k)
// 	vals = append(vals, v)

// 	// more KV pairs
// 	for toks.Current().Is(slang.TokenCOMMA) {
// 		toks.Advance()

// 		k, v, ok := p.parseMapKVPair(toks)
// 		if !ok {
// 			return nil
// 		}

// 		keys = append(keys, k)
// 		vals = append(vals, v)
// 	}

// 	if !p.ExpectToken(toks.Current(), slang.TokenGREATER) {
// 		return nil
// 	}

// 	return expressions.NewMapAuto(keys, vals)
// }

// func (p *ExprParser) parseMapKVPair(toks slang.TokenSeq) (key, val slang.Expression, ok bool) {
// 	k := p.parseExpression(toks, PrecedenceLOWEST)

// 	toks.Skip(slang.TokenNEWLINE)

// 	if !p.ExpectToken(toks.Current(), slang.TokenCOLON) {
// 		return nil, nil, false
// 	}

// 	toks.AdvanceSkip(slang.TokenNEWLINE)

// 	v := p.parseExpression(toks, PrecedenceLOWEST)

// 	return k, v, true
// }

func (p *ExprParser) parseFuncAnon(toks slang.TokenSeq) slang.Expression {
	toks.AdvanceSkip(slang.TokenNEWLINE) // past func

	sigParser := NewFuncSignatureParser()
	if !p.RunSubParser(toks, sigParser) {
		return nil
	}

	bodyParser := NewFuncBodyParser()
	if !p.RunSubParser(toks, bodyParser) {
		return nil
	}

	bodyStatements := make([]slang.Statement, len(bodyParser.Statements))

	for i, stmt := range bodyParser.Statements {
		bodyStatements[i] = stmt
	}

	return expressions.NewFunc(
		sigParser.InParams, sigParser.OutParams, bodyStatements...)
}

func (p *ExprParser) parseIdentifier(toks slang.TokenSeq) slang.Expression {
	name := toks.Current().Value()

	toks.Advance()

	return expressions.NewIdentifier(name)
}

func (p *ExprParser) parseBoolVal(toks slang.TokenSeq) slang.Expression {
	str := toks.Current().Value()

	b, err := strconv.ParseBool(str)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s' as bool: %w", str, err)

		p.errors = append(p.errors, NewParseError(err, toks.Current()))

		return nil
	}

	toks.Advance()

	return expressions.NewBool(b)
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

func (p *ExprParser) parseIntVal(toks slang.TokenSeq) slang.Expression {
	str := toks.Current().Value()

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s' as int: %w", str, err)

		p.errors = append(p.errors, NewParseError(err, toks.Current()))

		return nil
	}

	toks.Advance()

	return expressions.NewInt(i)
}

func (p *ExprParser) parseStrVal(toks slang.TokenSeq) slang.Expression {
	strExprs := []slang.Expression{expressions.NewStr(toks.Current().Value())}

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

		if !p.ExpectToken(toks.Current(), slang.TokenSTRVAL) {
			return nil
		}

		strExprs = append(strExprs, expressions.NewStr(toks.Current().Value()))

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

	return expressions.NewStr(val)
}

func (p *ExprParser) parseFloatVal(toks slang.TokenSeq) slang.Expression {
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

func (p *ExprParser) parseAccessMember(toks slang.TokenSeq, object slang.Expression) slang.Expression {
	toks.Advance() // past the DOT

	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return nil
	}

	member := toks.Current().Value()

	toks.Advance()

	return expressions.NewAccessMember(object, member)
}

func (p *ExprParser) parseInvoke(toks slang.TokenSeq, subject slang.Expression) slang.Expression {
	toks.AdvanceSkip(slang.TokenNEWLINE) // past the LPAREN and any newlines

	// check for no args
	if toks.Current().Is(slang.TokenRPAREN) {
		toks.Advance()

		return expressions.NewInvoke(subject)
	}

	args := []*expressions.InvokeArg{}
	addArg := func() bool {
		var nameTok *slang.Token

		if toks.Current().Is(slang.TokenSYMBOL) && toks.Next().Is(slang.TokenCOLON) {
			nameTok = toks.Current()

			toks.Advance()
			toks.AdvanceSkip(slang.TokenNEWLINE)
		}

		val := p.parseExpression(toks, PrecedenceLOWEST)
		if val == nil {
			return false
		}

		toks.Skip(slang.TokenNEWLINE)

		var name string

		if nameTok != nil {
			name = nameTok.Value()
		}

		args = append(args, &expressions.InvokeArg{Name: name, Value: val})

		return true
	}

	if !addArg() {
		return nil
	}

	for toks.Current().Is(slang.TokenCOMMA, slang.TokenNEWLINE) {
		toks.AdvanceSkip(slang.TokenNEWLINE)

		if !addArg() {
			return nil
		}
	}

	toks.Skip(slang.TokenNEWLINE)

	if !p.ExpectToken(toks.Current(), slang.TokenRPAREN) {
		return nil
	}

	toks.Advance()

	return expressions.NewInvoke(subject, args...)
}

func (p *ExprParser) parseStructAnon(toks slang.TokenSeq) slang.Expression {
	toks.Advance() // past struct

	sigParser := NewDataSignatureParser()
	if !p.RunSubParser(toks, sigParser) {
		return nil
	}

	return expressions.NewStruct(sigParser.NameTypes...)
}

// func (p *ExprParser) parseAccessElem(toks slang.TokenSeq, ary slang.Expression) slang.Expression {
// 	toks.Advance()

// 	keyExpr := p.parseExpression(toks, PrecedenceLOWEST)
// 	if keyExpr == nil {
// 		return nil
// 	}

// 	if !p.ExpectToken(toks.Current(), slang.TokenRBRACKET) {
// 		return nil
// 	}

// 	toks.Advance()

// 	return expressions.NewAccessElem(ary, keyExpr)
// }

func (p *ExprParser) parseAnd(toks slang.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewAnd)
}

func (p *ExprParser) parseOr(toks slang.TokenSeq, left slang.Expression) slang.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewOr)
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

type newInfixExprFn func(left, right slang.Expression) *expressions.Expression

func (p *ExprParser) parseInfixExpr(toks slang.TokenSeq, left slang.Expression, fn newInfixExprFn) slang.Expression {
	prec := TokenPrecedence(toks.Current().Type())

	toks.Advance()

	right := p.parseExpression(toks, prec)

	return fn(left, right)
}
