package runtime

type Opcode byte

const (
	OpCONST Opcode = iota
	OpPOP

	OpGETGLOBAL
	OpSETGLOBAL

	OpJUMP
	OpJUMPIFFALSE

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
