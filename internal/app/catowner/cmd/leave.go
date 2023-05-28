package cmd

import (
	"context"

	"catownerclub/internal/domain/catowner"
)

type Leave struct {
	catOwnerRepository catowner.Repository
}

func NewLeave(catOwnerRepository catowner.Repository) *Leave {
	return &Leave{
		catOwnerRepository: catOwnerRepository,
	}
}

func (cmd *Leave) Execute(
	ctx context.Context,
	catOwnerID int,
) error {
	catOwner, err := cmd.catOwnerRepository.Get(catOwnerID).
		ForUpdate().
		WithCats().
		Load(ctx)
	if err != nil {
		return err
	}

	catOwner.GiveUpAllCats()

	if err := cmd.catOwnerRepository.Delete(ctx, catOwner); err != nil {
		return err
	}

	return nil
}
