package statements

import (
	"strings"

	"github.com/jamestunnell/slang"
)

type Base struct {
	CommentLines []string            `json:"commentLines,omitempty"`
	StmtType     slang.StatementType `json:"statementType"`
}

func NewBase(typ slang.StatementType) *Base {
	return &Base{StmtType: typ}
}

func (b *Base) SetComment(lines []string) {
	b.CommentLines = lines
}

func (b *Base) GetComment() string {
	return strings.Join(b.CommentLines, "")
}

func (b *Base) Type() slang.StatementType {
	return b.StmtType
}
