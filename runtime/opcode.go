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
