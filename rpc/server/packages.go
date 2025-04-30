package server

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/rpc/models"
)

type Packages struct {
	VM slang.VirtualMachine
}

func (api *Packages) Add(args *models.AddPackageArgs, reply *models.AddPackageReply) error {
	if err := api.VM.AddPackage(args.Archive); err != nil {
		reply.ErrorMsg = err.Error()
	}

	return nil
}

func (api *Packages) List(args *models.ListPackagesArgs, reply *models.ListPackagesReply) error {
	reply.Metas = api.VM.ListPackages()

	return nil
}

func (api *Packages) GetPackageArchive(args *models.GetArchiveArgs, reply *models.GetArchiveReply) error {
	pkg, found := api.VM.GetPackageArchive(args.Meta)
	if !found {
		reply.Archive = nil
		reply.Found = false
	}

	reply.Archive = pkg
	reply.Found = true

	return nil
}

func (api *Packages) Remove(args *models.RemovePackageArgs, reply *models.RemovePackageReply) error {
	reply.Removed = api.VM.RemovePackage(args.Meta)

	return nil
}
