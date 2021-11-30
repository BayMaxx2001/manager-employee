package service

import (
	"context"
	"encoding/json"

	"github.com/BayMaxx2001/manager-employee/team/internal/persistence"
	"github.com/asaskevich/govalidator"
)

type TeamDeleteEmployeeCommand struct {
	EmployeeId string `json:"eid"`
	TeamId     string `json:"tid"`
}

var teamDeleteEmployee TeamDeleteEmployeeCommand

// publish
func (e *TeamDeleteEmployeeCommand) Name() string {
	return "employee-team"
}

func (e *TeamDeleteEmployeeCommand) JSON() []byte {
	b, err := json.Marshal(teamDeleteEmployee)
	if err != nil {
		return nil
	}
	return b
}

func (c TeamDeleteEmployeeCommand) Valid(ctx context.Context) error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	return nil
}
func DeleteTeamToEmployee(ctx context.Context, command TeamDeleteEmployeeCommand) error {
	teamDeleteEmployee = TeamDeleteEmployeeCommand{
		EmployeeId: command.EmployeeId,
		TeamId:     command.TeamId,
	}

	if err := teamDeleteEmployee.Valid(ctx); err != nil {
		return err
	}

	team, err := persistence.Teams().FindByUID(ctx, teamDeleteEmployee.TeamId)
	if err != nil {
		return err
	}

	err = persistence.Teams().DeleteTeamToEmployee(ctx, team, teamDeleteEmployee.EmployeeId)
	return err
}
