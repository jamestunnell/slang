package customerrs

import "fmt"

type ErrDuplicate struct {
	name string
}

func NewErrDuplicateName(name string) *ErrDuplicate {
	return &ErrDuplicate{
		name: name,
	}
}

func (err *ErrDuplicate) Error() string {
	const strFmt = "name %s is already used"

	return fmt.Sprintf(strFmt, err.name)
}
