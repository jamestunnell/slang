package slang

type PackageInfo struct {
	Name string `json:"name"`
}

func (info *PackageInfo) GetName() string {
	return info.Name
}
