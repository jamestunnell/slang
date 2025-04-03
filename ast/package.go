package ast

import "github.com/jamestunnell/slang"

type Package struct {
	*slang.PackageInfo

	Modules []*Module `json:"modules"`
}

func NewPackage(info *slang.PackageInfo, modules ...*Module) *Package {
	return &Package{
		PackageInfo: info,
		Modules:     modules,
	}
}
