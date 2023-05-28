package cmd

import (
	"context"

	"catownerclub/internal/domain/catowner"
)

type RenameCat struct {
	catOwnerRepository catowner.Repository
}

func NewRenameCat(catOwnerRepository catowner.Repository) *RenameCat {
	return &RenameCat{
		catOwnerRepository: catOwnerRepository,
	}
}

func (cmd *RenameCat) Execute(
	ctx context.Context,
	catOwnerID int,
	catID int,
	newCatName string,
) error {
	catOwner, err := cmd.catOwnerRepository.
		Get(catOwnerID).
		ForUpdate().
		WithCat(catID).
		Load(ctx)

	if err != nil {
		return err
	}

	catOwner.RenameCat(catID, newCatName)

	if err := cmd.catOwnerRepository.Update(ctx, catOwner); err != nil {
		return err
	}

	return nil
}
