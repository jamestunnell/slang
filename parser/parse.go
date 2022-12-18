package parser

import (
	"strings"

	"github.com/jamestunnell/slang/lexer"
)

func Parse(input string) *ParseResults {
	r := strings.NewReader(input)
	l := lexer.New(r)
	p := New(l)

	return p.Run()
}
