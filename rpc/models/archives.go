package models

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/archives"
)

type AddTarGzArgs struct {
	Archive *archives.TarGz
}

type AddTarGzReply struct {
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
	Archive *archives.TarGz
	Found   bool
}

type ListArchivesArgs struct {
}

type ListArchivesReply struct {
	Metas []slang.PackageMeta
}
