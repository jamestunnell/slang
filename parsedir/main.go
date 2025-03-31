package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/parsing"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("missing directory arg")

		os.Exit(1)
	}

	if len(os.Args) > 2 {
		fmt.Printf("too many args")

		os.Exit(1)
	}

	dirpath := os.Args[1]

	entries, err := os.ReadDir(dirpath)
	if err != nil {
		fmt.Printf("failed to read dir '%s': %v", dirpath, err)

		os.Exit(1)
	}

	mod := ast.NewModule()

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if path.Ext(entry.Name()) != ".sl" {
			continue
		}

		fname := path.Join(dirpath, entry.Name())

		f, err := os.Open(fname)
		if err != nil {
			fmt.Printf("failed to open file '%s': %v", fname, err)

			os.Exit(1)
		}

		fmt.Printf("parsing %s -> ", fname)

		l := lexing.NewLexer(bufio.NewReader(f))
		toks := parsing.NewTokenSeq(l)
		p := parsing.NewFileParser()

		if !p.Run(toks) {
			fmt.Print("failed\n")

			for _, parseErr := range p.GetErrors() {
				fmt.Printf("* %v\n", parseErr)
			}

			os.Exit(1)
		}

		fmt.Print("success\n")

		mod.Statements = append(mod.Statements, p.Statements...)
	}

	d, err := json.MarshalIndent(mod, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal module JSON: %v", err)

		os.Exit(1)
	}

	outpath := path.Join(path.Dir(dirpath), "module.json")

	err = os.WriteFile(outpath, d, 0777)
	if err != nil {
		fmt.Printf("failed to write module JSON: %v", err)

		os.Exit(1)
	}

	fmt.Printf("wrote module JSON to '%s'\n", outpath)
}
