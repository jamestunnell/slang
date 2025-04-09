package parsers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/customerrs"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/sliceutil"
)

var errNonIdentInTypeChain = errors.New("not a valid type chain, non-identifier found")

func (p *ExprParser) parseExpression(
	toks parsing.TokenSeq,
	prec parsing.Precedence,
) *expressions.Expression {
	prefixParse, foundPrefix := p.findPrefixParseFn(toks.Current().Type())
	if !foundPrefix {
		err := customerrs.NewErrMissingPrefixParseFn(toks.Current().Type())

		p.AddError(parsing.NewParseError(err, toks.Current()))

		return nil
	}

	leftExpr := prefixParse(toks)

	for prec < parsing.TokenPrecedence(toks.Current().Type()) {
		infix, foundInfix := p.findInfixParseFn(toks.Current().Type())
		if !foundInfix {
			break
		}

		leftExpr = infix(toks, leftExpr)
	}

	return leftExpr
}

func (p *ExprParser) parseGroupedExpression(toks parsing.TokenSeq) *expressions.Expression {
	toks.Advance()

	expr := p.parseExpression(toks, parsing.PrecedenceLOWEST)

	if !p.ExpectToken(toks.Current(), slang.TokenRPAREN) {
		return nil
	}

	toks.Advance()

	return expr
}

var errEmptyAutoArrayOrMap = errors.New("empty auto array/map")

func (p *ExprParser) parseAutoArrayOrMap(toks parsing.TokenSeq) *expressions.Expression {
	_ = toks.AdvanceSkip(slang.TokenNEWLINE) // past the [ and any newlines

	// check for no args
	if toks.Current().Is(slang.TokenRBRACKET) {
		p.AddError(parsing.NewParseError(errEmptyAutoArrayOrMap, toks.Current()))

		return nil
	}

	expr := p.parseExpression(toks, parsing.PrecedenceLOWEST)

	if toks.Current().Is(slang.TokenCOLON) {
		toks.Advance()

		return p.finishParsingAutoMap(toks, expr)
	}

	return p.finishParsingAutoArray(toks, expr)
}

func (p *ExprParser) finishParsingAutoMap(
	toks parsing.TokenSeq,
	firstKey *expressions.Expression,
) *expressions.Expression {
	// finish first key-val pair
	firstVal := p.parseExpression(toks, parsing.PrecedenceLOWEST)
	if firstVal == nil {
		return nil
	}

	keys := []*expressions.Expression{firstKey}
	vals := []*expressions.Expression{firstVal}

	for {
		_ = toks.Skip(slang.TokenNEWLINE)

		if toks.Current().Is(slang.TokenRBRACKET) {
			toks.Advance()

			break
		}

		key := p.parseExpression(toks, parsing.PrecedenceLOWEST)
		if key == nil {
			return nil
		}

		if !p.ExpectToken(toks.Current(), slang.TokenCOLON) {
			return nil
		}

		toks.Advance()

		val := p.parseExpression(toks, parsing.PrecedenceLOWEST)
		if val == nil {
			return nil
		}

		keys = append(keys, key)
		vals = append(vals, val)
	}

	return expressions.NewMapAuto(keys, vals)
}

func (p *ExprParser) finishParsingAutoArray(
	toks parsing.TokenSeq,
	firstVal *expressions.Expression,
) *expressions.Expression {
	vals := []*expressions.Expression{firstVal}

	for {
		_ = toks.Skip(slang.TokenNEWLINE)

		if toks.Current().Is(slang.TokenRBRACKET) {
			toks.Advance()

			break
		}

		val := p.parseExpression(toks, parsing.PrecedenceLOWEST)
		if val == nil {
			return nil
		}

		vals = append(vals, val)
	}

	return expressions.NewArrayAuto(vals...)
}

// func (p *ExprParser) parseMapAuto(toks parsing.TokenSeq) *expressions.Expression {
// 	toks.Advance() // past <

// 	// typ, ok := p.ParseMapType(toks)
// 	// if !ok {
// 	// 	return nil
// 	// }

// 	// if !p.ExpectToken(toks.Current(), slang.TokenLBRACE) {
// 	// 	return nil
// 	// }

// 	keys := []*expressions.Expression{}
// 	vals := []*expressions.Expression{}

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

// func (p *ExprParser) parseMapKVPair(toks parsing.TokenSeq) (key, val *expressions.Expression, ok bool) {
// 	k := p.parseExpression(toks, PrecedenceLOWEST)

// 	toks.Skip(slang.TokenNEWLINE)

// 	if !p.ExpectToken(toks.Current(), slang.TokenCOLON) {
// 		return nil, nil, false
// 	}

