package ast

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type Module struct {
	Path       string            `json:"path"`
	Statements []slang.Statement `json:"statements"`
	// // Errors    []error          `json:"errors"`
	// Structures []slang.Structure `json:"structures"`
	// Functions  []slang.Function  `json:"functions"`
}

func NewModule(path string, stmts ...slang.Statement) *Module {
	return &Module{
		Path:       path,
		Statements: stmts,
		// Structures: []slang.Structure{},
		// Functions:  []slang.Function{},
		// // Errors:    []error{},
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

func (m *Module) GetPath() string {
	return m.Path
}

func (m *Module) GetStructures() []slang.Structure {
	structs := []slang.Structure{}

	for _, s := range m.Statements {
		if str, ok := s.(*statements.Struct); ok {
			structs = append(structs, str)
		}
	}

	return structs
}

func (m *Module) GetFunctions() []slang.Function {
	funcs := []slang.Function{}

	for _, s := range m.Statements {
		if fn, ok := s.(*statements.Func); ok {
			funcs = append(funcs, fn)
		}
	}

	return funcs
}
