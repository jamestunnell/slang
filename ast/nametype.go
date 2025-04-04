package ast

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/types"
)

type NameType struct {
	Name string      `json:"name"`
	Type *types.Type `json:"type"`
}

func NewNameType(name string, typ *types.Type) *NameType {
	return &NameType{
		Name: name,
		Type: typ,
	}
}

func (p *NameType) GetName() string {
	return p.Name
}

func (p *NameType) GetType() slang.Type {
	return p.Type
}
