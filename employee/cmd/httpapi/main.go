package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"

	"github.com/BayMaxx2001/manager-employee/employee/internal/app/httpapi"
	"github.com/BayMaxx2001/manager-employee/employee/internal/app/terminal"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Employee micro service"
	app.Usage = "Employee micro service"
	app.Copyright = "Copyright © 2021 CyRadar. All Rights Reserved."
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "env",
			Aliases: []string{"e"},
			Value:   "../../configs/.env",
			Usage:   "set path to environment file",
		},
		&cli.StringFlag{
			Name:    "env_prefix",
			Aliases: []string{"p"},
			Value:   "employee",
			Usage:   "set path to environment prefix",
		},
	}
	app.Commands = []*cli.Command{
		httpapi.CommandRunServer(),
		terminal.CommandRunTerminal(),
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	endSignal := make(chan os.Signal, 1)
	signal.Notify(endSignal, syscall.SIGINT, syscall.SIGTERM)

	errChan := make(chan error, 1)

	go func(ctx context.Context, errChan chan error) {
		err := app.RunContext(ctx, os.Args)

		errChan <- err

	}(ctx, errChan)

	select {
	case sign := <-endSignal:
		log.Println("shutting down. reason:", sign)
		return
	case err := <-errChan:
		log.Println("encountered error:", err)
		return
	}
}
