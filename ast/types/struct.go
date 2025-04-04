package types

import "github.com/jamestunnell/slang"

type Struct struct {
	ModuleName string `json:"module"`
	StructName string `json:"struct"`
}

func NewStruct(moduleName, structName string) *Type {
	core := &Struct{
		ModuleName: moduleName,
		StructName: structName,
	}

	return NewType(slang.TypeSTRUCT, core)
}

func (t *Struct) IsEqual(other Core) bool {
	t2, ok := other.(*Struct)
	if !ok {
		return false
	}

	if t.ModuleName != t2.ModuleName {
		return false
	}

	return t.StructName == t2.StructName
}

func (t *Struct) String() string {
	if t.ModuleName == "" {
		return t.StructName
	}

	return t.ModuleName + "." + t.StructName
}
