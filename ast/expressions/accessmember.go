package expressions

import (
	"github.com/jamestunnell/slang"
)

type AccessMember struct {
	*Base

	Receiver slang.Expression `json:"receiver"`
	Member   string           `json:"member"`
}

func NewAccessMember(object slang.Expression, member string) slang.Expression {
	return &AccessMember{
		Base:     NewBase(slang.ExprACCESSMEMBER),
		Receiver: object,
		Member:   member,
	}
}

func (c *AccessMember) Equal(other slang.Expression) bool {
	c2, ok := other.(*AccessMember)
	if !ok {
		return false
	}

	return c2.Receiver.Equal(c.Receiver) && c2.Member == c.Member
}

// func (c *Member) Eval(env *slang.Environment) (slang.Object, error) {
// 	obj, err := c.Object.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	vals := make([]slang.Object, len(c.Arguments))
// 	for i := 0; i < len(c.Arguments); i++ {
// 		val, err := c.Arguments[i].Eval(env)
// 		if err != nil {
// 			return objects.NULL(), err
// 		}

// 		vals[i] = val
// 	}

// 	return obj.Send(c.MethodName.Name, vals...)
// }
