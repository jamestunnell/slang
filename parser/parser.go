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

	p.registerPrefix(slang.TokenSYMBOL, p.parseIdentifier)
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

	// Read two tokens to get current and next
	p.nextTokenSkipComments()
	p.nextTokenSkipComments()

	return p
}
func (p *Parser) registerPrefix(tokenType slang.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType slang.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) nextTokenSkipComments() {
	tok := p.lexer.NextToken()
	for tok.Info.Type() == slang.TokenCOMMENT {
		tok = p.lexer.NextToken()
	}

	p.curToken = p.peekToken
	p.peekToken = tok
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
	statments := p.parseFileStatements()

	p.Statements = append(p.Statements, statments...)

	return &ParseResults{
		Statements: p.Statements,
		Errors:     p.Errors,
	}
}

func (p *Parser) parseFileStatements() []slang.Statement {
	// skip past empty lines at the beginning
	for p.curTokenIs(slang.TokenLINE) {
		p.nextTokenSkipComments()
	}

	statements := []slang.Statement{}

	for !p.curTokenIs(slang.TokenEOF) {
		if st := p.parseFileStatement(); st != nil {
			statements = append(statements, st)
		}

		if !p.expectPeek(slang.TokenLINE, slang.TokenEOF) {
			p.nextTokenSkipComments()
		}

		p.nextTokenSkipComments()

		// remove extra lines
		for p.curTokenIs(slang.TokenLINE) {
			p.nextTokenSkipComments()
		}
	}

	return statements
}

func (p *Parser) parseBlockStatements() []slang.Statement {
	// skip past empty lines at the beginning
	for p.curTokenIs(slang.TokenLINE) {
		p.nextTokenSkipComments()
	}

	statements := []slang.Statement{}

	// read each statement until RBRACE or EOF
	for !p.curTokenIs(slang.TokenEOF, slang.TokenRBRACE) {
		if st := p.parseBlockStatement(); st != nil {
			statements = append(statements, st)
		}

		if !p.expectPeek(slang.TokenLINE, slang.TokenRBRACE) {
			p.nextTokenSkipComments()
		}

		p.nextTokenSkipComments()

		// remove extra lines
		for p.curTokenIs(slang.TokenLINE) {
			p.nextTokenSkipComments()
		}
	}

	// if we quit loop because of EOF then the block wasn't terminated correctly
	if !p.curTokenIs(slang.TokenRBRACE) {
		err := NewParseError(
			NewErrWrongTokenType(slang.TokenRBRACE), p.curToken)

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

func (p *Parser) expectCur(expectedTypes ...slang.TokenType) bool {
	if !p.curTokenIs(expectedTypes...) {
		p.tokenErr(p.curToken, expectedTypes...)

		return false
	}

	return true
}

func (p *Parser) expectPeek(expectedTypes ...slang.TokenType) bool {
	if !p.peekTokenIs(expectedTypes...) {
		p.tokenErr(p.peekToken, expectedTypes...)

		return false
	}

	return true
}

func (p *Parser) expectPeekAndAdvance(expectedTypes ...slang.TokenType) bool {
	if !p.peekTokenIs(expectedTypes...) {
		p.tokenErr(p.peekToken, expectedTypes...)

		return false
	}

	p.nextTokenSkipComments()

	return true
}

func (p *Parser) tokenErr(tok *slang.Token, expectedTypes ...slang.TokenType) {
	err := NewErrWrongTokenType(expectedTypes...)
	pErr := NewParseError(err, tok)

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
