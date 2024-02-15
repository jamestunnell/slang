package expressions

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type AccessMember struct {
	*Base

	Object slang.Expression `json:"object"`
	Member string           `json:"member"`
}

func NewAccessMember(object slang.Expression, member string) slang.Expression {
	return &AccessMember{
		Base:   NewBase(slang.ExprACCESSMEMBER),
		Object: object,
		Member: member,
	}
}

func (c *AccessMember) Equal(other slang.Expression) bool {
	c2, ok := other.(*AccessMember)
	if !ok {
		return false
	}

	return c2.Object.Equal(c.Object) && c2.Member == c.Member
}

func (c *AccessMember) Eval(env slang.Environment) (slang.Object, error) {
	obj, err := c.Object.Eval(env)
	if err != nil {
		return nil, err
	}

	memberVal, found := obj.Get(c.Member)
	if !found {
		return nil, customerrs.NewErrObjectNotFound(c.Member)
	}

	return memberVal, nil
}
