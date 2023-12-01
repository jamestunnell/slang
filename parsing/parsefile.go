package parsing

import (
	"bufio"
	"io"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/lexer"
)

func ParseFile(r io.Reader) ([]slang.Statement, []*ParseErr) {
	l := lexer.New(bufio.NewReader(r))
	toks := NewTokenSeq(l)
	p := NewFileParser()

	p.Run(toks)

	return p.Statements, p.GetErrors()
}
