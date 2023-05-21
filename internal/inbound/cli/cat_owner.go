package cli

import (
	"catownerclub/internal/app/catowner/cmd"

	"github.com/urfave/cli/v2"
)

type CatOwnerCommands struct {
	giveUpAllCatsCommand  *cmd.GiveUpAllCats
	giveUpCatCommand      *cmd.GiveUpCat
	leaveCommand          *cmd.Leave
	registerCommand       *cmd.Register
	renameCatCommand      *cmd.RenameCat
	renameCommand         *cmd.Rename
	takePossessionCommand *cmd.TakePossession
}

func NewCatOwnerCommands(
	giveUpAllCatsCommand *cmd.GiveUpAllCats,
	giveUpCatCommand *cmd.GiveUpCat,
	leaveCommand *cmd.Leave,
	registerCommand *cmd.Register,
	renameCatCommand *cmd.RenameCat,
	renameCommand *cmd.Rename,
	takePossessionCommand *cmd.TakePossession,
) *CatOwnerCommands {
	return &CatOwnerCommands{
		giveUpAllCatsCommand:  giveUpAllCatsCommand,
		giveUpCatCommand:      giveUpCatCommand,
		leaveCommand:          leaveCommand,
		registerCommand:       registerCommand,
		renameCatCommand:      renameCatCommand,
		renameCommand:         renameCommand,
		takePossessionCommand: takePossessionCommand,
	}
}

func (cmds *CatOwnerCommands) Register() *cli.Command {
	return &cli.Command{
		Name:  "register",
		Usage: "Register a new cat owner.",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "cat-owner-id", Aliases: []string{"id"}},
			&cli.StringFlag{Name: "cat-owner-name", Aliases: []string{"name"}},
		},
		Action: func(cCtx *cli.Context) error {
			if err := cmds.registerCommand.Execute(cCtx.Context,
				cCtx.Int("cat-owner-id"),
				cCtx.String("cat-owner-name"),
			); err != nil {
				return err
			}

			return nil
		},
	}
}

func (cmds *CatOwnerCommands) TakePossession() *cli.Command {
	return &cli.Command{
		Name:  "take-possession",
		Usage: "Take possession of a cat.",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "cat-owner-id"},
			&cli.IntFlag{Name: "cat-id"},
			&cli.StringFlag{Name: "cat-name"},
		},
		Action: func(cCtx *cli.Context) error {
			if err := cmds.takePossessionCommand.Execute(cCtx.Context,
				cCtx.Int("cat-owner-id"),
				cCtx.Int("cat-id"),
				cCtx.String("cat-name"),
			); err != nil {
				return err
			}

			return nil
		},
	}
}

func (cmds *CatOwnerCommands) GiveUpAllCats() *cli.Command {
	return &cli.Command{
		Name:  "give-up-all-cats",
		Usage: "Give up all owner cats.",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "cat-owner-id", Aliases: []string{"id"}},
		},
		Action: func(cCtx *cli.Context) error {
			if err := cmds.giveUpAllCatsCommand.Execute(cCtx.Context,
				cCtx.Int("cat-owner-id"),
			); err != nil {
				return err
			}

			return nil
		},
	}
}

func (cmds *CatOwnerCommands) GiveUpCat() *cli.Command {
	return &cli.Command{
		Name:  "give-up-cat",
		Usage: "Give up a cat.",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "cat-owner-id"},
			&cli.IntFlag{Name: "cat-id"},
		},
		Action: func(cCtx *cli.Context) error {
			if err := cmds.giveUpCatCommand.Execute(cCtx.Context,
				cCtx.Int("cat-owner-id"),
				cCtx.Int("cat-id"),
			); err != nil {
				return err
			}

			return nil
		},
	}
}

func (cmds *CatOwnerCommands) Leave() *cli.Command {
	return &cli.Command{
		Name:  "leave",
		Usage: "Leave the club.",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "cat-owner-id", Aliases: []string{"id"}},
		},
		Action: func(cCtx *cli.Context) error {
			if err := cmds.leaveCommand.Execute(cCtx.Context,
				cCtx.Int("cat-owner-id"),
			); err != nil {
				return err
			}

			return nil
		},
	}
}

func (cmds *CatOwnerCommands) Rename() *cli.Command {
	return &cli.Command{
		Name:  "rename",
		Usage: "Change owner's name.",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "cat-owner-id", Aliases: []string{"id"}},
			&cli.StringFlag{Name: "cat-owner-new-name", Aliases: []string{"name"}},
		},
		Action: func(cCtx *cli.Context) error {
			if err := cmds.renameCommand.Execute(cCtx.Context,
				cCtx.Int("cat-owner-id"),
				cCtx.String("cat-owner-new-name"),
			); err != nil {
				return err
			}

			return nil
		},
	}
}

func (cmds *CatOwnerCommands) RenameCat() *cli.Command {
	return &cli.Command{
		Name:  "rename-cat",
		Usage: "Change cat's name.",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "cat-owner-id"},
			&cli.IntFlag{Name: "cat-id"},
			&cli.StringFlag{Name: "cat-new-name"},
		},
		Action: func(cCtx *cli.Context) error {
			if err := cmds.renameCatCommand.Execute(cCtx.Context,
				cCtx.Int("cat-owner-id"),
				cCtx.Int("cat-id"),
				cCtx.String("cat-new-name"),
			); err != nil {
				return err
			}

			return nil
		},
	}
}
