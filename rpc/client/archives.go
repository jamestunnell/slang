package client

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/rpc/models"
)

func (client *Client) AddPackage(archive slang.PackageArchive) error {
	const method = "Archives.Add"

	args := &models.AddArchiveArgs{Archive: archive}

	var reply models.AddArchiveReply

	err := client.rpcClient.Call(method, args, &reply)
	if err != nil {
		return newErrMethodFailed(method, err)
	}

	return nil
}

func (client *Client) ListPackages() ([]slang.PackageMeta, error) {
	const method = "Archives.List"
	args := &models.ListArchivesArgs{}

	var reply models.ListArchivesReply

	err := client.rpcClient.Call(method, args, &reply)
	if err != nil {
		return []slang.PackageMeta{}, newErrMethodFailed(method, err)
	}

	return reply.Metas, nil
}

func (client *Client) GetPackage(meta slang.PackageMeta) (slang.PackageArchive, bool, error) {
	const method = "Archives.Get"

	args := &models.GetArchiveArgs{Meta: meta}

	var reply models.GetArchiveReply

	err := client.rpcClient.Call(method, args, &reply)
	if err != nil {
		return nil, false, newErrMethodFailed(method, err)
	}

	return reply.Archive, reply.Found, nil
}

func (client *Client) RemovePackage(meta slang.PackageMeta) (bool, error) {
	const method = "Archives.Remove"

	args := &models.RemoveArchiveArgs{Meta: meta}

	var reply models.RemoveArchiveReply

	err := client.rpcClient.Call(method, args, &reply)
	if err != nil {
		return false, newErrMethodFailed(method, err)
	}

	return reply.Removed, nil
}
