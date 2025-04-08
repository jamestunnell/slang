package ast

import "github.com/jamestunnell/slang/ast/statements"

type Import struct {
	Rename    string
	PathParts []string
}

func NewImport(s *statements.Statement) *Import {
	core, ok := s.Core.(*statements.Use)
	if !ok {
		return &Import{
			Rename:    "",
			PathParts: []string{},
		}
	}

	return &Import{
		Rename:    core.Rename,
		PathParts: core.PathParts,
	}
}

func (i *Import) GetRename() string {
	return i.Rename
}

func (i *Import) GetPathParts() []string {
	return i.PathParts
}
