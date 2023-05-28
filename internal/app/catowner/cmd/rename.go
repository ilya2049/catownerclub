package cmd

import (
	"catownerclub/internal/domain/catowner"
	"context"
)

type Rename struct {
	catOwnerRepository catowner.Repository
}

func NewRename(catOwnerRepository catowner.Repository) *Rename {
	return &Rename{
		catOwnerRepository: catOwnerRepository,
	}
}

func (cmd *Rename) Execute(
	ctx context.Context,
	catOwnerID int,
	newName string,
) error {
	catOwner, err := cmd.catOwnerRepository.
		Get(catOwnerID).
		ForUpdate().
		Load(ctx)

	if err != nil {
		return err
	}

	catOwner.Rename(newName)

	if err := cmd.catOwnerRepository.Update(ctx, catOwner); err != nil {
		return err
	}

	return nil
}
