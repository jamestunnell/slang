package parsing

import (
	"github.com/jamestunnell/slang"
	"go.uber.org/multierr"
)

func RunParser(l slang.Lexer, p Parser) error {
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
