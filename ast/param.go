package ast

type Param struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewParam(name, typ string) *Param {
	return &Param{Name: name, Type: typ}
}

func (p *Param) Equal(other *Param) bool {
	return p.Name == other.Name && p.Type == other.Type
}

func (p *Param) GetName() string {
	return p.Name
}

func (p *Param) GetType() string {
	return p.Name
}
