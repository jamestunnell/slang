package compilation

type Symbol struct {
	Scope string
	Name  string
	Type  SymbolType
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

func NewSymbol(scope, name string, typ SymbolType) Symbol {
	return Symbol{
		Name:  name,
		Scope: scope,
		Type:  typ,
	}
}
