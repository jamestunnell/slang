package field

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/types"
)

type Field struct {
	Name string      `json:"name"`
	Type *types.Type `json:"type"`
}

func New(name string, typ *types.Type) *Field {
	return &Field{
		Name: name,
		Type: typ,
	}
}

func (p *Field) GetName() string {
	return p.Name
}

func (p *Field) GetType() slang.Type {
	return p.Type
}

func (p *Field) IsEqual(other *Field) bool {
	return p.Name == other.Name && p.Type.IsEqual(other.Type)
}
