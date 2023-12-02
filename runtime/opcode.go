package runtime

type Opcode byte

const (
	OpCONST Opcode = iota
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
)
