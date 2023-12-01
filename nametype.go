package slang

type NameType struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (nt *NameType) GetName() string {
	return nt.Name
}

func (nt *NameType) GetType() string {
	return nt.Type
}

func (nt *NameType) Equal(other *NameType) bool {
	return nt.Name == other.Name && nt.Type == other.Type
}
