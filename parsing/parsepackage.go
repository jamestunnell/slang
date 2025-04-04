package parsing

import (
	"bufio"
	"fmt"
	"io/fs"
	"path"

	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/statements"
)

func ParsePackage(rootFS fs.FS, fp FileParser) ([]*ast.Module, error) {
	moduleStmts := map[string][]*statements.Statement{}
	err := fs.WalkDir(rootFS, ".", func(entryPath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			walkErr = fmt.Errorf("failed to walk into dir '%s': %v", entryPath, walkErr)

			return walkErr
		}

		if entry.IsDir() || path.Ext(entry.Name()) != ".sl" {
			return nil
		}

		f, openErr := rootFS.Open(entryPath)
		if openErr != nil {
			openErr = fmt.Errorf("failed to open file '%s': %v", entryPath, openErr)

			return openErr
		}

		fmt.Printf("parsing %s -> ", entryPath)

		if parseErr := RunParser(fp, bufio.NewReader(f)); parseErr != nil {
			return parseErr
		}

		log.Debug().Str("path", entryPath).Msg("parsed package file")

		dir := path.Dir(entryPath)

		moduleStmts[dir] = append(moduleStmts[dir], fp.GetStatements()...)

		return nil
	})
	if err != nil {
		return nil, err
	}

	modules := []*ast.Module{}

	for relativePath, stmts := range moduleStmts {
		modules = append(modules, ast.NewModule(relativePath, stmts...))
	}

	return modules, nil
}
