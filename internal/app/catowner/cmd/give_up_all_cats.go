package cmd

import (
	"context"
	"fmt"

	"catownerclub/internal/domain/catowner"
)

type GiveUpAllCats struct {
	catOwnerRepository catowner.Repository
}

func NewGiveUpAllCats(catOwnerRepository catowner.Repository) *GiveUpAllCats {
	return &GiveUpAllCats{
		catOwnerRepository: catOwnerRepository,
	}
}

func (cmd *GiveUpAllCats) Execute(
	ctx context.Context,
	catOwnerID int,
) error {
	catOwner, err := cmd.catOwnerRepository.
		Get(catOwnerID).
		ForUpdate().
		WithCats().
		Load(ctx)
	if err != nil {
		return fmt.Errorf("failed to get cat owner: %w", err)
	}

	catOwner.GiveUpAllCats()

	if err := cmd.catOwnerRepository.Update(ctx, catOwner); err != nil {
		return fmt.Errorf("failed to update cat owner: %w", err)
	}

	return nil
}
