package slang

import (
	"io/fs"

	"github.com/google/uuid"
)

type VirtualMachine interface {
	GetName() string
	GetID() uuid.UUID

	IsRunning() bool

	UpsertPackage(PackageMeta, PackageArchive)
	RemovePackage(PackageAddress) bool
	ListPackages() []PackageAddress

	GetPackageState(PackageAddress) (PackageState, bool)
	GetPackageArchive(PackageAddress) (PackageArchive, bool)
	GetPackageFiles(PackageAddress) (fs.FS, bool)
	GetPackageAST(PackageAddress) (PackageAST, bool)
	GetPackageDependencies(PackageAddress) ([]PackageAddress, bool)
	GetPackageAnalysis(PackageAddress) (PackageAnalysis, bool)
	GetPackageBytecode(PackageAddress) (PackageBytecode, bool)
	GetPackageFailure(PackageAddress) (PackageFailure, bool)
}
