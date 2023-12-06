package compilation

type Symbol struct {
	Name string
	Type SymbolType
}

type SymbolType int

const (
	SymbolCLASS SymbolType = iota
	SymbolFUNC
	SymbolFIELD
	SymbolMETHOD
	SymbolMODULE
	SymbolVAR
)

func NewSymbol(name string, typ SymbolType) Symbol {
	return Symbol{
		Name: name,
		Type: typ,
	}
}
