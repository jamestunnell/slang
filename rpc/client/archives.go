package client

import (
	"fmt"
	"net/rpc"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/archives"
	"github.com/jamestunnell/slang/rpc/models"
)

type Archives interface {
	Add(*archives.TarGz) error
	List() ([]slang.PackageMeta, error)
	Get(slang.PackageMeta) (slang.PackageArchive, bool, error)
	Remove(slang.PackageMeta) (bool, error)
}

type archivesClient struct {
	rpcClient *rpc.Client
}

func NewArchives(serverAddress string) (Archives, error) {
	rpcc, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to dial server: %w", err)
	}

	client := &archivesClient{
		rpcClient: rpcc,
	}

	return client, nil
}

func (client *archivesClient) Add(archive *archives.TarGz) error {
	args := &models.AddArchiveArgs{Archive: archive}

	var reply models.AddArchiveReply

	err := client.rpcClient.Call("Archives.Add", args, &reply)
	if err != nil {
		return fmt.Errorf("add-archive error: %w", err)
	}

	return nil
}

func (client *archivesClient) List() ([]slang.PackageMeta, error) {
	args := &models.ListArchivesArgs{}

	var reply models.ListArchivesReply

	err := client.rpcClient.Call("Archives.List", args, &reply)
	if err != nil {
		return []slang.PackageMeta{}, fmt.Errorf("list-archives error: %w", err)
	}

	return reply.Metas, nil
}

func (client *archivesClient) Get(meta slang.PackageMeta) (slang.PackageArchive, bool, error) {
	args := &models.GetArchiveArgs{Meta: meta}

	var reply models.GetArchiveReply

	err := client.rpcClient.Call("Archives.Get", args, &reply)
	if err != nil {
		return nil, false, fmt.Errorf("get-archive error: %w", err)
	}

	return reply.Archive, reply.Found, nil
}

func (client *archivesClient) Remove(meta slang.PackageMeta) (bool, error) {
	args := &models.RemoveArchiveArgs{Meta: meta}

	var reply models.RemoveArchiveReply

	err := client.rpcClient.Call("Archives.Remove", args, &reply)
	if err != nil {
		return false, fmt.Errorf("get-archive error: %w", err)
	}

	return reply.Removed, nil
}
