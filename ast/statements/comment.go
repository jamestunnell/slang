package statements

import "github.com/jamestunnell/slang"

type Comment struct {
	*Base

	Lines []string
}

func NewComment(lines ...string) *Comment {
	return &Comment{
		Base:  NewBase(slang.StatementCOMMENT),
		Lines: lines,
	}
}

func (c *Comment) Equal(other slang.Statement) bool {
	c2, ok := other.(*Comment)
	if !ok {
		return false
	}

	if len(c.Lines) != len(c2.Lines) {
		return false
	}

	for i, line := range c.Lines {
		if line != c2.Lines[i] {
			return false
		}
	}

	return true
}
