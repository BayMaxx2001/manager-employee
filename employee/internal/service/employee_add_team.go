package service

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/BayMaxx2001/manager-employee/employee/internal/persistence"
	"github.com/asaskevich/govalidator"
)

type EmployeeAddTeamCommand struct {
	EmployeeId string `json:"eid"`
	TeamId     string `json:"tid"`
}

var employeeAddTeam EmployeeAddTeamCommand

// publish
func (e *EmployeeAddTeamCommand) Name() string {
	return "employee-team"
}

func (e *EmployeeAddTeamCommand) JSON() []byte {
	b, err := json.Marshal(employeeAddTeam)
	if err != nil {
		return nil
	}
	return b
}

func (c EmployeeAddTeamCommand) Valid(ctx context.Context) error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	employee, err := FindEmployeeByUID(ctx, FindEmployeeByUIDCommand(employeeAddTeam.EmployeeId))
	if err != nil {
		return err
	}

	for _, lsTeams := range employee.ListTeams {
		if strings.Compare(lsTeams, employeeAddTeam.TeamId) == 0 {
			return errors.New("employee exits in team")
		}
	}

	return nil
}

func AddEmployeeToTeam(ctx context.Context, command EmployeeAddTeamCommand) (EmployeeAddTeamCommand, error) {
	employeeAddTeam = EmployeeAddTeamCommand{
		EmployeeId: command.EmployeeId,
		TeamId:     command.TeamId,
	}

	if err := employeeAddTeam.Valid(ctx); err != nil {
		return employeeAddTeam, err
	}

	employee, err := persistence.Employees().FindByUID(ctx, employeeAddTeam.EmployeeId)
	if err != nil {
		return employeeAddTeam, err
	}

	err = persistence.Employees().AddEmployeeToTeam(ctx, employee, employeeAddTeam.TeamId)
	return employeeAddTeam, err
}
