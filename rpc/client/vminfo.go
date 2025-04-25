package client

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/rpc/models"
)

func (client *Client) GetInfo() (slang.VMInfo, error) {
	const method = "VMInfo.GetInfo"

	args := &models.GetVMInfoArgs{}

	var reply slang.VMInfo

	err := client.rpcClient.Call(method, args, &reply)
	if err != nil {
		return slang.VMInfo{}, newErrMethodFailed(method, err)
	}

	return reply, nil
}
