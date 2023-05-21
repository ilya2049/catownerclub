package main

import (
	"catownerclub/internal/app/catowner/cmd"
	catownercli "catownerclub/internal/inbound/cli"
	"catownerclub/internal/outbound/pg"
	"catownerclub/internal/outbound/pg/catowner"

	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/urfave/cli/v2"
)

// export DATABASE_URL=postgres://postgres:password@localhost:5432/postgres

func main() {
	connectionPoolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)

		return
	}

	connectionPoolConfig.ConnConfig.Tracer = pg.NewTraceLog()

	connectionPool, err := pgxpool.NewWithConfig(context.Background(), connectionPoolConfig)
	if err != nil {
		log.Println(err)

		return
	}

	catOwnerRepository := catowner.NewRepository(connectionPool)

	cliCommands := catownercli.NewCatOwnerCommands(
		cmd.NewGiveUpAllCats(catOwnerRepository),
		cmd.NewGiveUpCat(catOwnerRepository),
		cmd.NewLeave(catOwnerRepository),
		cmd.NewRegister(catOwnerRepository),
		cmd.NewRenameCat(catOwnerRepository),
		cmd.NewRename(catOwnerRepository),
		cmd.NewTakePossession(catOwnerRepository),
	)

	app := &cli.App{
		Commands: []*cli.Command{
			cliCommands.GiveUpAllCats(),
			cliCommands.GiveUpCat(),
			cliCommands.Leave(),
			cliCommands.Register(),
			cliCommands.RenameCat(),
			cliCommands.Rename(),
			cliCommands.TakePossession(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Println(err)

		return
	}
}
