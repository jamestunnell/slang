package virtualmachine

import (
	"fmt"
	"net/rpc"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/rpc/client"
)

type Client interface {
	GetInfo() (slang.VMInfo, error)

	ListPackages() ([]slang.PackageMeta, error)
	GetPackage(slang.PackageMeta) (slang.PackageArchive, bool, error)
	AddPackage(slang.PackageArchive) error
	RemovePackage(slang.PackageMeta) (bool, error)

	EvaluateExpr(slang.Expression) (slang.Object, error)
}

func MakeClient(tcpAddr string) (Client, error) {
	rpcClient, err := rpc.Dial("tcp", tcpAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	return client.New(rpcClient), nil
}
