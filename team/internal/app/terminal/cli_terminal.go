package terminal

import (
	"fmt"

	"github.com/BayMaxx2001/manager-employee/team/internal/model"
	"github.com/BayMaxx2001/manager-employee/team/internal/service"
	"github.com/urfave/cli/v2"
)

type CommandTerminal interface {
	CommandAddTeam() *cli.Command
	CommandUpdateTeam() *cli.Command
	CommandDeleteTeamByUUID() *cli.Command
	CommandGetAllTeam() *cli.Command
	CommandFindTeam() *cli.Command
}

func CommandRunTerminal() *cli.Command {
	cli := cli.Command{
		Name:  "terminal",
		Usage: "Start the core in terminal",
		Subcommands: []*cli.Command{
			CommandTeam(),
		},
	}
	return &cli
}

func CommandTeam() *cli.Command {
	cli := cli.Command{
		Name:  "team",
		Usage: "CRUD for Team ",
		Subcommands: []*cli.Command{
			CommandAddTeam(),
			CommandUpdateTeam(),
			CommandDeleteTeamByUUID(),
			CommandGetAllTeam(),
			CommandFindTeam(),
		},
	}
	return &cli
}

func CommandFindTeam() *cli.Command {
	var uuid string
	cli := cli.Command{
		Name:    "find",
		Usage:   "Find Team by uuid",
		Aliases: []string{"f"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "uuid",
				Usage:       "Id to search Team ",
				Destination: &uuid,
			},
		},
		Action: func(c *cli.Context) error {
			team, err := service.FindTeamByUID(c.Context, service.FindTeamByUIDCommand(uuid))
			//logger
			fmt.Printf("%s %s\n", team.UID, team.Name)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return &cli
}

func CommandGetAllTeam() *cli.Command {
	cli := cli.Command{
		Name:  "all",
		Usage: "get all team",
		Action: func(c *cli.Context) error {
			teams, err := service.GetAllTeams(c.Context)
			if err != nil {
				return err
			}
			for _, team := range teams {
				fmt.Printf("%s %s\n", team.UID, team.Name)
			}
			return nil
		},
	}
	return &cli
}

func CommandDeleteTeamByUUID() *cli.Command {
	var id string
	cli := cli.Command{
		Name:    "delete",
		Usage:   "delete team by id",
		Aliases: []string{"del"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "id",
				Usage:       "Id to remove",
				Destination: &id,
			},
		},
		Action: func(c *cli.Context) error {
			err := service.DeleteTeamByUID(c.Context, service.DeleteTeamByUIDCommand(id))
			if err != nil {
				return err
			}
			return nil
		},
	}
	return &cli
}

func CommandAddTeam() *cli.Command {
	var team model.Team
	cli := cli.Command{
		Name:  "add",
		Usage: "add new team",
		Flags: *FlagSaveTeam(&team),
		Action: func(c *cli.Context) error {
			_, err := service.AddTeam(c.Context, service.NewAddTeamCommand(team))
			return err
		},
	}
	return &cli
}

func CommandUpdateTeam() *cli.Command {
	var team model.Team
	cli := cli.Command{
		Name:  "update",
		Usage: "update team by id ",
		Flags: *FlagSaveTeam(&team),
		Action: func(c *cli.Context) error {
			err := service.UpdateTeamById(c.Context, service.NewUpdateTeamCommand(team))
			return err
		},
	}
	return &cli
}

func FlagSaveTeam(team *model.Team) *[]cli.Flag {
	flag := []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Usage:       "new name team",
			Aliases:     []string{"n"},
			Destination: &team.Name,
		},
	}

	return &flag
}
