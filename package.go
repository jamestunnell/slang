package slang

import (
	"fmt"
	"io/fs"
	"slices"
	"strings"
)

type PackageState int
type PackageAction int

type PackageMeta struct {
	Address      PackageAddress   `json:"address"`
	Dependencies []PackageAddress `json:"dependencies,omitempty"`
}

type PackageArchive interface {
	GetDigest() string
	GetDataFormat() string
	GetDigestType() string

	Pack(fs.FS) error
	Unpack() (fs.FS, error)
}

type PackageAST interface {
	GetModules() []Module
}

type PackageAnalysis struct {
}

type PackageBytecode struct {
}

type PackageFailure struct {
	FailedAction PackageAction
	ErrorMsg     string
}

type PackageAddress struct {
	Path    string `json:"path"`
	Version string `json:"version"`
}

const (
	PkgAdded PackageState = iota
	PkgUnpacked
	PkgParsed
	PkgResolved
	PkgAnalyzed
	PkgCompiled
	PkgFailed

	PkgUnpack PackageAction = iota
	PkgParse
	PkgResolve
	PkgAnalyze
	PkgCompile
)

func (s PackageState) String() string {
	var str string

	switch s {
	case PkgAdded:
		str = "ADDED"
	case PkgUnpacked:
		str = "UNPACKED"
	case PkgParsed:
		str = "PARSED"
	case PkgResolved:
		str = "RESOLVED"
	case PkgAnalyzed:
		str = "ANALYZED"
	case PkgCompiled:
		str = "COMPILED"
	case PkgFailed:
		str = "FAILED"
	}

	return str
}

func (s PackageAction) String() string {
	var str string

	switch s {
	case PkgUnpack:
		str = "UNPACK"
	case PkgParse:
		str = "PARSE"
	case PkgResolve:
		str = "RESOLVE"
	case PkgAnalyze:
		str = "ANALYZE"
	case PkgCompile:
		str = "COMPILE"
	}

	return str
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

func (addr *PackageAddress) Parse(s string) {
	parts := strings.SplitN(s, "@", 2)
	switch len(parts) {
	case 0:
	case 1:
		addr.Path = s
	default:
		addr.Path = parts[0]
		addr.Version = parts[1]
	}
}

func PackageAddressesEqual(a, b PackageAddress) bool {
	return (a.Path == b.Path) && (a.Version == b.Version)
}
