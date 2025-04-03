package slang

type Structure interface {
	GetName() string
	GetComment() string
	GetFields() []Field
}
