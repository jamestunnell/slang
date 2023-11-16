package slang

import "fmt"

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

func (loc SourceLocation) String() string {
	return fmt.Sprintf("(line: %d, col: %d)", loc.Line, loc.Column)
}

const (
	TokenASSIGN TokenType = iota
	TokenBANG
	TokenCOLON
	TokenCOMMA
	TokenCOMMENT
	TokenDOT
	TokenIF
	TokenELSE
	TokenEOF
	TokenEQUAL
	TokenFALSE
	TokenFLOAT
	TokenFUNC
	TokenGREATER
	TokenGREATEREQUAL
	TokenILLEGAL
	TokenINT
	TokenLBRACE
	TokenLBRACKET
	TokenLESS
	TokenLESSEQUAL
	TokenLPAREN
	TokenMINUS
	TokenMINUSEQUAL
	TokenMINUSMINUS
	TokenNEWLINE
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
	TokenSTRING
	TokenSTRUCT
	TokenSYMBOL
	TokenTRUE
	TokenUSE
)

func (tt TokenType) String() string {
	var str string

	switch tt {
	case TokenASSIGN:
		str = "ASSIGN"
	case TokenBANG:
		str = "BANG"
	case TokenCOLON:
		str = "COLON"
	case TokenCOMMA:
		str = "COMMA"
	case TokenDOT:
		str = "DOT"
	case TokenIF:
		str = "IF"
	case TokenELSE:
		str = "ELSE"
	case TokenEOF:
		str = "EOF"
	case TokenEQUAL:
		str = "EQUAL"
	case TokenFALSE:
		str = "FALSE"
	case TokenFLOAT:
		str = "FLOAT"
	case TokenFUNC:
		str = "FUNC"
	case TokenGREATER:
		str = "GREATER"
	case TokenGREATEREQUAL:
		str = "GREATEREQUAL"
	case TokenILLEGAL:
		str = "ILLEGAL"
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
	case TokenSTRING:
		str = "STRING"
	case TokenSYMBOL:
		str = "SYMBOL"
	case TokenTRUE:
		str = "TRUE"
	}

	return str
}
