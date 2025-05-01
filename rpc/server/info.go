package server

import (
	"github.com/google/uuid"
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/rpc/models"
)

type Info struct {
	VM slang.VirtualMachine
}

func (api *Info) IsRunning(_ *models.Empty, reply *bool) error {
	*reply = api.VM.IsRunning()

	return nil
}

func (api *Info) GetName(_ *models.Empty, reply *string) error {
	*reply = api.VM.GetName()

	return nil
}

func (api *Info) GetID(_ *models.Empty, reply *uuid.UUID) error {
	*reply = api.VM.GetID()

	return nil
}
