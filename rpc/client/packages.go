package client

import (
	"errors"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/rpc/models"
)

func (client *Client) AddPackage(archive slang.PackageArchive) error {
	const method = "Packages.Add"

	args := &models.AddPackageArgs{Archive: archive}

	var reply models.AddPackageReply

	err := client.rpcClient.Call(method, args, &reply)
	if err != nil {
		return newErrMethodFailed(method, err)
	}

	if reply.ErrorMsg != "" {
		return errors.New(reply.ErrorMsg)
	}

	return nil
}

func (client *Client) ListPackages() ([]slang.PackageMeta, error) {
	const method = "Packages.List"
	args := &models.ListPackagesArgs{}

	var reply models.ListPackagesReply

	err := client.rpcClient.Call(method, args, &reply)
	if err != nil {
		return []slang.PackageMeta{}, newErrMethodFailed(method, err)
	}

	return reply.Metas, nil
}

func (client *Client) GetPackageArchive(meta slang.PackageMeta) (slang.PackageArchive, bool, error) {
	const method = "Packages.GetArchive"

	args := &models.GetArchiveArgs{Meta: meta}

	var reply models.GetArchiveReply

	err := client.rpcClient.Call(method, args, &reply)
	if err != nil {
		return nil, false, newErrMethodFailed(method, err)
	}

	return reply.Archive, reply.Found, nil
}

func (client *Client) RemovePackage(meta slang.PackageMeta) (bool, error) {
	const method = "Packages.Remove"

	args := &models.RemovePackageArgs{Meta: meta}

	var reply models.RemovePackageReply

	err := client.rpcClient.Call(method, args, &reply)
	if err != nil {
		return false, newErrMethodFailed(method, err)
	}

	return reply.Removed, nil
}
