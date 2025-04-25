package slang

import "github.com/google/uuid"

type VirtualMachine interface {
	GetInfo() VMInfo

	IsRunning() bool

	ListPackages() []PackageMeta
	GetPackage(PackageMeta) (PackageArchive, bool)
	AddPackage(PackageArchive) error
	RemovePackage(PackageMeta) bool

	EvaluateExpr(Expression) (Object, error)
}

type VMInfo struct {
	Name string
	ID   uuid.UUID
}
