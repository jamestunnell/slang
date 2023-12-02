package runtime

type Opcode byte

const (
	OpCONST Opcode = iota
	OpPOP

	OpADD
	OpSUB
	OpMUL
	OpDIV

	OpEQ
	OpNEQ
	OpLT
	OpLEQ
	OpGT
	OpGEQ

	OpAND
	OpOR
)
