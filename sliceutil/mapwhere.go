package sliceutil

func MapWhere[T, U any](
	ts []T,
	keep func(T) bool,
	convert func(T) U,
) []U {
	us := []U{}

	for _, t := range ts {
		if keep(t) {
			us = append(us, convert(t))
		}
	}

	return us
}
