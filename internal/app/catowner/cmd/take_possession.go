package cmd

import (
	"catownerclub/internal/domain/catowner"

	"context"
)

type TakePossession struct {
	catOwnerRepository catowner.Repository
}

func NewTakePossession(catOwnerRepository catowner.Repository) *TakePossession {
	return &TakePossession{
		catOwnerRepository: catOwnerRepository,
	}
}

func (cmd *TakePossession) Execute(
	ctx context.Context,
	catOwnerID int,
	catID int,
	catName string,
) error {
	catOwner, err := cmd.catOwnerRepository.
		Get(catOwnerID).
		ForUpdate().
		Load(ctx)

	if err != nil {
		return err
	}

	newCat := catowner.NewCat(catID, catName)
	catOwner.TakePossessionOfCat(newCat)

	if err := cmd.catOwnerRepository.Update(ctx, catOwner); err != nil {
		return err
	}

	return nil
}
