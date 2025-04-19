package ast

import "github.com/jamestunnell/slang"

type Package struct {
	Meta    slang.PackageMeta
	Modules []*Module `json:"modules"`
}

func NewPackage(meta slang.PackageMeta, modules ...*Module) *Package {
	return &Package{
		Meta:    meta,
		Modules: modules,
	}
}

func (pkg *Package) GetMeta() slang.PackageMeta {
	return pkg.Meta
}

func (pkg *Package) GetModules() []slang.Module {
	modules := make([]slang.Module, len(pkg.Modules))

	for i, m := range pkg.Modules {
		modules[i] = m
	}

	return modules
}
