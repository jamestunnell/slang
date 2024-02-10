package parsing

import (
	"github.com/jamestunnell/slang"
)

type ExprParser struct {
	*ParserBase

	prec Precedence
	Expr slang.Expression
}

type prefixParseFn func(slang.TokenSeq) slang.Expression
type infixParseFn func(slang.TokenSeq, slang.Expression) slang.Expression

func NewExprParser(prec Precedence) *ExprParser {
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
	case slang.TokenINT:
		prefixParse = p.parseInteger
	case slang.TokenFLOAT:
		prefixParse = p.parseFloat
	case slang.TokenSTRING:
		prefixParse = p.parseString
	case slang.TokenVERBATIMSTRING:
		prefixParse = p.parseVerbatimString
	case slang.TokenTRUE:
		prefixParse = p.parseTrue
	case slang.TokenFALSE:
		prefixParse = p.parseFalse
	case slang.TokenMINUS:
		prefixParse = p.parseNegative
	case slang.TokenBANG:
		prefixParse = p.parseNot
	case slang.TokenLPAREN:
		prefixParse = p.parseGroupedExpression
	case slang.TokenLBRACKET:
		prefixParse = p.parseArrayOrMap
	case slang.TokenFUNC:
		prefixParse = p.parseFuncLiteral
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
	case slang.TokenEQUAL:
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
		infixParse = p.parseCall
	case slang.TokenLBRACKET:
		infixParse = p.parseAccessElem
	}

	return infixParse, infixParse != nil
}
