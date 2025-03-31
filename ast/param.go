package ast

import "github.com/jamestunnell/slang"

type Param struct {
	Name string     `json:"name"`
	Type slang.Type `json:"paramType"`
}

func NewParam(name string, typ slang.Type) *Param {
	return &Param{
		Name: name,
		Type: typ,
	}
}

func (p *Param) GetName() string {
	return p.Name
}

func (p *Param) GetType() slang.Type {
	return p.Type
}

func (p *Param) Equal(other *Param) bool {
	return p.Name == other.Name && p.Type == other.Type
}