// 	toks.AdvanceSkip(slang.TokenNEWLINE)

// 	v := p.parseExpression(toks, PrecedenceLOWEST)

// 	return k, v, true
// }

func (p *ExprParser) parseFuncAnon(toks parsing.TokenSeq) *expressions.Expression {
	toks.AdvanceSkip(slang.TokenNEWLINE) // past func

	sigParser := NewFuncSignatureParser()
	if !p.RunSubParser(toks, sigParser) {
		return nil
	}

	bodyParser := NewFuncBodyParser()
	if !p.RunSubParser(toks, bodyParser) {
		return nil
	}

	bodyStatements := sliceutil.Map(bodyParser.GetStatements(), func(s *statements.Statement) slang.Statement { return s })

	return expressions.NewFunc(
		sigParser.Inputs, sigParser.Outputs, bodyStatements...)
}

func (p *ExprParser) parseIdentifier(toks parsing.TokenSeq) *expressions.Expression {
	name := toks.Current().Value()

	toks.Advance()

	return expressions.NewIdentifier(name)
}

func (p *ExprParser) parseBoolVal(toks parsing.TokenSeq) *expressions.Expression {
	str := toks.Current().Value()

	b, err := strconv.ParseBool(str)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s' as bool: %w", str, err)

		p.AddError(parsing.NewParseError(err, toks.Current()))

		return nil
	}

	toks.Advance()

	return expressions.NewBool(b)
}

func (p *ExprParser) parseNegative(toks parsing.TokenSeq) *expressions.Expression {
	toks.Advance()

	return expressions.NewNegative(p.parseExpression(toks, parsing.PrecedencePREFIX))
}

func (p *ExprParser) parseNot(toks parsing.TokenSeq) *expressions.Expression {
	toks.Advance()

	return expressions.NewNot(p.parseExpression(toks, parsing.PrecedencePREFIX))
}

func (p *ExprParser) parseIntVal(toks parsing.TokenSeq) *expressions.Expression {
	str := toks.Current().Value()

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s' as int: %w", str, err)

		p.AddError(parsing.NewParseError(err, toks.Current()))

		return nil
	}

	toks.Advance()

	return expressions.NewInt(i)
}

func (p *ExprParser) parseStrVal(toks parsing.TokenSeq) *expressions.Expression {
	strExprs := []*expressions.Expression{expressions.NewStr(toks.Current().Value())}

	toks.Advance()

	for toks.Current().Is(slang.TokenDOLLARLBRACE) {
		toks.AdvanceSkip(slang.TokenNEWLINE)

		expr := p.parseExpression(toks, parsing.PrecedenceLOWEST)
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

func (p *ExprParser) parseVerbatimString(toks parsing.TokenSeq) *expressions.Expression {
	val := toks.Current().Value()

	toks.Advance()

	return expressions.NewStr(val)
}

func (p *ExprParser) parseFloatVal(toks parsing.TokenSeq) *expressions.Expression {
	str := toks.Current().Value()

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s' as float: %w", str, err)

		p.AddError(parsing.NewParseError(err, toks.Current()))

		return nil
	}

	toks.Advance()

	return expressions.NewFloat(f)
}

func (p *ExprParser) parseAccessMember(toks parsing.TokenSeq, object *expressions.Expression) *expressions.Expression {
	toks.Advance() // past the DOT

	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return nil
	}

	member := toks.Current().Value()

	toks.Advance()

	return expressions.NewAccessMember(object, member)
}

func (p *ExprParser) parseAnyInvokeArg(toks parsing.TokenSeq) (*expressions.InvokeArg, bool) {
	var nameTok *slang.Token

	if toks.Current().Is(slang.TokenSYMBOL) && toks.Next().Is(slang.TokenCOLON) {
		nameTok = toks.Current()

		toks.Advance()
		toks.AdvanceSkip(slang.TokenNEWLINE)
	}

	val := p.parseExpression(toks, parsing.PrecedenceLOWEST)
	if val == nil {
		return nil, false
	}

	var name string

	if nameTok != nil {
		name = nameTok.Value()
	}

	return &expressions.InvokeArg{Name: name, Value: val}, true
}

func (p *ExprParser) parseInvokePosArgs(toks parsing.TokenSeq) ([]*expressions.InvokeArg, bool) {
	args := []*expressions.InvokeArg{}

	for {
		toks.Skip(slang.TokenNEWLINE)

		if toks.Current().Is(slang.TokenRPAREN) {
			toks.Advance()

			break
		}

		val := p.parseExpression(toks, parsing.PrecedenceLOWEST)
		if val == nil {
			return []*expressions.InvokeArg{}, false
		}

		args = append(args, expressions.NewInvokeArgPos(val))
	}

	return args, true
}

