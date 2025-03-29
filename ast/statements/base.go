package statements

import (
	"github.com/jamestunnell/slang"
)

type Base struct {
	CommentLines []string
	StmtType     slang.StatementType `json:"type"`
}

func NewBase(typ slang.StatementType) *Base {
	return &Base{StmtType: typ}
}

func (b *Base) SetComment(lines []string) {
	b.CommentLines = lines
}

func (b *Base) GetComment() []string {
	return b.CommentLines
}

func (b *Base) Type() slang.StatementType {
	return b.StmtType
}
