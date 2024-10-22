package lexing

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Scan(r io.Reader) []*Token {
	l := NewLexer(bufio.NewReader(r))
	toks := []*Token{}
	keepGoing := func(tok *Token) bool {
		return tok != nil && tok.Type != TokenEOF
	}

	for tok := l.NextToken(); keepGoing(tok); tok = l.NextToken() {
		toks = append(toks, tok)
	}

	return toks
}

func ScanFile(path string) ([]*Token, error) {
	f, err := os.Open(path)
	if err != nil {
		return []*Token{}, fmt.Errorf("failed to read file '%s': %w", path, err)
	}

	return Scan(bufio.NewReader(f)), nil
}

func ScanString(input string) []*Token {
	return Scan(strings.NewReader(input))
}
