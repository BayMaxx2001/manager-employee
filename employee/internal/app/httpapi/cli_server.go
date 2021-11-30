package httpapi

import (
	"github.com/urfave/cli/v2"

	"github.com/BayMaxx2001/manager-employee/employee/internal/config"
	"github.com/BayMaxx2001/manager-employee/employee/internal/persistence"
)

func CommandRunServer() *cli.Command {
	cli := cli.Command{
		Name:   "serve",
		Usage:  "Start the core server",
		Action: ServeActionCommand,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "address",
				Aliases: []string{"addr"},
				Value:   "localhost:8282",
				Usage:   "specify which address to serve on",
			},
		},
	}

	return &cli
}

func ServeActionCommand(c *cli.Context) error {
	if err := config.LoadEnvFromFile(c.String("env_prefix"), c.String("env")); err != nil {
		return err
	}

	if err := persistence.LoadEmployeesMongoRedisRepository(); err != nil {
		return err
	}

	return Serve(c.Context, c.String("addr"))
}
