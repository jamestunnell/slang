package lexer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jamestunnell/slang"
)

func Scan(r io.Reader) []*slang.Token {
	l := New(bufio.NewReader(r))
	toks := []*slang.Token{}
	keepGoing := func(tok *slang.Token) bool {
		return tok != nil && tok.Info.Type() != slang.TokenEOF
	}

	for tok := l.NextToken(); keepGoing(tok); tok = l.NextToken() {
		toks = append(toks, tok)
	}

	return toks
}

func ScanFile(path string) ([]*slang.Token, error) {
	f, err := os.Open(path)
	if err != nil {
		return []*slang.Token{}, fmt.Errorf("failed to read file '%s': %w", path, err)
	}

	return Scan(bufio.NewReader(f)), nil
}

func ScanString(input string) []*slang.Token {
	return Scan(strings.NewReader(input))
}
