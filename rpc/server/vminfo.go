package server

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/rpc/models"
)

type VMInfo struct {
	VM slang.VirtualMachine
}

func (api *VMInfo) GetInfo(_ *models.GetVMInfoArgs, reply *slang.VMInfo) error {
	info := api.VM.GetInfo()

	reply.ID = info.ID
	reply.Name = info.Name

	return nil
}
