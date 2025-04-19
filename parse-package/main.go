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
	Root    string `arg:"positional"`
	Meta    string `arg:"-m" help:"package meta JSON" default:""`
	OutPath string `arg:"-o" help:"JSON output file path" default:"./ast.json"`
}

func main() {
	var args Args

	arg.MustParse(&args)

	rootInfo, err := os.Stat(args.Root)
	if err != nil {
		fmt.Printf("failed to stat root dir '%s': %v\n", args.Root, err)

		os.Exit(1)
	}

	if !rootInfo.IsDir() {
		fmt.Printf("package root is '%s' is not a dir\n", args.Root)

		os.Exit(1)
	}

	meta := loadMeta(&args)
	pkgFS := os.DirFS(args.Root)

	modules, err := parsing.ParsePackage(pkgFS, parsers.NewFileParser())
	if err != nil {
		fmt.Printf("failed to parse package: %v\n", err)

		os.Exit(1)
	}

	pkg := ast.NewPackage(meta, modules...)

	d, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal package JSON: %v\n", err)

		os.Exit(1)
	}

	err = os.WriteFile(args.OutPath, d, 0777)
	if err != nil {
		fmt.Printf("failed to write package JSON: %v\n", err)

		os.Exit(1)
	}

	fmt.Printf("wrote package JSON to '%s'\n", args.OutPath)
}

func loadMeta(args *Args) slang.PackageMeta {
	var metaData []byte

	if args.Meta != "" {
		metaData = []byte(args.Meta)
	} else {
		var err error

		metaPath := path.Join(args.Root, "meta.json")
		if _, err = os.Stat(metaPath); err != nil && os.IsNotExist(err) {
			fmt.Printf("missing meta JSON files '%s'", metaPath)

			os.Exit(1)
		}

		metaData, err = os.ReadFile(metaPath)
		if err != nil {
			fmt.Printf("failed to read meta JSON: %v\n", err)

			os.Exit(1)
		}
	}

	var meta slang.PackageMeta

	if err := json.Unmarshal(metaData, &meta); err != nil {
		fmt.Printf("failed to unmarshal meta JSON: %v\n", err)

		os.Exit(1)
	}

	return meta
}
