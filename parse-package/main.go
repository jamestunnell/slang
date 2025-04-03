package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"

	arg "github.com/alexflint/go-arg"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/parsing"
)

type Args struct {
	PackageRoot string `arg:"positional"`

	PackageName string `arg:"-n" help:"package name (root dir name by default)" default:""`
	OutDir      string `arg:"-o" help:"package JSON output dir" default:"."`
}

func main() {
	var args Args

	arg.MustParse(&args)

	rootInfo, err := os.Stat(args.PackageRoot)
	if err != nil {
		fmt.Printf("failed to stat root dir '%s': %v\n", args.PackageRoot, err)

		os.Exit(1)
	}

	if !rootInfo.IsDir() {
		fmt.Printf("package root is '%s' is not a dir\n", args.PackageRoot)

		os.Exit(1)
	}

	pkgName := args.PackageName
	if pkgName == "" {
		pkgName = path.Base(args.PackageRoot)
	}

	moduleStmts := map[string][]slang.Statement{}
	fileSystem := os.DirFS(args.PackageRoot)

	fs.WalkDir(fileSystem, ".", func(entryPath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			fmt.Printf("failed to walk into dir '%s': %v\n", entryPath, walkErr)

			os.Exit(1)
		}

		if entry.IsDir() || path.Ext(entry.Name()) != ".sl" {
			return nil
		}

		f, err := os.Open(path.Join(args.PackageRoot, entryPath))
		if err != nil {
			fmt.Printf("failed to open file '%s': %v\n", entryPath, err)

			os.Exit(1)
		}

		fmt.Printf("parsing %s -> ", entryPath)

		l := lexing.NewLexer(bufio.NewReader(f))
		toks := parsing.NewTokenSeq(l)
		p := parsing.NewFileParser()

		if !p.Run(toks) {
			fmt.Print("failed\n")

			for _, parseErr := range p.GetErrors() {
				fmt.Printf("* %s: %v\n", parseErr.Token.Location, parseErr.Error)
			}

			os.Exit(1)
		}

		fmt.Print("success\n")

		dir := path.Dir(entryPath)

		moduleStmts[dir] = append(moduleStmts[dir], p.Statements...)

		return nil
	})

	modules := []*ast.Module{}

	for relativePath, stmts := range moduleStmts {
		modules = append(modules, ast.NewModule(relativePath, stmts...))
	}

	pkgInfo := &slang.PackageInfo{Name: pkgName}
	pkg := ast.NewPackage(pkgInfo, modules...)

	d, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal package JSON: %v\n", err)

		os.Exit(1)
	}

	outpath := path.Join(args.OutDir, pkgInfo.Name+".json")

	err = os.WriteFile(outpath, d, 0777)
	if err != nil {
		fmt.Printf("failed to write package JSON: %v\n", err)

		os.Exit(1)
	}

	fmt.Printf("wrote package JSON to '%s'\n", outpath)
}
