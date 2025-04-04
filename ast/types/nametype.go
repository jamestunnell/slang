package types

import (
	"github.com/jamestunnell/slang"
)

type NameType struct {
	Name string `json:"name"`
	Type *Type  `json:"type"`
}

func NewNameType(name string, typ *Type) *NameType {
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

func (p *NameType) IsEqual(other *NameType) bool {
	return p.Name == other.Name && p.Type.IsEqual(other.Type)
}
