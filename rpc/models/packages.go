package models

import (
	"github.com/jamestunnell/slang"
)

type AddPackageArgs struct {
	Archive slang.PackageArchive
}

type AddPackageReply struct {
	ErrorMsg string
}

type RemovePackageArgs struct {
	Meta slang.PackageMeta
}

type RemovePackageReply struct {
	Removed bool
}

type GetArchiveArgs struct {
	Meta slang.PackageMeta
}

type GetArchiveReply struct {
	Archive slang.PackageArchive
	Found   bool
}

type ListPackagesArgs struct {
}

type ListPackagesReply struct {
	Metas []slang.PackageMeta
}
