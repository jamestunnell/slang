package ast

import "github.com/jamestunnell/slang"

type Package struct {
	*slang.PackageInfo

	Modules []*Module `json:"modules"`
}

func NewPackage(name string, modules ...*Module) *Package {
	return &Package{
		PackageInfo: &slang.PackageInfo{Name: name},
		Modules:     modules,
	}
}

func (pkg *Package) GetModules() []slang.Module {
	modules := make([]slang.Module, len(pkg.Modules))

	for i, m := range pkg.Modules {
		modules[i] = m
	}

	return modules
}
