package compilation

import "github.com/jamestunnell/slang"

type StmtSeq interface {
	Reset()

	Next() (slang.Statement, bool)
}
