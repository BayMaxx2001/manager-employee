package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/BayMaxx2001/manager-employee/team/internal/persistence"
	"github.com/asaskevich/govalidator"
)

type TeamAddEmployeeCommand struct {
	EmployeeId string `json:"eid"`
	TeamId     string `json:"tid"`
}

var teamAddEmployee TeamAddEmployeeCommand

func (c *TeamAddEmployeeCommand) Valid(ctx context.Context) error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	team, err := FindTeamByUID(ctx, FindTeamByUIDCommand(teamAddEmployee.TeamId))
	if err != nil {
		return err
	}

	for _, lsEms := range team.ListEmployees {
		if strings.Compare(lsEms, teamAddEmployee.EmployeeId) == 0 {
			return errors.New("employee exits in team")
		}
	}

	return nil
}
func AddTeamToEmployee(ctx context.Context, command TeamAddEmployeeCommand) (TeamAddEmployeeCommand, error) {
	teamAddEmployee = TeamAddEmployeeCommand{
		EmployeeId: command.EmployeeId,
		TeamId:     command.TeamId,
	}

	if err := teamAddEmployee.Valid(ctx); err != nil {
		fmt.Println(err)
		return teamAddEmployee, err
	}

	team, err := persistence.Teams().FindByUID(ctx, teamAddEmployee.TeamId)
	if err != nil {
		return teamAddEmployee, err
	}

	err = persistence.Teams().AddTeamToEmplopyee(ctx, team, teamAddEmployee.EmployeeId)
	return teamAddEmployee, err
}
