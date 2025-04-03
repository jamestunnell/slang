package ast

type InsideType struct {
	StructName string `json:"structName"`
}

func NewInsideType(structName string) *InsideType {
	return &InsideType{
		StructName: structName,
	}
}

func (t *InsideType) String() string {
	return t.StructName
}
