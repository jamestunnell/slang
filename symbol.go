package slang

import (
	"strings"

	"golang.org/x/exp/slices"
)

type Symbol struct {
	Parts []string
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

func NewRootSymbol(name string, typ SymbolType) *Symbol {
	return &Symbol{
		Parts: []string{name},
		Type:  typ,
	}
}

func NewChildSymbol(name string, typ SymbolType, parent *Symbol) *Symbol {
	return &Symbol{
		Parts: append(slices.Clone(parent.Parts), name),
		Type:  typ,
	}
}

func (sym *Symbol) Scope() string {
	n := len(sym.Parts)

	if n < 2 {
		return ""
	}

	return strings.Join(sym.Parts[:n-1], ".")
}

func (sym *Symbol) Name() string {
	return sym.Parts[len(sym.Parts)-1]
}

func (sym *Symbol) ScopedName() string {
	return strings.Join(sym.Parts, ".")
}
