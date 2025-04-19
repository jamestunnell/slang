package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/sliceutil"
	"go.uber.org/multierr"
)

func RunParser(l slang.Lexer, p Parser) error {
	toks := NewTokenSeq(l)

	if !p.Run(toks) {
		parseErrs := p.GetErrors()

		errs := sliceutil.Map(parseErrs, func(parseErr *ParseErr) error { return parseErr })

		return multierr.Combine(errs...)
	}

	return nil
}
