package parsing

import "github.com/jamestunnell/slang"

type Precedence int

const (
	PrecedenceLOWEST      Precedence = iota
	PrecedenceEQUALS                 // ==
	PrecedenceLESSGREATER            // >, <, >=, <=
	PrecedenceSUM                    // +
	PrecedencePRODUCT                // *
	PrecedencePREFIX                 // -X or !X
	PrecedenceDOTCALLIDX             // a.b.c, myFunction(X), myArray[0]
)

var precedences = map[slang.TokenType]Precedence{
	slang.TokenEQUAL:        PrecedenceEQUALS,
	slang.TokenNOTEQUAL:     PrecedenceEQUALS,
	slang.TokenLESS:         PrecedenceLESSGREATER,
	slang.TokenLESSEQUAL:    PrecedenceLESSGREATER,
	slang.TokenGREATER:      PrecedenceLESSGREATER,
	slang.TokenGREATEREQUAL: PrecedenceLESSGREATER,
	slang.TokenPLUS:         PrecedenceSUM,
	slang.TokenMINUS:        PrecedenceSUM,
	slang.TokenSLASH:        PrecedencePRODUCT,
	slang.TokenSTAR:         PrecedencePRODUCT,
	slang.TokenDOT:          PrecedenceDOTCALLIDX,
	slang.TokenLPAREN:       PrecedenceDOTCALLIDX,
	slang.TokenLBRACKET:     PrecedenceDOTCALLIDX,
}

func TokenPrecedence(tokType slang.TokenType) Precedence {
	if p, ok := precedences[tokType]; ok {
		return p
	}

	return PrecedenceLOWEST
}
