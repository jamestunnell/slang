package server

import (
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/archives"
	"github.com/jamestunnell/slang/rpc/models"
)

type Archives struct {
	VM slang.VirtualMachine
}

func (api *Archives) AddTarGz(args *models.AddTarGzArgs, reply *models.AddTarGzReply) error {
	api.VM.AddPackage(args.Archive)

	return nil
}

func (api *Archives) List(args *models.ListArchivesArgs, reply *models.ListArchivesReply) error {
	reply.Metas = api.VM.ListPackages()

	return nil
}

func (api *Archives) GetTarGz(args *models.GetArchiveArgs, reply *models.GetArchiveReply) error {
	pkg, found := api.VM.GetPackage(args.Meta)
	if !found {
		reply.Archive = nil
		reply.Found = false
	}

	tgz, ok := pkg.(*archives.TarGz)
	if !ok {
		return fmt.Errorf("package format '%s' is not TarGz", pkg.GetDataFormat())
	}

	reply.Archive = tgz
	reply.Found = true

	return nil
}

func (api *Archives) Remove(args *models.RemoveArchiveArgs, reply *models.RemoveArchiveReply) error {
	reply.Removed = api.VM.RemovePackage(args.Meta)

	return nil
}
