package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/parsing"
)

type ExprParser struct {
	*ParserBase

	prec parsing.Precedence
	Expr *expressions.Expression
}

type prefixParseFn func(slang.TokenSeq) *expressions.Expression
type infixParseFn func(slang.TokenSeq, *expressions.Expression) *expressions.Expression

func NewExprParser(prec parsing.Precedence) *ExprParser {
	return &ExprParser{
		prec:       prec,
		ParserBase: NewParserBase(),
	}
}

func (p *ExprParser) Run(toks slang.TokenSeq) bool {
	p.Expr = p.parseExpression(toks, p.prec)

	return p.Expr != nil
}

func (p *ExprParser) findPrefixParseFn(
	tokType slang.TokenType) (prefixParseFn, bool) {
	var prefixParse prefixParseFn

	switch tokType {
	case slang.TokenSYMBOL:
		prefixParse = p.parseIdentifier
	case slang.TokenINTVAL:
		prefixParse = p.parseIntVal
	case slang.TokenFLTVAL:
		prefixParse = p.parseFloatVal
	case slang.TokenSTRVAL:
		prefixParse = p.parseStrVal
	case slang.TokenVERBATIMSTRING:
		prefixParse = p.parseVerbatimString
	case slang.TokenBOOLVAL:
		prefixParse = p.parseBoolVal
	case slang.TokenMINUS:
		prefixParse = p.parseNegative
	case slang.TokenBANG:
		prefixParse = p.parseNot
	case slang.TokenLPAREN:
		prefixParse = p.parseGroupedExpression
	// case slang.TokenLBRACE:
	// 	prefixParse = p.parseStructAnon
	// case slang.TokenLBRACKET:
	// 	prefixParse = p.parseArrayAuto
	// case slang.TokenLESS:
	// 	prefixParse = p.parseMapAuto
	case slang.TokenFUNC:
		prefixParse = p.parseFuncAnon
	case slang.TokenSTRUCT:
		prefixParse = p.parseStructAnon
	}

	return prefixParse, prefixParse != nil
}

func (p *ExprParser) findInfixParseFn(
	tokType slang.TokenType) (infixParseFn, bool) {
	var infixParse infixParseFn

	switch tokType {
	case slang.TokenAND:
		infixParse = p.parseAnd
	case slang.TokenOR:
		infixParse = p.parseOr
	case slang.TokenPLUS:
		infixParse = p.parseAdd
	case slang.TokenMINUS:
		infixParse = p.parseSubtract
	case slang.TokenSTAR:
		infixParse = p.parseMultiply
	case slang.TokenSLASH:
		infixParse = p.parseDivide
	case slang.TokenEQUALEQUAL:
		infixParse = p.parseEqual
	case slang.TokenNOTEQUAL:
		infixParse = p.parseNotEqual
	case slang.TokenLESS:
		infixParse = p.parseLess
	case slang.TokenLESSEQUAL:
		infixParse = p.parseLessEqual
	case slang.TokenGREATER:
		infixParse = p.parseGreater
	case slang.TokenGREATEREQUAL:
		infixParse = p.parseGreaterEqual
	case slang.TokenDOT:
		infixParse = p.parseAccessMember
	case slang.TokenLPAREN:
		infixParse = p.parseInvoke
	}

	return infixParse, infixParse != nil
}
