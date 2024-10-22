package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/lexing"
)

type ExprParser struct {
	*ParserBase

	prec Precedence
	Expr slang.Expression
}

type prefixParseFn func(lexing.TokenSeq) slang.Expression
type infixParseFn func(lexing.TokenSeq, slang.Expression) slang.Expression

func NewExprParser(prec Precedence) *ExprParser {
	return &ExprParser{
		prec:       prec,
		ParserBase: NewParserBase(),
	}
}

func (p *ExprParser) Run(toks lexing.TokenSeq) bool {
	p.Expr = p.parseExpression(toks, p.prec)

	return p.Expr != nil
}

func (p *ExprParser) findPrefixParseFn(
	tokType lexing.TokenType) (prefixParseFn, bool) {
	var prefixParse prefixParseFn

	switch tokType {
	case lexing.TokenSYMBOL:
		prefixParse = p.parseIdentifier
	case lexing.TokenINT:
		prefixParse = p.parseInteger
	case lexing.TokenFLOAT:
		prefixParse = p.parseFloat
	case lexing.TokenSTRING:
		prefixParse = p.parseString
	case lexing.TokenVERBATIMSTRING:
		prefixParse = p.parseVerbatimString
	case lexing.TokenTRUE:
		prefixParse = p.parseTrue
	case lexing.TokenFALSE:
		prefixParse = p.parseFalse
	case lexing.TokenMINUS:
		prefixParse = p.parseNegative
	case lexing.TokenBANG:
		prefixParse = p.parseNot
	case lexing.TokenLPAREN:
		prefixParse = p.parseGroupedExpression
	case lexing.TokenLBRACKET:
		prefixParse = p.parseArrayOrMap
	case lexing.TokenFUNC:
		prefixParse = p.parseFuncLiteral
	}

	return prefixParse, prefixParse != nil
}

func (p *ExprParser) findInfixParseFn(
	tokType lexing.TokenType) (infixParseFn, bool) {
	var infixParse infixParseFn

	switch tokType {
	case lexing.TokenAND:
		infixParse = p.parseAnd
	case lexing.TokenOR:
		infixParse = p.parseOr
	case lexing.TokenPLUS:
		infixParse = p.parseAdd
	case lexing.TokenMINUS:
		infixParse = p.parseSubtract
	case lexing.TokenSTAR:
		infixParse = p.parseMultiply
	case lexing.TokenSLASH:
		infixParse = p.parseDivide
	case lexing.TokenEQUAL:
		infixParse = p.parseEqual
	case lexing.TokenNOTEQUAL:
		infixParse = p.parseNotEqual
	case lexing.TokenLESS:
		infixParse = p.parseLess
	case lexing.TokenLESSEQUAL:
		infixParse = p.parseLessEqual
	case lexing.TokenGREATER:
		infixParse = p.parseGreater
	case lexing.TokenGREATEREQUAL:
		infixParse = p.parseGreaterEqual
	case lexing.TokenDOT:
		infixParse = p.parseAccessMember
	case lexing.TokenLPAREN:
		infixParse = p.parseCall
	case lexing.TokenLBRACKET:
		infixParse = p.parseAccessElem
	}

	return infixParse, infixParse != nil
}
