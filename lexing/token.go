package lexing

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type TokenType int

type TokenInfo struct {
	Type  TokenType
	Value string
}

type SourceLocation struct {
	Line, Column int
}

type Token struct {
	*TokenInfo

	Location SourceLocation
}

func NewTokenInfo(typ TokenType, val string) *TokenInfo {
	return &TokenInfo{
		Type:  typ,
		Value: val,
	}
}

func NewToken(info *TokenInfo, loc SourceLocation) *Token {
	return &Token{
		TokenInfo: info,
		Location:  loc,
	}
}

func NewLoc(line, col int) SourceLocation {
	return SourceLocation{Line: line, Column: col}
}

func (tok *Token) Is(tokTypes ...TokenType) bool {
	return slices.Contains(tokTypes, tok.Type)
}

func (loc SourceLocation) String() string {
	return fmt.Sprintf("(line: %d, col: %d)", loc.Line, loc.Column)
}

const (
	TokenAND TokenType = iota
	TokenBANG
	TokenBREAK
	TokenCLASS
	TokenCOLON
	TokenCOMMA
	TokenCOMMENT
	TokenCONST
	TokenCONTINUE
	TokenDOLLARLBRACE
	TokenDOT
	TokenELSE
	TokenEOF
	TokenEQUAL
	TokenEQUALEQUAL
	TokenFALSE
	TokenFIELD
	TokenFLOAT
	TokenFOREACH
	TokenFUNC
	TokenGREATER
	TokenGREATEREQUAL
	TokenIF
	TokenILLEGAL
	TokenIN
	TokenINT
	TokenLBRACE
	TokenLBRACKET
	TokenLESS
	TokenLESSEQUAL
	TokenLPAREN
	TokenMETHOD
	TokenMINUS
	TokenMODULE
	TokenNEWLINE
	TokenOR
	TokenNOTEQUAL
	TokenPLUS
	TokenRBRACE
	TokenRBRACKET
	TokenRETURN
	TokenRPAREN
	TokenSEMICOLON
	TokenSLASH
	TokenSTAR
	TokenSTRING
	TokenSYMBOL
	TokenTRUE
	TokenVAR
	TokenVERBATIMSTRING
	TokenUSE
)

func (tt TokenType) String() string {
	var str string

	switch tt {
	case TokenAND:
		str = "AND"
	case TokenBANG:
		str = "BANG"
	case TokenBREAK:
		str = "BREAK"
	case TokenCLASS:
		str = "CLASS"
	case TokenCOLON:
		str = "COLON"
	case TokenCOMMA:
		str = "COMMA"
	case TokenCOMMENT:
		str = "COMMENT"
	case TokenCONST:
		str = "CONST"
	case TokenCONTINUE:
		str = "CONTINUE"
	case TokenDOT:
		str = "DOT"
	case TokenELSE:
		str = "ELSE"
	case TokenEOF:
		str = "EOF"
	case TokenEQUAL:
		str = "EQUAL"
	case TokenEQUALEQUAL:
		str = "EQUALEQUAL"
	case TokenFALSE:
		str = "FALSE"
	case TokenFIELD:
		str = "FIELD"
	case TokenFLOAT:
		str = "FLOAT"
	case TokenFOREACH:
		str = "FOREACH"
	case TokenFUNC:
		str = "FUNC"
	case TokenGREATER:
		str = "GREATER"
	case TokenGREATEREQUAL:
		str = "GREATEREQUAL"
	case TokenIF:
		str = "IF"
	case TokenILLEGAL:
		str = "ILLEGAL"
	case TokenIN:
		str = "IN"
	case TokenINT:
		str = "INT"
	case TokenLBRACE:
		str = "LBRACE"
	case TokenLBRACKET:
		str = "LBRACKET"
	case TokenLESS:
		str = "LESS"
	case TokenLESSEQUAL:
		str = "LESSEQUAL"
	case TokenLPAREN:
		str = "LPAREN"
	case TokenMODULE:
		str = "MODULE"
	case TokenMETHOD:
		str = "METHOD"
	case TokenMINUS:
		str = "MINUS"
	case TokenNEWLINE:
		str = "NEWLINE"
	case TokenNOTEQUAL:
		str = "NOTEQUAL"
	case TokenOR:
		str = "OR"
	case TokenPLUS:
		str = "PLUS"
	case TokenRBRACE:
		str = "RBRACE"
	case TokenRBRACKET:
		str = "RBRACKET"
	case TokenRETURN:
		str = "RETURN"
	case TokenRPAREN:
		str = "RPAREN"
	case TokenSEMICOLON:
		str = "SEMICOLON"
	case TokenSLASH:
		str = "SLASH"
	case TokenSTAR:
		str = "STAR"
	case TokenSTRING:
		str = "STRING"
	case TokenSYMBOL:
		str = "SYMBOL"
	case TokenTRUE:
		str = "TRUE"
	case TokenVERBATIMSTRING:
		str = "VERBATIMSTRING"
	case TokenUSE:
		str = "USE"
	case TokenVAR:
		str = "VAR"
	}

	return str
}
