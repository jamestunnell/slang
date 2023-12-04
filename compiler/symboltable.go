package compiler

type SymbolScope string

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

const (
	GlobalScope = "GLOBAL"
	LocalScope  = "LOCAL"
	FreeScope   = "FREE"
)

type SymbolTable struct {
	store       map[string]Symbol
	numDefs     int
	outer       *SymbolTable
	freeSymbols []Symbol
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		store: make(map[string]Symbol),
		outer: nil,
	}
}

func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	return &SymbolTable{
		store: make(map[string]Symbol),
		outer: outer,
	}
}

func (st *SymbolTable) NumDefs() int {
	return st.numDefs
}

func (st *SymbolTable) FreeSymbols() []Symbol {
	return st.freeSymbols
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

func (st *SymbolTable) defineFree(original Symbol) Symbol {
	st.freeSymbols = append(st.freeSymbols, original)

	symbol := Symbol{Name: original.Name, Index: len(st.freeSymbols) - 1}
	symbol.Scope = FreeScope

	st.store[original.Name] = symbol

	return symbol
}

func (st *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := st.store[name]
	if !ok && st.outer != nil {
		obj, ok = st.outer.Resolve(name)
		if !ok {
			return obj, false
		}

		if obj.Scope == GlobalScope {
			return obj, true
		}

		free := st.defineFree(obj)

		return free, true
	}

	return obj, ok
}
