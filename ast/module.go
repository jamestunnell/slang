package ast

import (
	"github.com/jamestunnell/slang"
)

type Module struct {
	Comment   string           `json:"comment"`
	Errors    []error          `json:"errors"`
	Structs   []slang.Struct   `json:"structs"`
	Functions []slang.Function `json:"functions"`
}

func NewModule() *Module {
	return &Module{
		Comment:   "",
		Structs:   []slang.Struct{},
		Functions: []slang.Function{},
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

func (m *Module) GetStructNames() []string {
	names := make([]string, len(m.Structs))
	for i, s := range m.Structs {
		names[i] = s.GetName()
	}

	return names
}

func (m *Module) GetStruct(name string) (slang.Struct, bool) {
	for _, s := range m.Structs {
		if s.GetName() == name {
			return s, true
		}
	}

	return nil, false
}

func (m *Module) GetFunctionNames() []string {
	names := make([]string, len(m.Functions))
	for i, s := range m.Functions {
		names[i] = s.GetName()
	}

	return names
}

func (m *Module) GetFunction(name string) (slang.Function, bool) {
	for _, f := range m.Functions {
		if f.GetName() == name {
			return f, true
		}
	}

	return nil, false
}
