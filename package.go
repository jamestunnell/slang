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
	Path         string   `json:"path"`
	Version      string   `json:"version"`
	Dependencies []string `json:"dependencies,omitempty"`
}

type PackageRepoEntry struct {
	Archive PackageArchive
	Errors  []error
	AST     Package
}

func (m PackageMeta) String() string {
	if m.Version == "" {
		return m.Path
	}

	return fmt.Sprintf("%s-%s", m.Path, m.Version)
}

func (m PackageMeta) IsEqual(other PackageMeta) bool {
	if m.Path != other.Path {
		return false
	}

	return m.Version == other.Version
}
