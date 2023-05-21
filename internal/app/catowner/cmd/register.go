package cmd

import (
	"context"

	"catownerclub/internal/domain/catowner"
)

type Register struct {
	catOwnerRepository catowner.Repository
}

func NewRegister(catOwnerRepository catowner.Repository) *Register {
	return &Register{
		catOwnerRepository: catOwnerRepository,
	}
}

func (cmd *Register) Execute(
	ctx context.Context,
	catOwnerID int,
	catOwnerName string,
) error {
	newCatOwner := catowner.New(catOwnerID, catOwnerName)

	if err := cmd.catOwnerRepository.Add(ctx, newCatOwner); err != nil {
		return err
	}

	return nil
}
