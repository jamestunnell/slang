package slang

type Structure interface {
	GetName() string
	GetComment() string

	GetFieldNames() []string
	GetFieldType(name string) (string, bool)
}
