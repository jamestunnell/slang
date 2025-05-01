package server

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/rpc/models"
)

type Packages struct {
	VM slang.VirtualMachine
}

func (api *Packages) Upsert(args *models.UpsertPackageArgs, reply *models.Empty) error {
	api.VM.UpsertPackage(args.Meta, args.Archive)

	return nil
}

func (api *Packages) Remove(addr *slang.PackageAddress, removed *bool) error {
	*removed = api.VM.RemovePackage(*addr)

	return nil
}

func (api *Packages) List(args *models.Empty, reply *models.ListPackagesReply) error {
	reply.Addresses = api.VM.ListPackages()

	return nil
}

func (api *Packages) GetState(
	addr *slang.PackageAddress,
	reply *models.GetPackageStateReply,
) error {
	reply.State, reply.Found = api.VM.GetPackageState(*addr)

	return nil
}

func (api *Packages) GetArchive(
	addr *slang.PackageAddress,
	reply *models.GetPackageArchiveReply,
) error {
	reply.Archive, reply.Found = api.VM.GetPackageArchive(*addr)

	return nil
}

func (api *Packages) GetFiles(
	addr *slang.PackageAddress,
	reply *models.GetPackageFilesReply,
) error {
	reply.Files, reply.Found = api.VM.GetPackageFiles(*addr)

	return nil
}

func (api *Packages) GetAST(
	addr *slang.PackageAddress,
	reply *models.GetPackageASTReply,
) error {
	reply.AST, reply.Found = api.VM.GetPackageAST(*addr)

	return nil
}

func (api *Packages) GetDependencies(
	addr *slang.PackageAddress,
	reply *models.GetPackageDependenciesReply,
) error {
	reply.Dependencies, reply.Found = api.VM.GetPackageDependencies(*addr)

	return nil
}

func (api *Packages) GetAnalysis(
	addr *slang.PackageAddress,
	reply *models.GetPackageAnalysisReply,
) error {
	reply.Analysis, reply.Found = api.VM.GetPackageAnalysis(*addr)

	return nil
}

func (api *Packages) GetBytecode(
	addr *slang.PackageAddress,
	reply *models.GetPackageBytecodeReply,
) error {
	reply.Bytecode, reply.Found = api.VM.GetPackageBytecode(*addr)

	return nil
}

func (api *Packages) GetFailure(
	addr *slang.PackageAddress,
	reply *models.GetPackageFailureReply,
) error {
	reply.Failure, reply.Found = api.VM.GetPackageFailure(*addr)

	return nil
}
