package statements

import "github.com/jamestunnell/slang"

func WithComment(s slang.Statement, lines ...string) slang.Statement {
	s.SetComment(lines)

	return s
}
