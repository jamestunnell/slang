package parsing

import (
	"bufio"
	"io"

	"github.com/jamestunnell/slang/lexing"
)

func RunParser(p Parser, r io.Reader) bool {
	l := lexing.NewLexer(bufio.NewReader(r))
	toks := NewTokenSeq(l)

	return p.Run(toks)
}
