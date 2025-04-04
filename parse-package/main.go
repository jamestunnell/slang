package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	arg "github.com/alexflint/go-arg"
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
)

type Args struct {
	PackageRoot string `arg:"positional"`

	OutDir string `arg:"-o" help:"package JSON output dir" default:"."`
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

	pkgFS := os.DirFS(args.PackageRoot)

	modules, err := parsing.ParsePackage(pkgFS, parsers.NewFileParser())
	if err != nil {
		fmt.Printf("failed to parse package: %v\n", err)

		os.Exit(1)
	}

	pkgInfo := &slang.PackageInfo{
		Name: path.Base(args.PackageRoot),
	}
	pkg := ast.NewPackage(pkgInfo, modules...)

	d, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal package JSON: %v\n", err)

		os.Exit(1)
	}

	outpath := path.Join(args.OutDir, pkg.PackageInfo.Name+".json")

	err = os.WriteFile(outpath, d, 0777)
	if err != nil {
		fmt.Printf("failed to write package JSON: %v\n", err)

		os.Exit(1)
	}

	fmt.Printf("wrote package JSON to '%s'\n", outpath)
}
