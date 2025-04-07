package sliceutil

func Map[T, U any](ts []T, convert func(T) U) []U {
	us := make([]U, len(ts))

	for i, t := range ts {
		us[i] = convert(t)
	}

	return us
}
