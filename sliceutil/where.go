package sliceutil

func Where[T any](ts []T, keep func(T) bool) []T {
	kept := []T{}

	for _, t := range ts {
		if keep(t) {
			kept = append(kept, t)
		}
	}

	return kept
}
