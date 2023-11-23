package slang

type Package interface {
	GetVersion() string

	GetModulePaths() []string
	GetModule(path string) (Module, bool)
}
