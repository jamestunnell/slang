package server

import (
	"github.com/jamestunnell/slang/archives"
	"github.com/jamestunnell/slang/rpc/models"
)

type Archives struct {
	Repo *archives.Repo
}

func NewArchives() *Archives {
	return &Archives{Repo: archives.NewRepo()}
}

func (api *Archives) Add(args *models.AddArchiveArgs, reply *models.AddArchiveReply) error {
	api.Repo.Add(args.Archive)

	return nil
}

func (api *Archives) List(args *models.ListArchivesArgs, reply *models.ListArchivesReply) error {
	reply.Metas = api.Repo.GetAllMeta()

	return nil
}

func (api *Archives) Get(args *models.GetArchiveArgs, reply *models.GetArchiveReply) error {
	reply.Archive, reply.Found = api.Repo.Get(args.Meta)

	return nil
}

func (api *Archives) Remove(args *models.RemoveArchiveArgs, reply *models.RemoveArchiveReply) error {
	reply.Removed = api.Repo.Remove(args.Meta)

	return nil
}
