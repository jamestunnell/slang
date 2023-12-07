package compilation

import "github.com/jamestunnell/slang"

type StmtSeq interface {
	Advance()
	Current() (slang.Statement, bool)
	Reset()
}

type stmtSeq struct {
	Statements []slang.Statement
	Index      int
	Length     int
}

func NewEmptyStmtSeq() StmtSeq {
	return NewStmtSeq([]slang.Statement{})
}

func NewStmtSeq(stmts []slang.Statement) StmtSeq {
	return &stmtSeq{
		Statements: stmts,
		Index:      0,
		Length:     len(stmts),
	}
}

func (seq *stmtSeq) Advance() {
	seq.Index++
}

func (seq *stmtSeq) Current() (slang.Statement, bool) {
	if seq.Index >= seq.Length {
		return nil, false
	}

	stmt := seq.Statements[seq.Index]

	return stmt, true
}

func (seq *stmtSeq) Reset() {
	seq.Index = 0
}
