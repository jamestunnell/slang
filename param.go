package slang

type Param struct {
	Name, Type string
}

func NewParam(name, typ string) *Param {
	return &Param{Name: name, Type: typ}
}

func (p *Param) String() string {
	return p.Name + " " + p.Type
}
