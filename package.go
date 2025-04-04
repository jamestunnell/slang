package slang

type Package interface {
	GetName() string
	GetModules() ([]Module, bool)
}
