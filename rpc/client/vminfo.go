package client

import (
	"fmt"
	"net/rpc"

	"github.com/jamestunnell/slang/rpc/models"
)

type VMInfo interface {
	GetID() (string, error)
}

type vmInfoClient struct {
	rpcClient *rpc.Client
}

func NewVMInfo(serverAddress string) (VMInfo, error) {
	rpcc, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to dial server: %w", err)
	}

	client := &vmInfoClient{
		rpcClient: rpcc,
	}

	return client, nil
}

func (client *vmInfoClient) GetID() (string, error) {
	args := &models.GetVMIDArgs{}

	var reply models.GetVMIDReply

	err := client.rpcClient.Call("VMInfo.GetID", args, &reply)
	if err != nil {
		return "", fmt.Errorf("get-vminfo-id error: %w", err)
	}

	return reply.ID, nil
}
