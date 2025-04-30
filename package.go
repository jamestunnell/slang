package slang

import (
	"fmt"
	"io/fs"
	"slices"
)

type Package interface {
	GetMeta() PackageMeta
	GetModules() []Module
}

type PackageArchive interface {
	GetMeta() PackageMeta
	GetDigest() string
	GetDataFormat() string
	GetDigestType() string

	Pack(fs.FS) error
	Unpack() (fs.FS, error)
}

type PackageMeta struct {
	Address      PackageAddress   `json:"address"`
	Dependencies []PackageAddress `json:"dependencies,omitempty"`
}

type PackageAddress struct {
	Path    string `json:"path"`
	Version string `json:"version"`
}

type PackageRepoEntry struct {
	Archive PackageArchive
	Errors  []error
	AST     Package
}

func (m PackageMeta) String() string {
	return m.Address.String()
}

func (m PackageMeta) IsEqual(other PackageMeta) bool {
	if !PackageAddressesEqual(m.Address, other.Address) {
		return false
	}

	return slices.EqualFunc(m.Dependencies, other.Dependencies, PackageAddressesEqual)
}

func (addr PackageAddress) String() string {
	if addr.Version == "" {
		return addr.Path
	}

	return fmt.Sprintf("%s@%s", addr.Path, addr.Version)
}

func PackageAddressesEqual(a, b PackageAddress) bool {
	return (a.Path == b.Path) && (a.Version == b.Version)
}
