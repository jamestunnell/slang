package statements

import "github.com/jamestunnell/slang"

type Comment struct{}

func NewComment() *Statement {
	return NewStatement(slang.StatementCOMMENT, &Comment{})
}

func (c *Comment) GetName() (string, bool) {
	return "", false
}

func (c *Comment) IsEqual(other Core) bool {
	_, ok := other.(*Comment)

	return ok
}

func (c *Comment) Render(level int, w slang.CodeWriter) {
}
