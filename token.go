package slang

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type TokenType int

type TokenInfo interface {
	Type() TokenType
	Value() string
}

type SourceLocation struct {
	Line, Column int
}

type Token struct {
	Info     TokenInfo
	Location SourceLocation
}

func NewToken(info TokenInfo, loc SourceLocation) *Token {
	return &Token{
		Info:     info,
		Location: loc,
	}
}

func NewLoc(line, col int) SourceLocation {
	return SourceLocation{Line: line, Column: col}
}

func (tok *Token) Type() TokenType {
	return tok.Info.Type()
}

func (tok *Token) Value() string {
	return tok.Info.Value()
}

func (tok *Token) Is(tokTypes ...TokenType) bool {
	return slices.Contains(tokTypes, tok.Info.Type())
}

func (loc SourceLocation) String() string {
	return fmt.Sprintf("(line: %d, col: %d)", loc.Line, loc.Column)
}

const (
	TokenAND TokenType = iota
	TokenARY
	TokenEQUAL
	TokenBANG
	TokenBOOL
	TokenBOOLVAL
	TokenBREAK
	TokenCOLON
	TokenCOMMA
	TokenCOMMENT
	TokenCONST
	TokenCONTINUE
	TokenDOLLARLBRACE
	TokenDOT
	TokenELSE
	TokenERR
	TokenEOF
	TokenEQUALEQUAL
	TokenFLT
	TokenFLTVAL
	TokenFOREACH
	TokenFUNC
	TokenGREATER
	TokenGREATEREQUAL
	TokenIF
	TokenILLEGAL
	TokenUSE
	TokenIN
	TokenINT
	TokenINTERFACE
	TokenINTVAL
	TokenLBRACE
	TokenLBRACKET
	TokenLESS
	TokenLESSEQUAL
	TokenLESSLESS
	TokenLAMBDAOP
	TokenLPAREN
	TokenMAP
	TokenMINUS
	TokenMINUSEQUAL
	TokenMINUSMINUS
	TokenMODULE
	TokenNEWLINE
	TokenOR
	TokenNOTEQUAL
	TokenPLUS
	TokenPLUSEQUAL
	TokenPLUSPLUS
	TokenRBRACE
	TokenRBRACKET
	TokenRETURN
	TokenRPAREN
	TokenSEMICOLON
	TokenSLASH
	TokenSLASHEQUAL
	TokenSTAR
	TokenSTAREQUAL
	TokenSTR
	TokenSTRVAL
	TokenSTRUCT
	TokenSYMBOL
	TokenVAR
	TokenVERBATIMSTRING
)

func (tt TokenType) String() string {
	var str string

	switch tt {
	case TokenAND:
		str = "AND"
	case TokenARY:
		str = "ARY"
	case TokenEQUAL:
		str = "ASSIGN"
	case TokenBANG:
		str = "BANG"
	case TokenBOOL:
		str = "BOOl"
	case TokenBOOLVAL:
		str = "BOOLVAL"
	case TokenBREAK:
		str = "BREAK"
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
	case TokenERR:
		str = "ERR"
	case TokenEOF:
		str = "EOF"
	case TokenEQUALEQUAL:
		str = "EQUALEQUAL"
	case TokenFLT:
		str = "FLOAT"
	case TokenFLTVAL:
		str = "FLOATVAL"
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
	case TokenUSE:
		str = "USE"
	case TokenIN:
		str = "IN"
	case TokenINT:
		str = "INT"
	case TokenINTVAL:
		str = "INTVAL"
	case TokenLAMBDAOP:
		str = "LAMBDAOP"
	case TokenLBRACE:
		str = "LBRACE"
	case TokenLBRACKET:
		str = "LBRACKET"
	case TokenLESS:
		str = "LESS"
	case TokenLESSEQUAL:
		str = "LESSEQUAL"
	case TokenLESSLESS:
		str = "LESSLESS"
	case TokenLPAREN:
		str = "LPAREN"
	case TokenMAP:
		str = "MAP"
	case TokenMODULE:
		str = "MODULE"
	case TokenMINUS:
		str = "MINUS"
	case TokenMINUSEQUAL:
		str = "MINUSEQUAL"
	case TokenMINUSMINUS:
		str = "MINUSMINUS"
	case TokenNEWLINE:
		str = "NEWLINE"
	case TokenNOTEQUAL:
		str = "NOTEQUAL"
	case TokenOR:
		str = "OR"
	case TokenPLUS:
		str = "PLUS"
	case TokenPLUSEQUAL:
		str = "PLUSEQUAL"
	case TokenPLUSPLUS:
		str = "PLUSPLUS"
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
	case TokenSLASHEQUAL:
		str = "SLASHEQUAL"
	case TokenSTAR:
		str = "STAR"
	case TokenSTAREQUAL:
		str = "STAREQUAL"
	case TokenSTR:
		str = "STR"
	case TokenSTRVAL:
		str = "STRVAL"
	case TokenSTRUCT:
		str = "STRUCT"
	case TokenSYMBOL:
		str = "SYMBOL"
	case TokenVERBATIMSTRING:
		str = "VERBATIMSTRING"
	case TokenVAR:
		str = "VAR"
	}

	return str
}
