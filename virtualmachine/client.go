package virtualmachine

import (
	"fmt"
	"net/rpc"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/rpc/client"
)

func MakeClient(tcpAddr string) (slang.VirtualMachine, error) {
	rpcClient, err := rpc.Dial("tcp", tcpAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	return client.New(rpcClient), nil
}
