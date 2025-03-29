package slang

type Struct interface {
	GetName() string
	GetComment() string

	GetFieldNames() []string
	GetFieldType(name string) (string, bool)
}
