package compiler

type SymbolScope string

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

const (
	GlobalScope = "GLOBAL"
)

type SymbolTable struct {
	store   map[string]Symbol
	numDefs int
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		store: make(map[string]Symbol),
	}
}

func (st *SymbolTable) Define(name string) Symbol {
	sym := Symbol{
		Name:  name,
		Index: st.numDefs,
		Scope: GlobalScope,
	}

	st.store[name] = sym

	st.numDefs++

	return sym
}

func (st *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := st.store[name]

	return obj, ok
}
