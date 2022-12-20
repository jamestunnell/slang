package slang

import "fmt"

type ErrArrayBounds struct {
	index, size int64
}

func NewErrArrayBounds(index, size int64) *ErrArrayBounds {
	return &ErrArrayBounds{
		index: index,
		size:  size,
	}
}

func (err *ErrArrayBounds) Error() string {
	const strFmt = "index %d is out of bounds for a size %d array"

	return fmt.Sprintf(strFmt, err.index, err.size)
}
