package objects

type Class interface {
	Name() string
	AddInstanceMethod(string, Method)
}
type Method interface {
	ParamNames() []string
	Run(args []Object) (Object, error)
}
