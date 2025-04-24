package server

import (
	"github.com/jamestunnell/slang/rpc/models"
)

type VMInfo struct {
	id string
}

func NewVMInfo(id string) *VMInfo {
	return &VMInfo{id: id}
}

func (api *VMInfo) GetID(_ *models.GetVMIDArgs, reply *models.GetVMIDReply) error {
	reply.ID = api.id

	return nil
}
