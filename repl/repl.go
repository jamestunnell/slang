package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/lexer"
	"github.com/jamestunnell/slang/parser"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := slang.NewEnvironment(nil)

	for {
		fmt.Fprint(out, prompt)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(strings.NewReader(line))
		p := parser.New(l)

		results := p.Run()
		if len(results.Errors) > 0 {
			for _, pErr := range results.Errors {
				l := pErr.Token.Location

				fmt.Fprintln(out, line)
				fmt.Fprintln(out, spaces(l.Column-1)+"^")

				fmt.Fprintf(out, "parse error at %s: %v\n", l, pErr.Error)
			}

			continue
		}

		for _, st := range results.Statements {
			obj, err := st.Eval(env)
			if err != nil {
				fmt.Fprintf(out, "runtime error: %v\n", err)

				break
			}

			fmt.Fprintln(out, obj.Inspect())
		}
	}
}

func spaces(n int) string {
	return strings.Repeat(" ", n)
}
