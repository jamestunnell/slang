package parsing

import (
	"bufio"
	"io"

	"go.uber.org/multierr"

	"github.com/jamestunnell/slang/lexing"
)

func RunParser(p Parser, r io.Reader) error {
	l := lexing.NewLexer(bufio.NewReader(r))
	toks := NewTokenSeq(l)

	if !p.Run(toks) {
		parseErrs := p.GetErrors()
		errs := make([]error, len(parseErrs))

		for i, parseErr := range parseErrs {
			errs[i] = parseErr
		}

		return multierr.Combine(errs...)
	}

	return nil
}
