package parsing

import "github.com/jamestunnell/slang"

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

var precedences = map[slang.TokenType]Precedence{
	slang.TokenOR:           PrecedenceOR,
	slang.TokenAND:          PrecedenceAND,
	slang.TokenEQUAL:        PrecedenceEQUALITY,
	slang.TokenNOTEQUAL:     PrecedenceEQUALITY,
	slang.TokenLESS:         PrecedenceRELATIONAL,
	slang.TokenLESSEQUAL:    PrecedenceRELATIONAL,
	slang.TokenGREATER:      PrecedenceRELATIONAL,
	slang.TokenGREATEREQUAL: PrecedenceRELATIONAL,
	slang.TokenPLUS:         PrecedenceADDSUB,
	slang.TokenMINUS:        PrecedenceADDSUB,
	slang.TokenSLASH:        PrecedenceMULDIVREM,
	slang.TokenSTAR:         PrecedenceMULDIVREM,
	slang.TokenDOT:          PrecedenceDOTCALLIDX,
	slang.TokenLPAREN:       PrecedenceDOTCALLIDX,
	slang.TokenLBRACE:       PrecedenceDOTCALLIDX,
	slang.TokenLBRACKET:     PrecedenceDOTCALLIDX,
}

func TokenPrecedence(tokType slang.TokenType) Precedence {
	if p, ok := precedences[tokType]; ok {
		return p
	}

	return PrecedenceLOWEST
}
