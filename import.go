package slang

type Import interface {
	GetRename() string
	GetPathParts() []string
}
