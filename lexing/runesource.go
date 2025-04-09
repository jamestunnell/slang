package lexing

import (
	"bufio"
	"io"
)

type RuneSource interface {
	NextRune() (rune, bool)
}

type runeSource struct {
	reader io.RuneReader
}

func NewRuneSource(r io.Reader) RuneSource {
	return &runeSource{reader: bufio.NewReader(r)}
}

func (rs *runeSource) NextRune() (rune, bool) {
	r, _, err := rs.reader.ReadRune()

	return r, err == nil
}
