package terminal

import (
	"fmt"

	"github.com/BayMaxx2001/manager-employee/employee/internal/model"
	"github.com/BayMaxx2001/manager-employee/employee/internal/service"
	"github.com/urfave/cli/v2"
)

type CommandTerminal interface {
	CommandAddEmployee() *cli.Command
	CommandUpdateEmployee() *cli.Command
	CommandDeleteEmployeeByUUID() *cli.Command
	CommandGetAllEmployees() *cli.Command
	CommandFindEmployee() *cli.Command
}

func CommandRunTerminal() *cli.Command {
	cli := cli.Command{
		Name:  "terminal",
		Usage: "Start the core in terminal",
		Subcommands: []*cli.Command{
			CommandEmployee(),
		},
	}

	return &cli
}

func CommandEmployee() *cli.Command {
	cli := cli.Command{
		Name:  "employee",
		Usage: "CRUD for employee ",
		Subcommands: []*cli.Command{
			CommandAddEmployee(),
			CommandUpdateEmployee(),
			CommandDeleteEmployeeByUUID(),
			CommandGetAllEmployees(),
			CommandFindEmployee(),
		},
	}

	return &cli
}

func CommandFindEmployee() *cli.Command {
	var uuid string
	cli := cli.Command{
		Name:    "Find",
		Usage:   "Find employee by uuid",
		Aliases: []string{"f"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "uuid",
				Usage:       "Id to search employee ",
				Destination: &uuid,
			},
		},
		Action: func(c *cli.Context) error {
			employee, err := service.FindEmployeeByUID(c.Context, service.FindEmployeeByUIDCommand(uuid))
			fmt.Printf("%s %s %d %s\n", employee.UID, employee.Name, employee.Gender, employee.DobFormat("2006-02-01"))
			if err != nil {
				return err
			}
			return nil
		},
	}

	return &cli
}

func CommandGetAllEmployees() *cli.Command {
	cli := cli.Command{
		Name:  "all",
		Usage: "get all employee",
		Action: func(c *cli.Context) error {
			employees, err := service.GetAllEmployees(c.Context)
			if err != nil {
				return err
			}
			for _, employee := range employees {
				fmt.Printf("%s %s %d %s\n", employee.UID, employee.Name, employee.Gender, employee.DobFormat("2006-02-01"))
			}
			return nil
		},
	}

	return &cli
}

func CommandDeleteEmployeeByUUID() *cli.Command {
	var id string
	cli := cli.Command{
		Name:    "delete",
		Usage:   "delete employee by id",
		Aliases: []string{"del"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "id",
				Usage:       "Id to remove",
				Destination: &id,
			},
		},
		Action: func(c *cli.Context) error {
			err := service.DeleteEmployeeByUID(c.Context, service.DeleteEmployeeByUIDCommand(id))
			if err != nil {
				return err
			}
			return nil
		},
	}

	return &cli
}

func CommandAddEmployee() *cli.Command {
	var employee model.Employee
	cli := cli.Command{
		Name:  "add",
		Usage: "add new employee",
		Flags: *FlagSaveEmployee(&employee),
		Action: func(c *cli.Context) error {
			_, err := service.AddEmployee(c.Context, service.NewAddEmployeeCommand(employee))
			return err
		},
	}
	return &cli
}

func CommandUpdateEmployee() *cli.Command {
	var employee model.Employee
	cli := cli.Command{
		Name:  "update",
		Usage: "update employee by id ",
		Flags: *FlagSaveEmployee(&employee),
		Action: func(c *cli.Context) error {
			err := service.UpdateEmployeeById(c.Context, service.NewUpdateEmployeeCommand(employee))
			return err
		},
	}

	return &cli
}

func FlagSaveEmployee(employee *model.Employee) *[]cli.Flag {
	flag := []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Usage:       "new name employee",
			Aliases:     []string{"n"},
			Destination: &employee.Name,
		},
		&cli.IntFlag{
			Name:        "gender",
			Usage:       "new gender employee",
			Aliases:     []string{"g"},
			Destination: &employee.Gender,
		},
		&cli.StringFlag{
			Name:        "dob",
			Usage:       "new dob employee",
			Aliases:     []string{"d"},
			Destination: &employee.DOB,
		},
	}

	return &flag
}
