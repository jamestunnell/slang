package parsing

import (
	"bufio"
	"errors"
	"strings"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/lexer"
)

var (
	errParseStringEmptyExpr     = errors.New("empty string interpolation expression")
	errParseStringMissingRBrace = errors.New("missing string interpolation rbrace")
)

// ParseString parses a quoted string value, with escape sequences and
// interpolation expressions.
func ParseString(val string) ([]slang.Expression, error) {
	exprs := []slang.Expression{}
	r := bufio.NewReader(strings.NewReader(val))

	ch, _, _ := r.ReadRune()

	collected := []rune{}

	waitingForRBrace := false

	for ch != 0 {
		if ch == '{' {
			if len(collected) > 0 {
				exprs = append(exprs, expressions.NewString(string(collected)))

				collected = []rune{}
			}

			waitingForRBrace = true
		} else if ch == '}' && waitingForRBrace {
			if len(collected) == 0 {
				return []slang.Expression{}, errParseStringEmptyExpr
			}

			p := NewExprParser(PrecedenceLOWEST)
			l := lexer.New(bufio.NewReader(strings.NewReader(string(collected))))

			p.Run(NewTokenSeq(l))

			if len(p.GetErrors()) > 0 {
				return []slang.Expression{}, p.GetErrors()[0].Error
			}

			exprs = append(exprs, p.Expr)

			collected = []rune{}

			waitingForRBrace = false
		} else {
			collected = append(collected, ch)
		}

		ch, _, _ = r.ReadRune()
	}

	if waitingForRBrace {
		return []slang.Expression{}, errParseStringMissingRBrace
	}

	if len(collected) > 0 {
		exprs = append(exprs, expressions.NewString(string(collected)))
	}

	return exprs, nil
}
