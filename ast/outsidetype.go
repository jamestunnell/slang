package ast

type OutsideType struct {
	ModuleName string `json:"moduleName"`
	StructName string `json:"structName"`
}

func NewOutsideType(moduleName, structName string) *OutsideType {
	return &OutsideType{
		ModuleName: moduleName,
		StructName: structName,
	}
}

func (t *OutsideType) String() string {
	return t.ModuleName + "." + t.StructName
}
