package slang

import (
	"fmt"
	"io/fs"
)

type Package interface {
	GetMeta() PackageMeta
	GetModules() ([]Module, bool)
}

type PackageArchive interface {
	GetMeta() PackageMeta
	GetDataFormat() string
	GetDigestType() string

	Pack(fs.FS) error
	Unpack() (fs.FS, error)
}

type PackageMeta struct {
	Path    string `json:"path"`
	Commit  string `json:"commit,omitempty"`
	Version string `json:"version"`
}

type PackageRepoEntry struct {
	Archive PackageArchive
	Errors  []error
	AST     Package
}

func (m PackageMeta) String() string {
	return fmt.Sprintf("%s-%s-%s", m.Path, m.Version, m.Commit)
}

func (m PackageMeta) IsEqual(other PackageMeta) bool {
	if m.Path != other.Path {
		return false
	}

	if m.Version != other.Version {
		return false
	}

	return m.Commit == other.Commit
}
