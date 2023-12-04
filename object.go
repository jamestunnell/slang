package slang

type Object interface {
	Equal(Object) bool
	// Inspect() string
	Send(name string, args ...Object) (Object, error)
}

func ObjectsEqual(a, b []Object) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, obj := range a {
		if !obj.Equal(b[idx]) {
			return false
		}
	}

	return true
}
