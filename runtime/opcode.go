package runtime

type Opcode byte

const (
	OpPOP Opcode = iota

	OpGETCONST
	OpGETGLOBAL
	OpGETLOCAL
	OpGETFREE

	OpSETGLOBAL
	OpSETLOCAL

	OpCLOSURE
	OpCURRENTCLOSURE

	OpJUMP
	OpJUMPIFFALSE

	OpCALL
	OpRETURN
	OpRETURNVAL

	OpADD
	OpSUB
	OpMUL
	OpDIV
	OpNEG

	OpEQ
	OpNEQ
	OpLT
	OpLEQ
	OpGT
	OpGEQ

	OpNOT
	OpAND
	OpOR
)

func (o Opcode) String() string {
	var str string
	switch o {
	case OpPOP:
		str = "POP"
	case OpGETCONST:
		str = "GETCONST"
	case OpGETGLOBAL:
		str = "GETGLOBAL"
	case OpGETLOCAL:
		str = "GETLOCAL"
	case OpGETFREE:
		str = "GETFREE"
	case OpSETGLOBAL:
		str = "SETGLOBAL"
	case OpSETLOCAL:
		str = "SETLOCAL"
	case OpCLOSURE:
		str = "CLOSURE"
	case OpCURRENTCLOSURE:
		str = "CURRENTCLOSURE"
	case OpJUMP:
		str = "JUMP"
	case OpJUMPIFFALSE:
		str = "JUMPIFFALSE"
	case OpCALL:
		str = "CALL"
	case OpRETURN:
		str = "RETURN"
	case OpRETURNVAL:
		str = "RETURNVAL"
	case OpADD:
		str = "ADD"
	case OpSUB:
		str = "SUB"
	case OpMUL:
		str = "MUL"
	case OpDIV:
		str = "DIV"
	case OpNEG:
		str = "NEG"
	case OpEQ:
		str = "EQ"
	case OpNEQ:
		str = "NEQ"
	case OpLT:
		str = "LT"
	case OpLEQ:
		str = "LEQ"
	case OpGT:
		str = "GET"
	case OpGEQ:
		str = "GEQ"
	case OpNOT:
		str = "NOT"
	case OpAND:
		str = "AND"
	case OpOR:
		str = "OR"
	}

	return str
}
