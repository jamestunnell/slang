package parser

import (
	"github.com/jamestunnell/slang"
)

type Precedence int

const (
	PrecedenceLOWEST      Precedence = iota
	PrecedenceEQUALS                 // ==
	PrecedenceLESSGREATER            // >, <, >=, <=
	PrecedenceSUM                    // +
	PrecedencePRODUCT                // *
	PrecedencePREFIX                 // -X or !X
	PrecedenceCALL                   // myFunction(X)
	PrecedenceINDEX                  // myArray[0]
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
	slang.TokenLPAREN:       PrecedenceCALL,
	slang.TokenLBRACKET:     PrecedenceINDEX,
}
