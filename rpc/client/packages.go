package client

import (
	"io/fs"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/rpc/models"
)

func (client *Client) UpsertPackage(meta slang.PackageMeta, archive slang.PackageArchive) {
	const method = "Packages.Upsert"

	args := &models.UpsertPackageArgs{Meta: meta, Archive: archive}

	var reply models.Empty

	if err := client.rpcClient.Call(method, args, &reply); err != nil {
		client.logFailedMethodCall(method, err)
	}
}

func (client *Client) RemovePackage(addr slang.PackageAddress) bool {
	const method = "Packages.Remove"

	var removed bool

	if err := client.rpcClient.Call(method, &addr, &removed); err != nil {
		client.logFailedMethodCall(method, err)

		return false
	}

	return removed
}

func (client *Client) ListPackages() []slang.PackageAddress {
	const method = "Packages.List"
	args := &models.Empty{}

	var reply models.ListPackagesReply

	if err := client.rpcClient.Call(method, args, &reply); err != nil {
		client.logFailedMethodCall(method, err)

		return []slang.PackageAddress{}
	}

	return reply.Addresses
}

func (client *Client) GetPackageState(addr slang.PackageAddress) (slang.PackageState, bool) {
	const method = "Packages.GetState"

	var reply models.GetPackageStateReply

	if err := client.rpcClient.Call(method, &addr, &reply); err != nil {
		client.logFailedMethodCall(method, err)

		return 0, false
	}

	return reply.State, reply.Found
}

func (client *Client) GetPackageArchive(addr slang.PackageAddress) (slang.PackageArchive, bool) {
	const method = "Packages.GetArchive"

	var reply models.GetPackageArchiveReply

	if err := client.rpcClient.Call(method, &addr, &reply); err != nil {
		client.logFailedMethodCall(method, err)

		return nil, false
	}

	return reply.Archive, reply.Found
}

func (client *Client) GetPackageFiles(addr slang.PackageAddress) (fs.FS, bool) {
	const method = "Packages.GetFiles"

	var reply models.GetPackageFilesReply

	if err := client.rpcClient.Call(method, &addr, &reply); err != nil {
		client.logFailedMethodCall(method, err)

		return nil, false
	}

	return reply.Files, reply.Found
}

func (client *Client) GetPackageAST(addr slang.PackageAddress) (slang.PackageAST, bool) {
	const method = "Packages.GetAST"

	var reply models.GetPackageASTReply

	if err := client.rpcClient.Call(method, &addr, &reply); err != nil {
		client.logFailedMethodCall(method, err)

		return nil, false
	}

	return reply.AST, reply.Found
}

func (client *Client) GetPackageDependencies(addr slang.PackageAddress) ([]slang.PackageAddress, bool) {
	const method = "Packages.GetDependencies"

	var reply models.GetPackageDependenciesReply

	if err := client.rpcClient.Call(method, &addr, &reply); err != nil {
		client.logFailedMethodCall(method, err)

		return nil, false
	}

	return reply.Dependencies, reply.Found
}

func (client *Client) GetPackageAnalysis(addr slang.PackageAddress) (slang.PackageAnalysis, bool) {
	const method = "Packages.GetAnalysis"

	var reply models.GetPackageAnalysisReply

	if err := client.rpcClient.Call(method, &addr, &reply); err != nil {
		client.logFailedMethodCall(method, err)

		return slang.PackageAnalysis{}, false
	}

	return reply.Analysis, reply.Found
}

func (client *Client) GetPackageBytecode(addr slang.PackageAddress) (slang.PackageBytecode, bool) {
	const method = "Packages.GetBytecode"

	var reply models.GetPackageBytecodeReply

	if err := client.rpcClient.Call(method, &addr, &reply); err != nil {
		client.logFailedMethodCall(method, err)

		return slang.PackageBytecode{}, false
	}

	return reply.Bytecode, reply.Found
}

func (client *Client) GetPackageFailure(addr slang.PackageAddress) (slang.PackageFailure, bool) {
	const method = "Packages.GetFailure"

	var reply models.GetPackageFailureReply

	if err := client.rpcClient.Call(method, &addr, &reply); err != nil {
		client.logFailedMethodCall(method, err)

		return slang.PackageFailure{}, false
	}

	return reply.Failure, reply.Found
}
