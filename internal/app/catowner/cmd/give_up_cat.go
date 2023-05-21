package cmd

import (
	"context"

	"catownerclub/internal/domain/catowner"
)

type GiveUpCat struct {
	catOwnerRepository catowner.Repository
}

func NewGiveUpCat(catOwnerRepository catowner.Repository) *GiveUpCat {
	return &GiveUpCat{
		catOwnerRepository: catOwnerRepository,
	}
}

func (cmd *GiveUpCat) Execute(
	ctx context.Context,
	catOwnerID int,
	catID int,
) error {
	catOwner, err := cmd.catOwnerRepository.
		Get(catOwnerID).
		WithCat(catID).
		Load(ctx)
	if err != nil {
		return err
	}

	catOwner.GiveUpCat(catID)

	if err := cmd.catOwnerRepository.Update(ctx, catOwner); err != nil {
		return err
	}

	return nil
}
