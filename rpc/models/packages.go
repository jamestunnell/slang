package models

import (
	"io/fs"

	"github.com/jamestunnell/slang"
)

type UpsertPackageArgs struct {
	Meta    slang.PackageMeta
	Archive slang.PackageArchive
}

type ListPackagesReply struct {
	Addresses []slang.PackageAddress
}

type GetPackageStateReply struct {
	State slang.PackageState
	Found bool
}

type GetPackageArchiveReply struct {
	Archive slang.PackageArchive
	Found   bool
}

type GetPackageFilesReply struct {
	Files fs.FS
	Found bool
}

type GetPackageASTReply struct {
	AST   slang.PackageAST
	Found bool
}

type GetPackageDependenciesReply struct {
	Dependencies []slang.PackageAddress
	Found        bool
}

type GetPackageAnalysisReply struct {
	Analysis slang.PackageAnalysis
	Found    bool
}

type GetPackageBytecodeReply struct {
	Bytecode slang.PackageBytecode
	Found    bool
}

type GetPackageFailureReply struct {
	Failure slang.PackageFailure
	Found   bool
}