func (p *ExprParser) parseInvokeKWArgs(toks parsing.TokenSeq) ([]*expressions.InvokeArg, bool) {
	args := []*expressions.InvokeArg{}

	for {
		toks.Skip(slang.TokenNEWLINE)

		if toks.Current().Is(slang.TokenRPAREN) {
			toks.Advance()

			break
		}

		if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
			return []*expressions.InvokeArg{}, false
		}

		name := toks.Current().Value()

		toks.Advance()

		if !p.ExpectToken(toks.Current(), slang.TokenCOLON) {
			return []*expressions.InvokeArg{}, false
		}

		toks.Advance()

		val := p.parseExpression(toks, parsing.PrecedenceLOWEST)
		if val == nil {
			return []*expressions.InvokeArg{}, false
		}

		args = append(args, expressions.NewInvokeArgKW(name, val))
	}

	return args, true
}

func (p *ExprParser) parseInvoke(toks parsing.TokenSeq, subject *expressions.Expression) *expressions.Expression {
	_ = toks.AdvanceSkip(slang.TokenNEWLINE) // past the ( and any newlines

	// check for no args
	if toks.Current().Is(slang.TokenRPAREN) {
		toks.Advance()

		return expressions.NewInvoke(subject)
	}

	_ = toks.Skip(slang.TokenNEWLINE)

	firstArg, ok := p.parseAnyInvokeArg(toks)
	if !ok {
		return nil
	}

	var moreArgs []*expressions.InvokeArg

	if firstArg.Name == "" {
		moreArgs, ok = p.parseInvokePosArgs(toks)
		if !ok {
			return nil
		}
	} else {
		moreArgs, ok = p.parseInvokeKWArgs(toks)
		if !ok {
			return nil
		}
	}

	args := append([]*expressions.InvokeArg{firstArg}, moreArgs...)

	return expressions.NewInvoke(subject, args...)
}

func (p *ExprParser) parseStructAnon(toks parsing.TokenSeq) *expressions.Expression {
	toks.Advance() // past struct

	sigParser := NewFieldSeqParser()
	if !p.RunSubParser(toks, sigParser) {
		return nil
	}

	return expressions.NewStruct(sigParser.GetFields()...)
}

// func (p *ExprParser) parseAccessElem(toks parsing.TokenSeq, ary *expressions.Expression) *expressions.Expression {
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

func (p *ExprParser) parseAnd(toks parsing.TokenSeq, left *expressions.Expression) *expressions.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewAnd)
}

func (p *ExprParser) parseOr(toks parsing.TokenSeq, left *expressions.Expression) *expressions.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewOr)
}

func (p *ExprParser) parseAdd(toks parsing.TokenSeq, left *expressions.Expression) *expressions.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewAdd)
}

func (p *ExprParser) parseSubtract(toks parsing.TokenSeq, left *expressions.Expression) *expressions.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewSubtract)
}

func (p *ExprParser) parseMultiply(toks parsing.TokenSeq, left *expressions.Expression) *expressions.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewMultiply)
}

func (p *ExprParser) parseDivide(toks parsing.TokenSeq, left *expressions.Expression) *expressions.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewDivide)
}

func (p *ExprParser) parseEqual(toks parsing.TokenSeq, left *expressions.Expression) *expressions.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewEqual)
}

func (p *ExprParser) parseNotEqual(toks parsing.TokenSeq, left *expressions.Expression) *expressions.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewNotEqual)
}

func (p *ExprParser) parseLess(toks parsing.TokenSeq, left *expressions.Expression) *expressions.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewLess)
}

func (p *ExprParser) parseLessEqual(toks parsing.TokenSeq, left *expressions.Expression) *expressions.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewLessEqual)
}

func (p *ExprParser) parseGreater(toks parsing.TokenSeq, left *expressions.Expression) *expressions.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewGreater)
}

func (p *ExprParser) parseGreaterEqual(toks parsing.TokenSeq, left *expressions.Expression) *expressions.Expression {
	return p.parseInfixExpr(toks, left, expressions.NewGreaterEqual)
}

type newInfixExprFn func(left, right *expressions.Expression) *expressions.Expression

func (p *ExprParser) parseInfixExpr(toks parsing.TokenSeq, left *expressions.Expression, fn newInfixExprFn) *expressions.Expression {
	prec := parsing.TokenPrecedence(toks.Current().Type())

	toks.Advance()

	right := p.parseExpression(toks, prec)

	return fn(left, right)
}
