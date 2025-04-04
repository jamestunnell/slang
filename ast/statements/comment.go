package statements

import "github.com/jamestunnell/slang"

type Comment struct{}

func NewComment() *Statement {
	return NewStatement(slang.StatementCOMMENT, &Comment{})
}

func (c *Comment) IsEqual(other Core) bool {
	_, ok := other.(*Comment)

	return ok
}
