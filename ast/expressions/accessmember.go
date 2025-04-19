package expressions

import (
	"github.com/jamestunnell/slang"
)

type AccessMember struct {
	Receiver *Expression `json:"receiver"`
	Member   string      `json:"member"`
}

func NewAccessMember(object *Expression, member string) *Expression {
	return NewExpression(slang.ExprACCESSMEMBER, &AccessMember{
		Receiver: object,
		Member:   member,
	})
}

func (a *AccessMember) IsEqual(other Core) bool {
	a2, ok := other.(*AccessMember)
	if !ok {
		return false
	}

	return a2.Receiver.IsEqual(a.Receiver) && a2.Member == a.Member
}

func (a *AccessMember) Render(level int, w slang.CodeWriter) {
	a.Receiver.Render(level, w)

	w.WriteString(".")
	w.WriteString(a.Member)
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
