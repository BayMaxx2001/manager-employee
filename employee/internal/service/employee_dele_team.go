package service

import (
	"context"
	"encoding/json"

	"github.com/BayMaxx2001/manager-employee/employee/internal/persistence"
	"github.com/asaskevich/govalidator"
)

type EmployeeDeleteTeamCommand struct {
	EmployeeId string `json:"eid"`
	TeamId     string `json:"tid"`
}

var employeeDeleteTeam EmployeeDeleteTeamCommand

// publish
func (e *EmployeeDeleteTeamCommand) Name() string {
	return "employee-team"
}

func (e *EmployeeDeleteTeamCommand) JSON() []byte {
	b, err := json.Marshal(employeeDeleteTeam)
	if err != nil {
		return nil
	}
	return b
}

func (c EmployeeDeleteTeamCommand) Valid(ctx context.Context) error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	return nil
}
func DeleteEmployeeToTeam(ctx context.Context, command EmployeeDeleteTeamCommand) error {
	employeeDeleteTeam = EmployeeDeleteTeamCommand{
		EmployeeId: command.EmployeeId,
		TeamId:     command.TeamId,
	}
	if err := employeeDeleteTeam.Valid(ctx); err != nil {
		return err
	}

	employee, err := persistence.Employees().FindByUID(ctx, employeeDeleteTeam.EmployeeId)
	if err != nil {
		return err
	}
	err = persistence.Employees().DeleteEmployeeToTeam(ctx, employee, employeeDeleteTeam.TeamId)

	return err
}
