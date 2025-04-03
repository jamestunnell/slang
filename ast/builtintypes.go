package ast

import "github.com/jamestunnell/slang"

type BoolType struct {
}

type ErrType struct {
}

type FltType struct {
}

type IntType struct {
}

type StrType struct {
}

func (t *BoolType) String() string {
	return slang.TypeBOOL
}

func (t *ErrType) String() string {
	return slang.TypeERR
}

func (t *FltType) String() string {
	return slang.TypeFLT
}

func (t *IntType) String() string {
	return slang.TypeINT
}

func (t *StrType) String() string {
	return slang.TypeSTR
}
