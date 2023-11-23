package ast

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/maps"
)

type Module struct {
	Comment   string                    `json:"comment"`
	Errors    []error                   `json:"errors"`
	Classes   map[string]slang.Class    `json:"classes"`
	Functions map[string]slang.Function `json:"functions"`
}

func NewModule() *Module {
	return &Module{
		Comment:   "",
		Classes:   map[string]slang.Class{},
		Functions: map[string]slang.Function{},
		Errors:    []error{},
	}
}

// func FromFiles(fpaths ...string) (*Module, error) {
// 	m := New()
// 	rootNames := []string{}
// 	fileComments := map[string]string{}

// 	for _, fpath := range fpaths {
// 		f, err := os.Open(fpath)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to open fail: %w", err)
// 		}

// 		results := parsing.Parse(f)
// 		if len(results.Errors) > 0 {
// 			for _, parseErr := range results.Errors {
// 				err := fmt.Errorf("%s (%d,%d): %w", fpath, parseErr.Token.Location.Line, parseErr.Token.Location.Column, parseErr.Error)

// 				m.Errors = append(m.Errors, err)
// 			}
// 		}

// 		if len(results.FileComment) > 0 {
// 			fileComments[fpath] = results.FileComment
// 		}

// 		for name, cls := range results.Classes {
// 			if slices.Contains(rootNames, name) {
// 				err := customerrs.NewErrDuplicateName(name)

// 				m.Errors = append(m.Errors, err)

// 				continue
// 			}

// 			m.Classes[name] = cls

// 			rootNames = append(rootNames, name)
// 		}

// 		for name, fn := range results.Functions {
// 			if slices.Contains(rootNames, name) {
// 				err := customerrs.NewErrDuplicateName(name)

// 				m.Errors = append(m.Errors, err)

// 				continue
// 			}

// 			m.Functions[name] = fn

// 			rootNames = append(rootNames, name)
// 		}
// 	}

// 	// combine file comments into a module comment
// 	for fpath, comment := range fileComments {

// 	}

// 	return m, nil
// }

func (m *Module) GetComment() string {
	return m.Comment
}

func (m *Module) GetClassNames() []string {
	return maps.Keys(m.Classes)
}

func (m *Module) GetClass(name string) (slang.Class, bool) {
	c, found := m.Classes[name]
	if !found {
		return nil, false
	}

	return c, true
}

func (m *Module) GetFunctionNames() []string {
	return maps.Keys(m.Functions)
}

func (m *Module) GetFunction(name string) (slang.Function, bool) {
	c, found := m.Functions[name]
	if !found {
		return nil, false
	}

	return c, true
}
