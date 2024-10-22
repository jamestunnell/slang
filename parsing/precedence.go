package parsing

import (
	"github.com/jamestunnell/slang/lexing"
)

type Precedence int

const (
	PrecedenceLOWEST     Precedence = iota
	PrecedenceOR                    // or
	PrecedenceAND                   // and
	PrecedenceEQUALITY              // ==, !=
	PrecedenceRELATIONAL            // <, <=, >, >=
	PrecedenceADDSUB                // +, -
	PrecedenceMULDIVREM             // *, /, %
	PrecedencePREFIX                // -X or !X
	PrecedenceDOTCALLIDX            // a.b.c, myFunction(X), myArray[0]
)

var precedences = map[lexing.TokenType]Precedence{
	lexing.TokenOR:           PrecedenceOR,
	lexing.TokenAND:          PrecedenceAND,
	lexing.TokenEQUAL:        PrecedenceEQUALITY,
	lexing.TokenNOTEQUAL:     PrecedenceEQUALITY,
	lexing.TokenLESS:         PrecedenceRELATIONAL,
	lexing.TokenLESSEQUAL:    PrecedenceRELATIONAL,
	lexing.TokenGREATER:      PrecedenceRELATIONAL,
	lexing.TokenGREATEREQUAL: PrecedenceRELATIONAL,
	lexing.TokenPLUS:         PrecedenceADDSUB,
	lexing.TokenMINUS:        PrecedenceADDSUB,
	lexing.TokenSLASH:        PrecedenceMULDIVREM,
	lexing.TokenSTAR:         PrecedenceMULDIVREM,
	lexing.TokenDOT:          PrecedenceDOTCALLIDX,
	lexing.TokenLPAREN:       PrecedenceDOTCALLIDX,
	lexing.TokenLBRACKET:     PrecedenceDOTCALLIDX,
}

func TokenPrecedence(tokType lexing.TokenType) Precedence {
	if p, ok := precedences[tokType]; ok {
		return p
	}

	return PrecedenceLOWEST
}
