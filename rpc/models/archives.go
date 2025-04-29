package models

import (
	"github.com/jamestunnell/slang"
)

type AddArchiveArgs struct {
	Archive slang.PackageArchive
}

type AddArchiveReply struct {
}

type RemoveArchiveArgs struct {
	Meta slang.PackageMeta
}

type RemoveArchiveReply struct {
	Removed bool
}

type GetArchiveArgs struct {
	Meta slang.PackageMeta
}

type GetArchiveReply struct {
	Archive slang.PackageArchive
	Found   bool
}

type ListArchivesArgs struct {
}

type ListArchivesReply struct {
	Metas []slang.PackageMeta
}
