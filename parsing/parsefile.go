package parsing

import (
	"bufio"
	"io"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/lexing"
)

func ParseFile(r io.Reader) ([]slang.Statement, []*ParseErr) {
	l := lexing.NewLexer(bufio.NewReader(r))
	toks := NewTokenSeq(l)
	p := NewFileParser()

	p.Run(toks)

	return p.Statements, p.GetErrors()
}
