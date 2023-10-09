package parser

import (
	"errors"

	"github.com/jamestunnell/slang"
)

type prefixParseFn func() slang.Expression
type infixParseFn func(slang.Expression) slang.Expression

type Parser struct {
	Statements []slang.Statement
	Errors     []*ParseErr

	lexer slang.Lexer

	curToken  *slang.Token
	peekToken *slang.Token

	prefixParseFns map[slang.TokenType]prefixParseFn
	infixParseFns  map[slang.TokenType]infixParseFn
}

type ParseResults struct {
	Statements []slang.Statement
	Errors     []*ParseErr
}

var (
	errEmptyFuncBody = errors.New("function body is empty")
	errMissingReturn = errors.New("function missing return")
)

func New(l slang.Lexer) *Parser {
	p := &Parser{
		Statements:     []slang.Statement{},
		Errors:         []*ParseErr{},
		lexer:          l,
		prefixParseFns: map[slang.TokenType]prefixParseFn{},
		infixParseFns:  map[slang.TokenType]infixParseFn{},
	}

	p.registerPrefix(slang.TokenIDENT, p.parseIdentifier)
	p.registerPrefix(slang.TokenINT, p.parseInteger)
	p.registerPrefix(slang.TokenFLOAT, p.parseFloat)
	p.registerPrefix(slang.TokenSTRING, p.parseString)
	p.registerPrefix(slang.TokenTRUE, p.parseTrue)
	p.registerPrefix(slang.TokenFALSE, p.parseFalse)
	p.registerPrefix(slang.TokenMINUS, p.parseNegative)
	p.registerPrefix(slang.TokenBANG, p.parseNot)
	p.registerPrefix(slang.TokenLPAREN, p.parseGroupedExpression)
	p.registerPrefix(slang.TokenLBRACKET, p.parseArray)
	p.registerPrefix(slang.TokenIF, p.parseIfExpression)
	p.registerPrefix(slang.TokenFUNC, p.parseFuncLiteral)

	p.registerInfix(slang.TokenPLUS, p.parseAdd)
	p.registerInfix(slang.TokenMINUS, p.parseSubtract)
	p.registerInfix(slang.TokenSTAR, p.parseMultiply)
	p.registerInfix(slang.TokenSLASH, p.parseDivide)
	p.registerInfix(slang.TokenEQUAL, p.parseEqual)
	p.registerInfix(slang.TokenNOTEQUAL, p.parseNotEqual)
	p.registerInfix(slang.TokenLESS, p.parseLess)
	p.registerInfix(slang.TokenLESSEQUAL, p.parseLessEqual)
	p.registerInfix(slang.TokenGREATER, p.parseGreater)
	p.registerInfix(slang.TokenGREATEREQUAL, p.parseGreaterEqual)
	p.registerInfix(slang.TokenDOT, p.parseMethodCall)
	p.registerInfix(slang.TokenLPAREN, p.parseFunctionCall)
	p.registerInfix(slang.TokenLBRACKET, p.parseIndex)

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()

	p.nextToken()

	return p
}
func (p *Parser) registerPrefix(tokenType slang.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType slang.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

// func (p *Parser) nextTokenSkipLines() {
// 	p.nextToken()

// 	for p.curTokenIs(slang.TokenLINE) {
// 		p.nextToken()
// 	}
// }

// func (p *Parser) skipToNextLineOrEOF() {
// 	for !p.curTokenIs(slang.TokenEOF) && !p.curTokenIs(slang.TokenLINE) {
// 		p.nextToken()
// 	}

// 	if p.curTokenIs(slang.TokenLINE) {
// 		p.nextToken()
// 	}
// }

func (p *Parser) Run() *ParseResults {
	statments := p.parseStatementsUntil(slang.TokenEOF)

	p.Statements = append(p.Statements, statments...)

	return &ParseResults{
		Statements: p.Statements,
		Errors:     p.Errors,
	}
}

func (p *Parser) parseStatementsUntil(
	stopTokType slang.TokenType) []slang.Statement {
	// skip past comments, empty lines and semicolons at the beginning
	for p.curTokenIs(slang.TokenLINE, slang.TokenSEMICOLON, slang.TokenCOMMENT) {
		p.nextToken()
	}

	statements := []slang.Statement{}

	// read each statement until stop token
	for !p.curTokenIs(stopTokType, slang.TokenEOF) {
		if st := p.parseStatement(); st != nil {
			statements = append(statements, st)
		}

		// after a statement we expect semicolon, comment, newline, or EOF
		if !p.expectPeek(slang.TokenSEMICOLON, slang.TokenCOMMENT, slang.TokenLINE, slang.TokenEOF) {
			break
		}

		if p.peekTokenIs(stopTokType) {
			p.nextToken()

			break
		}

		// skip past lines, semicolons, and comments
		for p.peekTokenIs(slang.TokenLINE, slang.TokenSEMICOLON, slang.TokenCOMMENT) {
			p.nextToken()
		}

		p.nextToken()
	}

	// did we stop because of EOF or the expected stop token?
	if !p.curTokenIs(stopTokType) {
		err := NewParseError(
			NewErrWrongTokenType(stopTokType), p.curToken)

		p.Errors = append(p.Errors, err)
	}

	return statements
}

func (p *Parser) curTokenIs(expectedTypes ...slang.TokenType) bool {
	for _, t := range expectedTypes {
		if p.curToken.Info.Type() == t {
			return true
		}
	}

	return false
}

func (p *Parser) peekTokenIs(expectedTypes ...slang.TokenType) bool {
	for _, t := range expectedTypes {
		if p.peekToken.Info.Type() == t {
			return true
		}
	}

	return false
}

func (p *Parser) expectPeek(expectedTypes ...slang.TokenType) bool {
	if !p.peekTokenIs(expectedTypes...) {
		p.peekError(expectedTypes...)

		return false
	}

	return true
}

func (p *Parser) expectPeekAndAdvance(expectedTypes ...slang.TokenType) bool {
	if !p.peekTokenIs(expectedTypes...) {
		p.peekError(expectedTypes...)

		return false
	}

	p.nextToken()

	return true
}

func (p *Parser) peekError(expectedTypes ...slang.TokenType) {
	err := NewErrWrongTokenType(expectedTypes...)
	pErr := NewParseError(err, p.peekToken)

	p.Errors = append(p.Errors, pErr)
}

func (p *Parser) NewParseErr(err error) *ParseErr {
	return NewParseError(err, p.curToken)
}

func (p *Parser) peekPrecedence() Precedence {
	if p, ok := precedences[p.peekToken.Info.Type()]; ok {
		return p
	}

	return PrecedenceLOWEST
}

func (p *Parser) curPrecedence() Precedence {
	if p, ok := precedences[p.curToken.Info.Type()]; ok {
		return p
	}

	return PrecedenceLOWEST
}
